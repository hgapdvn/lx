package lxio

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// ======================== File Extension Constants ========================

// Document extensions
const (
	ExtPDF   = ".pdf"
	ExtDOC   = ".doc"
	ExtDOCX  = ".docx"
	ExtTXT   = ".txt"
	ExtRTF   = ".rtf"
	ExtODT   = ".odt"
	ExtPages = ".pages"
)

// Image extensions
const (
	ExtJPG  = ".jpg"
	ExtJPEG = ".jpeg"
	ExtPNG  = ".png"
	ExtGIF  = ".gif"
	ExtBMP  = ".bmp"
	ExtTIFF = ".tiff"
	ExtSVG  = ".svg"
	ExtWebP = ".webp"
	ExtICO  = ".ico"
)

// Archive extensions
const (
	ExtZIP   = ".zip"
	ExtRAR   = ".rar"
	ExtTAR   = ".tar"
	ExtGZ    = ".gz"
	ExtBZ2   = ".bz2"
	ExtTarGZ = ".tar.gz"
)

// Code extensions
const (
	ExtGO   = ".go"
	ExtPY   = ".py"
	ExtJS   = ".js"
	ExtTS   = ".ts"
	ExtJava = ".java"
	ExtCPP  = ".cpp"
	ExtC    = ".c"
	ExtH    = ".h"
	ExtRB   = ".rb"
	ExtPHP  = ".php"
	ExtJSON = ".json"
	ExtXML  = ".xml"
	ExtYAML = ".yaml"
	ExtYML  = ".yml"
	ExtCSV  = ".csv"
)

// Audio extensions
const (
	ExtMP3  = ".mp3"
	ExtWAV  = ".wav"
	ExtFLAC = ".flac"
	ExtAAC  = ".aac"
	ExtOGG  = ".ogg"
	ExtM4A  = ".m4a"
)

// Video extensions
const (
	ExtMP4  = ".mp4"
	ExtAVI  = ".avi"
	ExtMKV  = ".mkv"
	ExtMOV  = ".mov"
	ExtWMV  = ".wmv"
	ExtFLV  = ".flv"
	ExtWEBM = ".webm"
)

// Grouped extension slices for convenience
var (
	DocumentExts = []string{ExtPDF, ExtDOC, ExtDOCX, ExtTXT, ExtRTF, ExtODT, ExtPages}
	ImageExts    = []string{ExtJPG, ExtJPEG, ExtPNG, ExtGIF, ExtBMP, ExtTIFF, ExtSVG, ExtWebP, ExtICO}
	ArchiveExts  = []string{ExtZIP, ExtRAR, ExtTAR, ExtGZ, ExtBZ2, ExtTarGZ}
	CodeExts     = []string{ExtGO, ExtPY, ExtJS, ExtTS, ExtJava, ExtCPP, ExtC, ExtH, ExtRB, ExtPHP, ExtJSON, ExtXML, ExtYAML, ExtYML, ExtCSV}
	AudioExts    = []string{ExtMP3, ExtWAV, ExtFLAC, ExtAAC, ExtOGG, ExtM4A}
	VideoExts    = []string{ExtMP4, ExtAVI, ExtMKV, ExtMOV, ExtWMV, ExtFLV, ExtWEBM}
)

// ListFiles returns a sorted slice of file names (not directories) in the given directory.
// The slice contains only the base names of files, not full paths.
// It returns an error if the directory doesn't exist or cannot be read.
//
// Example:
//
//	files, err := lxio.ListFiles("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, file := range files {
//		fmt.Println(file)
//	}
func ListFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	sort.Strings(files)
	return files, nil
}

// ListDirs returns a sorted slice of directory names (not files) in the given directory.
// The slice contains only the base names of directories, not full paths.
// It returns an error if the directory doesn't exist or cannot be read.
//
// Example:
//
//	dirs, err := lxio.ListDirs("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, dir := range dirs {
//		fmt.Println(dir)
//	}
func ListDirs(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}

	sort.Strings(dirs)
	return dirs, nil
}

// ListAll returns a sorted slice of all entry names (both files and directories) in the given directory.
// The slice contains only the base names of entries, not full paths.
// It returns an error if the directory doesn't exist or cannot be read.
//
// Example:
//
//	entries, err := lxio.ListAll("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, entry := range entries {
//		fmt.Println(entry)
//	}
func ListAll(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, entry := range entries {
		names = append(names, entry.Name())
	}

	sort.Strings(names)
	return names, nil
}

