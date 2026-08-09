package lxcrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// gcmStreamChunkSize is the plaintext size of each chunk in EncryptGCMStream.
const gcmStreamChunkSize = 64 * 1024 // 64 KB

const (
	gcmStreamMagic        = "LXGCM1"
	gcmStreamIDSize       = 16
	gcmStreamDataRecord   = 1
	gcmStreamFinalRecord  = 2
	gcmStreamRecordHeader = 5 // record type (1 B) + payload length (4 B)
)

var (
	errChunkTooLarge          = errors.New("lxcrypto: chunk length exceeds maximum allowed size")
	errGCMStreamInvalidFormat = errors.New("lxcrypto: invalid GCM stream format")
	errGCMStreamMissingFinal  = errors.New("lxcrypto: GCM stream is missing its final record")
	errGCMStreamUnexpected    = errors.New("lxcrypto: unexpected GCM stream record")
	errGCMStreamTrailingData  = errors.New("lxcrypto: trailing data after GCM stream final record")
	errGCMStreamTooManyChunks = errors.New("lxcrypto: too many GCM stream chunks")
)

// EncryptGCMStream encrypts data from src and writes an order- and
// truncation-authenticated AES-GCM stream to dst.
//
// Each record is bound to a random stream ID and its sequence number, and a final
// authenticated record is required. This rejects reordered, inserted, or
// cleanly truncated records during decryption.
//
// The output format is: ["LXGCM1"][stream ID][records...]. Each record is
// [record type (1 byte)][uint32 payload length][nonce][ciphertext][tag].
//
// Example:
//
//	err := lxcrypto.EncryptGCMStream(src, dst, key)
func EncryptGCMStream(src io.Reader, dst io.Writer, key []byte) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("lxcrypto: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("lxcrypto: %w", err)
	}

	streamID := make([]byte, gcmStreamIDSize)
	if _, err := rand.Read(streamID); err != nil {
		return fmt.Errorf("lxcrypto: %w", err)
	}
	if err := writeAll(dst, append([]byte(gcmStreamMagic), streamID...)); err != nil {
		return fmt.Errorf("lxcrypto: %w", err)
	}

	nonce := make([]byte, aead.NonceSize())
	buf := make([]byte, gcmStreamChunkSize)
	var recordIndex uint64

	for {
		n, readErr := io.ReadFull(src, buf)
		if readErr != nil && !errors.Is(readErr, io.ErrUnexpectedEOF) && !errors.Is(readErr, io.EOF) {
			return readErr
		}
		if n == 0 {
			break
		}
		if _, err := rand.Read(nonce); err != nil {
			return fmt.Errorf("lxcrypto: %w", err)
		}
		payload := aead.Seal(nonce, nonce, buf[:n], gcmStreamAAD(streamID, gcmStreamDataRecord, recordIndex))
		if err := writeGCMStreamRecord(dst, gcmStreamDataRecord, payload); err != nil {
			return fmt.Errorf("lxcrypto: %w", err)
		}
		if recordIndex == ^uint64(0) {
			return errGCMStreamTooManyChunks
		}
		recordIndex++
		if errors.Is(readErr, io.ErrUnexpectedEOF) || errors.Is(readErr, io.EOF) {
			break
		}
	}

	if _, err := rand.Read(nonce); err != nil {
		return fmt.Errorf("lxcrypto: %w", err)
	}
	finalPayload := aead.Seal(nonce, nonce, nil, gcmStreamAAD(streamID, gcmStreamFinalRecord, recordIndex))
	if err := writeGCMStreamRecord(dst, gcmStreamFinalRecord, finalPayload); err != nil {
		return fmt.Errorf("lxcrypto: %w", err)
	}
	return nil
}

