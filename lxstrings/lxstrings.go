package lxstrings

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	// Empty represents an empty string.
	Empty = ""

	// LF represents a line feed character.
	LF = "\n"

	// CR represents a carriage return character.
	CR = "\r"

	// CRLF represents a carriage return followed by a line feed.
	CRLF = "\r\n"
)

// Abbreviate shortens the string to the specified maxWidth, adding "..." if truncated.
// If the string is shorter than or equal to maxWidth, it is returned unchanged.
// If maxWidth is less than or equal to 3, it returns the first maxWidth characters.
func Abbreviate(s string, maxWidth int) string {
	runes := []rune(s)
	if len(runes) <= maxWidth {
		return s
	}
	if maxWidth <= 3 {
		return string(runes[:maxWidth])
	}
	return string(runes[:maxWidth-3]) + "..."
}

// Capitalize capitalizes the first character of the string.
// If the string is empty or starts with a non-letter, it is returned unchanged.
func Capitalize(s string) string {
	if IsBlank(s) {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Compare compares two strings lexicographically.
// It returns an integer comparing two strings lexicographically.
// The result will be 0 if s1 == s2, -1 if s1 < s2, and +1 if s1 > s2.
func Compare(s1, s2 string) int {
	return strings.Compare(s1, s2)
}

// Contains checks if the substring is present in the string.
func Contains(str, sub string) bool {
	return strings.Contains(str, sub)
}

// ContainsIgnoreCase checks if the substring is present in the string, ignoring case.
func ContainsIgnoreCase(str, sub string) bool {
	sLower := strings.ToLower(str)
	subLower := strings.ToLower(sub)
	return strings.Contains(sLower, subLower)
}

// ContainsAny checks if any of the specified characters are present in the string.
func ContainsAny(str string, chars ...rune) bool {
	for _, c := range str {
		for _, char := range chars {
			if c == char {
				return true
			}
		}
	}
	return false
}

// CompareIgnoreCase compares two strings lexicographically, ignoring case.
// It returns an integer comparing two strings lexicographically, ignoring case.
// The result will be 0 if str1 == str2, -1 if str1 < str2, and +1 if str1 > str2.
func CompareIgnoreCase(str1, str2 string) int {
	s1Lower := strings.ToLower(str1)
	s2Lower := strings.ToLower(str2)
	return strings.Compare(s1Lower, s2Lower)
}

// IsEmpty checks if the given string is empty.
func IsEmpty(str string) bool {
	return len(str) == 0
}

// IsNotEmpty checks if the given string is not empty.
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

// IsBlank checks if the given string is blank (empty or only whitespace).
func IsBlank(str string) bool {
	for _, r := range str {
		if r != ' ' && r != '\n' && r != '\t' && r != '\r' {
			return false
		}
	}
	return true
}

// IsNotBlank checks if the given string is not blank.
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

// IsAlpha checks if the given string contains only alphabetic characters.
func IsAlpha(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, r := range str {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// IsNumeric checks if the given string contains only numeric characters.
func IsNumeric(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, r := range str {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsAlphaNumeric checks if the given string contains only alphanumeric characters.
func IsAlphaNumeric(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, r := range str {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsSpace reports whether str is non-empty and contains only Unicode whitespace.
func IsSpace(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, r := range str {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// Index returns the index of the first occurrence of substr in s, or -1 if not found.
func Index(str, sub string) int {
	return strings.Index(str, sub)
}

// IndexFrom returns the index of sub in str starting from fromInd,
// or -1 if sub is not present.
func IndexFrom(str, sub string, fromInd int) int {
	if fromInd < 0 || fromInd >= len(str) {
		return -1
	}
	ind := Index(str[fromInd:], sub)
	if ind == -1 {
		return -1
	}
	return fromInd + ind
}

// LastIndex returns the index of the last occurrence of substr in s, or -1 if not found.
func LastIndex(str, sub string) int {
	return strings.LastIndex(str, sub)
}

// LastIndexIgnoreCase returns the index of the last occurrence of substr in s, ignoring case, or -1 if not found.
func LastIndexIgnoreCase(str, sub string) int {
	sLower := strings.ToLower(str)
	substrLower := strings.ToLower(sub)
	return strings.LastIndex(sLower, substrLower)
}

// Length returns the length of the string.
func Length(s string) int {
	return len(s)
}

// RuneCount returns the number of runes (Unicode code points) in the string.
// This is useful when you need the character count as seen by users rather
// than the byte length returned by Length.
//
// Example:
//
//	RuneCount("こんにちは") // 5
func RuneCount(s string) int {
	return utf8.RuneCountInString(s)
}

// LowerCase converts the string to lowercase.
func LowerCase(s string) string {
	return strings.ToLower(s)
}

// UpperCase converts the string to uppercase.
func UpperCase(str string) string {
	return strings.ToUpper(str)
}

// IndexIgnoreCase returns the index of the first occurrence of substr in s, ignoring case, or -1 if not found.
func IndexIgnoreCase(str, substr string) int {
	sLower := LowerCase(str)
	substrLower := LowerCase(substr)
	return Index(sLower, substrLower)
}

// Equals checks if two strings are equal.
func Equals(str1, str2 string) bool {
	return str1 == str2
}

// NotEquals checks if two strings are not equal.
func NotEquals(str1, str2 string) bool {
	return !Equals(str1, str2)
}

// EqualsIgnoreCase checks if two strings are equal, ignoring case.
func EqualsIgnoreCase(str1, str2 string) bool {
	return strings.EqualFold(str1, str2)
}

// NotEqualsIgnoreCase checks if two strings are not equal, ignoring case.
func NotEqualsIgnoreCase(str1, str2 string) bool {
	return !EqualsIgnoreCase(str1, str2)
}

// TrimSpace removes leading and trailing whitespace from the string.
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// Trim removes all leading and trailing characters specified in cutset from the string.
func Trim(str string, cutset string) string {
	return strings.Trim(str, cutset)
}

// TrimLeft removes all leading characters specified in cutset from the string.
func TrimLeft(str string, cutset string) string {
	return strings.TrimLeft(str, cutset)
}

// TrimRight removes all trailing characters specified in cutset from the string.
func TrimRight(str string, cutset string) string {
	return strings.TrimRight(str, cutset)
}

func Truncate(str string, maxWidth int) string {
	runes := []rune(str)
	if len(runes) <= maxWidth {
		return str
	}
	return string(runes[:maxWidth])
}

// Split splits the string by the specified separator and returns a slice of substrings.
func Split(str, sep string) []string {
	return strings.Split(str, sep)
}

// Join joins a slice of strings into a single string with the specified separator.
func Join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}

// Repeat returns a new string consisting of count copies of the string s.
func Repeat(str string, count int) string {
	return strings.Repeat(str, count)
}

// StartBy checks if the string starts with the specified prefix.
func StartBy(str, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}

// StartByIgnoreCase checks if the string starts with the specified prefix, ignoring case.
func StartByIgnoreCase(str, prefix string) bool {
	sLower := LowerCase(str)
	prefixLower := LowerCase(prefix)
	return strings.HasPrefix(sLower, prefixLower)
}

// StartByAny checks if the string starts with any of the specified prefixes.
func StartByAny(str string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(str, prefix) {
			return true
		}
	}
	return false
}

func StartByAnyIgnoreCase(str string, prefixes ...string) bool {
	sLower := LowerCase(str)
	for _, prefix := range prefixes {
		prefixLower := LowerCase(prefix)
		if strings.HasPrefix(sLower, prefixLower) {
			return true
		}
	}
	return false
}

// EndBy checks if the string ends with the specified suffix.
func EndBy(str, suffix string) bool {
	return strings.HasSuffix(str, suffix)
}

// EndByIgnoreCase checks if the string ends with the specified suffix, ignoring case.
func EndByIgnoreCase(str, suffix string) bool {
	sLower := LowerCase(str)
	suffixLower := LowerCase(suffix)
	return strings.HasSuffix(sLower, suffixLower)
}

// EndByAny checks if the string ends with any of the specified suffixes.
func EndByAny(str string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(str, suffix) {
			return true
		}
	}
	return false
}

// EndByAnyIgnoreCase checks if the string ends with any of the specified suffixes, ignoring case.
func EndByAnyIgnoreCase(str string, suffixes ...string) bool {
	sLower := LowerCase(str)
	for _, suffix := range suffixes {
		suffixLower := LowerCase(suffix)
		if strings.HasSuffix(sLower, suffixLower) {
			return true
		}
	}
	return false
}

// Replace replaces occurrences of old with new in the string str, up to n times.
// If n is -1, all occurrences are replaced.
func Replace(str, old, new string, n int) string {
	return strings.Replace(str, old, new, n)
}

// ReplaceAll replaces all occurrences of old with new in the string str.
func ReplaceAll(str, old, new string) string {
	return strings.ReplaceAll(str, old, new)
}

// Remove removes all occurrences of sub from the string str.
func Remove(str, sub string) string {
	return strings.ReplaceAll(str, sub, "")
}

// RemoveIgnoreCase removes all occurrences of substr from str, ignoring case.
func RemoveIgnoreCase(str, substr string) string {
	if substr == "" {
		return str
	}

	lowerS := LowerCase(str)
	lowerSub := LowerCase(substr)

	var b strings.Builder
	i := 0

	for {
		j := Index(lowerS[i:], lowerSub)
		if j < 0 {
			b.WriteString(str[i:])
			break
		}

		j += i
		b.WriteString(str[i:j])
		i = j + len(substr)
	}

	return b.String()
}

// RemoveAny removes all occurrences of the provided substrings from str.
func RemoveAny(str string, subs ...string) string {
	result := str
	for _, sub := range subs {
		result = Remove(result, sub)
	}
	return result
}

// RemoveAnyIgnoreCase removes all occurrences of the specified substrings from the string s, ignoring case.
func RemoveAnyIgnoreCase(str string, subs ...string) string {
	result := str
	for _, sub := range subs {
		result = RemoveIgnoreCase(result, sub)
	}
	return result
}

// Reverse returns str with its characters in reverse order.
func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// SubString returns the substring of str from start index to end index.
// If end is -1, it returns the substring from start to the end of the string.
func SubString(str string, start, end int) string {
	runes := []rune(str)
	if end == -1 || end > len(runes) {
		end = len(runes)
	}
	if start < 0 || start > end {
		return ""
	}
	return string(runes[start:end])
}

// SubStringBefore returns the substring before the first occurrence of sep.
// If sep is not found, it returns an empty string.
func SubStringBefore(str, sep string) string {
	index := Index(str, sep)
	if index == -1 {
		return ""
	}
	return str[:index]
}

// SubStringBeforeIgnoreCase returns the substring before the first occurrence of sep, ignoring case.
// If sep is not found, it returns an empty string.
func SubStringBeforeIgnoreCase(str, sep string) string {
	sLower := LowerCase(str)
	sepLower := LowerCase(sep)
	index := strings.Index(sLower, sepLower)
	if index == -1 {
		return ""
	}
	return str[:index]
}

// SubStringAfter returns the substring after the first occurrence of sep.
// If sep is not found, it returns an empty string.
func SubStringAfter(str, sep string) string {
	index := Index(str, sep)
	if index == -1 {
		return ""
	}
	return str[index+len(sep):]
}

// SubStringAfterIgnoreCase returns the substring after the first occurrence of sep, ignoring case.
// If sep is not found, it returns an empty string.
func SubStringAfterIgnoreCase(str, sep string) string {
	sLower := LowerCase(str)
	sepLower := LowerCase(sep)
	index := strings.Index(sLower, sepLower)
	if index == -1 {
		return ""
	}
	return str[index+len(sep):]
}

// PadLeft pads the string on the left with the specified padStr until it reaches the desired length.
func PadLeft(str string, length int, padStr string) string {
	if IsEmpty(str) || IsEmpty(padStr) || Length(str) >= length {
		return str
	}

	padLen := length - Length(str)
	var b strings.Builder
	b.Grow(length)

	for b.Len() < padLen {
		b.WriteString(padStr)
	}

	padded := b.String()
	return padded[len(padded)-padLen:] + str
}

// PadRight pads the string on the right with the specified padStr until it reaches the desired length.
func PadRight(str string, length int, padStr string) string {
	if IsEmpty(str) || IsEmpty(padStr) || Length(str) >= length {
		return str
	}

	var b strings.Builder
	b.Grow(length)
	b.WriteString(str)

	for b.Len() < length {
		b.WriteString(padStr)
	}

	result := b.String()
	return result[:length]
}

// PadCenter pads the string on both sides with the specified padStr until it reaches the desired length.
func PadCenter(str string, length int, padStr string) string {
	if IsEmpty(str) || IsEmpty(padStr) {
		return str
	}

	sLen := len(str)
	if sLen >= length {
		return str
	}

	totalPad := length - sLen
	left := totalPad / 2
	right := totalPad - left

	var b strings.Builder
	b.Grow(length)

	for n := left; n > 0; {
		if len(padStr) <= n {
			b.WriteString(padStr)
			n -= len(padStr)
		} else {
			b.WriteString(padStr[:n])
			n = 0
		}
	}

	b.WriteString(str)

	for n := right; n > 0; {
		if len(padStr) <= n {
			b.WriteString(padStr)
			n -= len(padStr)
		} else {
			b.WriteString(padStr[:n])
			n = 0
		}
	}

	return b.String()
}

// CountMatches counts non-overlapping occurrences of sub in str.
func CountMatches(str, sub string) int {
	if IsEmpty(str) || IsEmpty(sub) {
		return 0
	}

	count := 0
	ind := 0

	for {
		ind = IndexFrom(str, sub, ind)
		if ind == -1 {
			break
		}
		count++
		ind += len(sub)
	}

	return count
}

// DefaultIfEmpty returns defaultStr if str is empty, otherwise returns str.
func DefaultIfEmpty(str, defaultStr string) string {
	if IsEmpty(str) {
		return defaultStr
	}
	return str
}

// DefaultIfBlank returns defaultStr if str is blank, otherwise returns str.
func DefaultIfBlank(str, defaultStr string) string {
	if IsBlank(str) {
		return defaultStr
	}
	return str
}

func StartWith(str, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}

// StartWithIgnoreCase checks if the string starts with the specified prefix, ignoring case.
func StartWithIgnoreCase(str, prefix string) bool {
	sLower := LowerCase(str)
	prefixLower := LowerCase(prefix)
	return strings.HasPrefix(sLower, prefixLower)
}

// StartWithAny checks if the string starts with any of the specified prefixes.
func StartWithAny(str string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(str, prefix) {
			return true
		}
	}
	return false
}

// StartWithAnyIgnoreCase checks if the string starts with any of the specified prefixes, ignoring case.
func StartWithAnyIgnoreCase(str string, prefixes ...string) bool {
	sLower := LowerCase(str)
	for _, prefix := range prefixes {
		prefixLower := LowerCase(prefix)
		if strings.HasPrefix(sLower, prefixLower) {
			return true
		}
	}
	return false
}

// EndWith checks if the string ends with the specified suffix.
func EndWith(str, suffix string) bool {
	return strings.HasSuffix(str, suffix)
}

// EndWithIgnoreCase checks if the string ends with the specified suffix, ignoring case.
func EndWithIgnoreCase(str, suffix string) bool {
	sLower := LowerCase(str)
	suffixLower := LowerCase(suffix)
	return strings.HasSuffix(sLower, suffixLower)
}

// EndWithAny checks if the string ends with any of the specified suffixes.
func EndWithAny(str string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(str, suffix) {
			return true
		}
	}
	return false
}

// EndWithAnyIgnoreCase checks if the string ends with any of the specified suffixes, ignoring case.
func EndWithAnyIgnoreCase(str string, suffixes ...string) bool {
	sLower := LowerCase(str)
	for _, suffix := range suffixes {
		suffixLower := LowerCase(suffix)
		if strings.HasSuffix(sLower, suffixLower) {
			return true
		}
	}
	return false
}