// CountFiles returns the number of files (not directories) in the given directory.
// It does not recursively count files in subdirectories.
// It returns an error if the directory doesn't exist or cannot be read.
//
// Example:
//
//	count, err := lxio.CountFiles("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Files: %d\n", count)
func CountFiles(dir string) (int, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			count++
		}
	}

	return count, nil
}

// CountDirs returns the number of directories (not files) in the given directory.
// It does not recursively count directories in subdirectories.
// It returns an error if the directory doesn't exist or cannot be read.
//
// Example:
//
//	count, err := lxio.CountDirs("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Directories: %d\n", count)
func CountDirs(dir string) (int, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, entry := range entries {
		if entry.IsDir() {
			count++
		}
	}

	return count, nil
}

// CountFilesRecursive returns the total number of files in the root directory
// and all subdirectories recursively.
// It returns an error if the root directory doesn't exist or cannot be read.
//
// Example:
//
//	count, err := lxio.CountFilesRecursive("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Total files: %d\n", count)
func CountFilesRecursive(root string) (int, error) {
	count := 0
	err := WalkFiles(root, func(path string) error {
		count++
		return nil
	})
	return count, err
}

// CountDirsRecursive returns the total number of directories in the root directory
// and all subdirectories recursively (excluding the root directory itself).
// It returns an error if the root directory doesn't exist or cannot be read.
//
// Example:
//
//	count, err := lxio.CountDirsRecursive("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Total directories: %d\n", count)
func CountDirsRecursive(root string) (int, error) {
	count := 0
	err := WalkDirs(root, func(path string) error {
		count++
		return nil
	})
	return count, err
}

// WalkFiles recursively walks through the root directory and all subdirectories,
// calling fn for each file found. The path passed to fn is relative to root
// and always uses forward slashes as separators, regardless of the operating system.
// If fn returns a non-nil error, the walk stops and that error is returned.
// It returns an error if the root directory doesn't exist or cannot be read.
//
// Example:
//
//	err := lxio.WalkFiles("/path/to/dir", func(path string) error {
//		fmt.Println(path) // e.g., "subdir/file.txt"
//		return nil
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
func WalkFiles(root string, fn func(path string) error) error {
	return walkFilesInternal(root, "", fn)
}

// walkFilesInternal is a helper function for WalkFiles.
// relRoot tracks the relative path from the original root.
func walkFilesInternal(currentPath string, relRoot string, fn func(path string) error) error {
	entries, err := os.ReadDir(currentPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		var relPath string
		if relRoot == "" {
			relPath = entry.Name()
		} else {
			relPath = relRoot + "/" + entry.Name()
		}

		if entry.IsDir() {
			// Recurse into subdirectory
			fullPath := filepath.Join(currentPath, entry.Name())
			if err := walkFilesInternal(fullPath, relPath, fn); err != nil {
				return err
			}
		} else {
			// Call fn for this file
			if err := fn(relPath); err != nil {
				return err
			}
		}
	}

	return nil
}

// WalkDirs recursively walks through the root directory and all subdirectories,
// calling fn for each directory found (excluding the root directory itself).
// The path passed to fn is relative to root and always uses forward slashes as separators,
// regardless of the operating system.
// If fn returns a non-nil error, the walk stops and that error is returned.
// It returns an error if the root directory doesn't exist or cannot be read.
//
// Example:
//
//	err := lxio.WalkDirs("/path/to/dir", func(path string) error {
//		fmt.Println(path) // e.g., "subdir", "subdir/nested"
//		return nil
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
func WalkDirs(root string, fn func(path string) error) error {
	return walkDirsInternal(root, "", fn)
}

// walkDirsInternal is a helper function for WalkDirs.
// relRoot tracks the relative path from the original root.
func walkDirsInternal(currentPath string, relRoot string, fn func(path string) error) error {
	entries, err := os.ReadDir(currentPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			var relPath string
			if relRoot == "" {
				relPath = entry.Name()
			} else {
				relPath = relRoot + "/" + entry.Name()
			}

			// Call fn for this directory
			if err := fn(relPath); err != nil {
				return err
			}

			// Recurse into subdirectory
			fullPath := filepath.Join(currentPath, entry.Name())
			if err := walkDirsInternal(fullPath, relPath, fn); err != nil {
				return err
			}
		}
	}

	return nil
}

// ======================== Extension Filter Functions ========================