// DecryptGCMStream decrypts an AES-GCM stream written by EncryptGCMStream.
// It returns an error when a record is tampered with,
// reordered, inserted from another stream, missing, or followed by data after
// the authenticated final record.
//
// Decrypted data may have been written to dst before a later authentication
// failure is found. To require all-or-nothing output, decrypt to a temporary
// file and rename it only after this function returns nil.
//
// Example:
//
//	err := lxcrypto.DecryptGCMStream(src, dst, key)
func DecryptGCMStream(src io.Reader, dst io.Writer, key []byte) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("lxcrypto: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("lxcrypto: %w", err)
	}

	header := make([]byte, len(gcmStreamMagic)+gcmStreamIDSize)
	if _, err := io.ReadFull(src, header); err != nil {
		return fmt.Errorf("lxcrypto: %w", err)
	}
	if !bytes.Equal(header[:len(gcmStreamMagic)], []byte(gcmStreamMagic)) {
		return errGCMStreamInvalidFormat
	}
	streamID := header[len(gcmStreamMagic):]
	maxPayloadLen := aead.NonceSize() + gcmStreamChunkSize + aead.Overhead()
	finalPayloadLen := aead.NonceSize() + aead.Overhead()
	var recordIndex uint64

	for {
		var recordHeader [gcmStreamRecordHeader]byte
		if _, err := io.ReadFull(src, recordHeader[:]); err != nil {
			if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
				return errGCMStreamMissingFinal
			}
			return fmt.Errorf("lxcrypto: %w", err)
		}

		recordType := recordHeader[0]
		payloadLen := binary.BigEndian.Uint32(recordHeader[1:])
		if int64(payloadLen) > int64(maxPayloadLen) {
			return errChunkTooLarge
		}
		payload := make([]byte, int(payloadLen))
		if _, err := io.ReadFull(src, payload); err != nil {
			return fmt.Errorf("lxcrypto: %w", err)
		}
		if len(payload) < finalPayloadLen {
			return errCiphertextTooShort
		}

		nonce, body := payload[:aead.NonceSize()], payload[aead.NonceSize():]
		switch recordType {
		case gcmStreamDataRecord:
			plaintext, err := aead.Open(nil, nonce, body, gcmStreamAAD(streamID, recordType, recordIndex))
			if err != nil {
				return fmt.Errorf("lxcrypto: %w", err)
			}
			if err := writeAll(dst, plaintext); err != nil {
				return fmt.Errorf("lxcrypto: %w", err)
			}
			if recordIndex == ^uint64(0) {
				return errGCMStreamTooManyChunks
			}
			recordIndex++

		case gcmStreamFinalRecord:
			if len(payload) != finalPayloadLen {
				return errGCMStreamUnexpected
			}
			plaintext, err := aead.Open(nil, nonce, body, gcmStreamAAD(streamID, recordType, recordIndex))
			if err != nil || len(plaintext) != 0 {
				return errGCMStreamUnexpected
			}
			return ensureGCMStreamEOF(src)

		default:
			return errGCMStreamUnexpected
		}
	}
}

func gcmStreamAAD(streamID []byte, recordType byte, recordIndex uint64) []byte {
	aad := make([]byte, len(streamID)+1+8)
	copy(aad, streamID)
	aad[len(streamID)] = recordType
	binary.BigEndian.PutUint64(aad[len(streamID)+1:], recordIndex)
	return aad
}

func writeGCMStreamRecord(dst io.Writer, recordType byte, payload []byte) error {
	var header [gcmStreamRecordHeader]byte
	header[0] = recordType
	binary.BigEndian.PutUint32(header[1:], uint32(len(payload)))
	if err := writeAll(dst, header[:]); err != nil {
		return err
	}
	return writeAll(dst, payload)
}

func writeAll(dst io.Writer, data []byte) error {
	for len(data) > 0 {
		n, err := dst.Write(data)
		if err != nil {
			return err
		}
		if n == 0 {
			return io.ErrShortWrite
		}
		data = data[n:]
	}
	return nil
}

func ensureGCMStreamEOF(src io.Reader) error {
	var trailing [1]byte
	n, err := src.Read(trailing[:])
	if n > 0 {
		return errGCMStreamTrailingData
	}
	if errors.Is(err, io.EOF) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("lxcrypto: %w", err)
	}
	return io.ErrNoProgress
}