// ListFilesByExt returns a sorted slice of file names that match any of the provided extensions.
// Extensions should include the dot (e.g., ".pdf", ".txt").
// The extension matching is case-insensitive.
// Returns an empty slice if no files match.
//
// Example:
//
//	files, err := lxio.ListFilesByExt("/path/to/dir", ".txt", ".md")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, file := range files {
//		fmt.Println(file) // e.g., "readme.txt", "guide.md"
//	}
func ListFilesByExt(dir string, ext ...string) ([]string, error) {
	if len(ext) == 0 {
		return []string{}, nil
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// Normalize extensions to lowercase for case-insensitive comparison
	extLower := make([]string, len(ext))
	for i, e := range ext {
		extLower[i] = strings.ToLower(e)
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			fileName := strings.ToLower(entry.Name())
			for _, e := range extLower {
				// Check if file ends with the extension or matches single extension with filepath.Ext
				if strings.HasSuffix(fileName, e) || strings.ToLower(filepath.Ext(entry.Name())) == e {
					files = append(files, entry.Name())
					break
				}
			}
		}
	}

	sort.Strings(files)
	return files, nil
}

// ListPdfFiles returns a sorted slice of PDF file names in the directory.
//
// Example:
//
//	pdfs, err := lxio.ListPdfFiles("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
func ListPdfFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtPDF)
}

// ListDocFiles returns a sorted slice of .doc file names in the directory.
func ListDocFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtDOC)
}

// ListDocxFiles returns a sorted slice of .docx file names in the directory.
func ListDocxFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtDOCX)
}

// ListTxtFiles returns a sorted slice of .txt file names in the directory.
func ListTxtFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtTXT)
}

// ListRtfFiles returns a sorted slice of .rtf file names in the directory.
func ListRtfFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtRTF)
}

// ListOdtFiles returns a sorted slice of .odt file names in the directory.
func ListOdtFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtODT)
}

// ListPagesFiles returns a sorted slice of .pages file names in the directory.
func ListPagesFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtPages)
}

// ListDocumentFiles returns a sorted slice of document file names (pdf, doc, docx, txt, rtf, odt, pages).
//
// Example:
//
//	docs, err := lxio.ListDocumentFiles("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
func ListDocumentFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, DocumentExts...)
}

// ======================== Image File Functions ========================

// ListJpgFiles returns a sorted slice of .jpg file names in the directory.
func ListJpgFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtJPG)
}

// ListJpegFiles returns a sorted slice of .jpeg file names in the directory.
func ListJpegFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtJPEG)
}

// ListPngFiles returns a sorted slice of .png file names in the directory.
func ListPngFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtPNG)
}

// ListGifFiles returns a sorted slice of .gif file names in the directory.
func ListGifFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtGIF)
}

// ListBmpFiles returns a sorted slice of .bmp file names in the directory.
func ListBmpFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtBMP)
}

// ListTiffFiles returns a sorted slice of .tiff file names in the directory.
func ListTiffFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtTIFF)
}

// ListSvgFiles returns a sorted slice of .svg file names in the directory.
func ListSvgFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtSVG)
}

// ListWebpFiles returns a sorted slice of .webp file names in the directory.
func ListWebpFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtWebP)
}

// ListIcoFiles returns a sorted slice of .ico file names in the directory.
func ListIcoFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtICO)
}

// ListImageFiles returns a sorted slice of image file names (jpg, jpeg, png, gif, bmp, tiff, svg, webp, ico).
//
// Example:
//
//	images, err := lxio.ListImageFiles("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
func ListImageFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ImageExts...)
}

// ======================== Archive File Functions ========================

// ListZipFiles returns a sorted slice of .zip file names in the directory.
func ListZipFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtZIP)
}

// ListRarFiles returns a sorted slice of .rar file names in the directory.
func ListRarFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtRAR)
}

// ListTarFiles returns a sorted slice of .tar file names in the directory.
func ListTarFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtTAR)
}

// ListGzFiles returns a sorted slice of .gz file names in the directory.
func ListGzFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtGZ)
}

// ListBz2Files returns a sorted slice of .bz2 file names in the directory.
func ListBz2Files(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtBZ2)
}

// ListTarGzFiles returns a sorted slice of .tar.gz file names in the directory.
func ListTarGzFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtTarGZ)
}

// ListArchiveFiles returns a sorted slice of archive file names (zip, rar, tar, gz, bz2, tar.gz).
//
// Example:
//
//	archives, err := lxio.ListArchiveFiles("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
func ListArchiveFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ArchiveExts...)
}

// ======================== Code File Functions ========================

// ListGoFiles returns a sorted slice of .go file names in the directory.
func ListGoFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtGO)
}

// ListPyFiles returns a sorted slice of .py file names in the directory.
func ListPyFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtPY)
}

// ListJsFiles returns a sorted slice of .js file names in the directory.
func ListJsFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtJS)
}

// ListTsFiles returns a sorted slice of .ts file names in the directory.
func ListTsFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtTS)
}

// ListJavaFiles returns a sorted slice of .java file names in the directory.
func ListJavaFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtJava)
}

// ListCppFiles returns a sorted slice of .cpp file names in the directory.
func ListCppFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtCPP)
}

// ListCFiles returns a sorted slice of .c file names in the directory.
func ListCFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtC)
}

// ListHFiles returns a sorted slice of .h file names in the directory.
func ListHFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtH)
}

// ListRbFiles returns a sorted slice of .rb file names in the directory.
func ListRbFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtRB)
}

// ListPhpFiles returns a sorted slice of .php file names in the directory.
func ListPhpFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtPHP)
}

// ListJsonFiles returns a sorted slice of .json file names in the directory.
func ListJsonFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtJSON)
}

// ListXmlFiles returns a sorted slice of .xml file names in the directory.
func ListXmlFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtXML)
}

// ListYamlFiles returns a sorted slice of .yaml file names in the directory.
func ListYamlFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtYAML)
}

// ListYmlFiles returns a sorted slice of .yml file names in the directory.
func ListYmlFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtYML)
}

// ListCsvFiles returns a sorted slice of .csv file names in the directory.
func ListCsvFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtCSV)
}

// ListCodeFiles returns a sorted slice of source code file names (go, py, js, ts, java, cpp, c, h, rb, php, json, xml, yaml, yml, csv).
//
// Example:
//
//	codes, err := lxio.ListCodeFiles("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
func ListCodeFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, CodeExts...)
}

// ======================== Audio File Functions ========================

// ListMp3Files returns a sorted slice of .mp3 file names in the directory.
func ListMp3Files(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtMP3)
}

// ListWavFiles returns a sorted slice of .wav file names in the directory.
func ListWavFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtWAV)
}

// ListFlacFiles returns a sorted slice of .flac file names in the directory.
func ListFlacFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtFLAC)
}

// ListAacFiles returns a sorted slice of .aac file names in the directory.
func ListAacFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtAAC)
}

// ListOggFiles returns a sorted slice of .ogg file names in the directory.
func ListOggFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtOGG)
}

// ListM4aFiles returns a sorted slice of .m4a file names in the directory.
func ListM4aFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtM4A)
}

// ListAudioFiles returns a sorted slice of audio file names (mp3, wav, flac, aac, ogg, m4a).
//
// Example:
//
//	audios, err := lxio.ListAudioFiles("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
func ListAudioFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, AudioExts...)
}

// ======================== Video File Functions ========================

// ListMp4Files returns a sorted slice of .mp4 file names in the directory.
func ListMp4Files(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtMP4)
}

// ListAviFiles returns a sorted slice of .avi file names in the directory.
func ListAviFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtAVI)
}

// ListMkvFiles returns a sorted slice of .mkv file names in the directory.
func ListMkvFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtMKV)
}

// ListMovFiles returns a sorted slice of .mov file names in the directory.
func ListMovFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtMOV)
}

// ListWmvFiles returns a sorted slice of .wmv file names in the directory.
func ListWmvFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtWMV)
}

// ListFlvFiles returns a sorted slice of .flv file names in the directory.
func ListFlvFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtFLV)
}

// ListWebmFiles returns a sorted slice of .webm file names in the directory.
func ListWebmFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, ExtWEBM)
}

// ListVideoFiles returns a sorted slice of video file names (mp4, avi, mkv, mov, wmv, flv, webm).
//
// Example:
//
//	videos, err := lxio.ListVideoFiles("/path/to/dir")
//	if err != nil {
//		log.Fatal(err)
//	}
func ListVideoFiles(dir string) ([]string, error) {
	return ListFilesByExt(dir, VideoExts...)
}

// ListFilesFunc returns a sorted slice of file names that match the predicate function.
// The predicate receives the base name of each file (not directories).
// Returns an empty slice if no files match.
//
// Example:
//
//	// Find all files starting with "test"
//	files, err := lxio.ListFilesFunc("/path/to/dir", func(name string) bool {
//		return strings.HasPrefix(name, "test")
//	})
func ListFilesFunc(dir string, filter func(name string) bool) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() && filter(entry.Name()) {
			files = append(files, entry.Name())
		}
	}

	sort.Strings(files)
	return files, nil
}
