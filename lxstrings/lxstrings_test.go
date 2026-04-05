package lxstrings_test

import (
	"testing"

	"github.com/hgapdvn/lx/lxstrings"
)

func TestAbbreviate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxWidth int
		expected string
	}{
		{"Happy case", "Hello, World!", 5, "He..."},
		{"No abbreviation", "Hello", 10, "Hello"},
		{"Short to abbreviate", "GoLang", 3, "GoL"},
		{"Empty string", "", 0, ""},
		{"Exact", "Exact", 5, "Exact"},
		{"Long string", "This is a longer string", 8, "This ..."},
		{"Emoij", "😊😊😊😊😊", 4, "😊..."},
		{"Japanese", "こんにちは世界", 6, "こんに..."},
		{"Empty string", "", 5, ""},
		{"Shorter than input", "Test", 2, "Te"},
		{"Equal to input", "Test", 4, "Test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.Abbreviate(tt.input, tt.maxWidth)
			if result != tt.expected {
				t.Errorf("Abbreviate(%q, %d) = %q; want %q", tt.input, tt.maxWidth, result, tt.expected)
			}
		})
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"One word", "hello", "Hello"},
		{"Already capitalized", "Hello", "Hello"},
		{"Mixed case", "hELLO", "HELLO"},
		{"Empty string", "", ""},
		{"Space", " ", " "},
		{"Newline", "\n", "\n"},
		{"Carriage return", "\r", "\r"},
		{"Tab", "\t", "\t"},
		{"Single char with a new line", "a \n", "A \n"},
		{"Carriage return test", "\rtest", "\rtest"},
		{"Single /", "/", "/"},
		{"New line windows", "\r\n", "\r\n"},
		{"All uppercase expect first char", "ABC", "ABC"},
		{"Numbers", "111", "111"},
		{"Single char uppercase", "A", "A"},
		{"First number in string", "1test", "1test"},
		{"emoji", "😊emoji", "😊emoji"},
		{"Japanese", "こんにちは", "こんにちは"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.Capitalize(tt.input)
			if result != tt.expected {
				t.Errorf("Capitalize(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		name     string
		s1, s2   string
		expected int
	}{
		{"Less than", "apple", "banana", -1},
		{"Greater than", "banana", "apple", 1},
		{"Equal", "cherry", "cherry", 0},
		{"Empty strings", "", "", 0},
		{"Case sensitive - lowercase first", "a", "A", 1},
		{"Case sensitive - uppercase first", "A", "a", -1},
		{"Less than number of characters", "abc", "abcd", -1},
		{"Greater than number of characters", "abcd", "abc", 1},
		{"Japanese", "こんにちは", "こんばんは", -1},
		{"Emoji equal", "😊", "😊", 0},
		{"Emoji less than", "😊", "😢", -1},
		{"Emoji greater than", "😢", "😊", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.Compare(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("Compare(%q, %q) = %d; want %d", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}

func TestCompareIgnoreCase(t *testing.T) {
	tests := []struct {
		name     string
		s1, s2   string
		expected int
	}{
		{"Less than case insensitive", "apple", "BANANA", -1},
		{"Greater than case insensitive", "BANANA", "apple", 1},
		{"Equal case insensitive", "cherry", "CHERRY", 0},
		{"Empty strings", "", "", 0},
		{"Case sensitive - lowercase first", "a", "A", 0},
		{"Case sensitive - uppercase first", "A", "a", 0},
		{"Less than number of characters", "abc", "ABCD", -1},
		{"Greater than number of characters", "ABCD", "abc", 1},
		{"Japanese equal", "こんにちは", "こんにちは", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.CompareIgnoreCase(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("CompareIgnoreCase(%q, %q) = %d; want %d", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		str, sub string
		expected bool
	}{
		{"Happy case", "hello world", "world", true},
		{"Insensitive", "hello world", "WORLD", false},
		{"Match first word", "golang", "go", true},
		{"Match last word", "golang", "lang", true},
		{"None match sensitive", "test", "TEST", false},
		{"Empty strings", "", "", true},
		{"Match empty string", "non-empty", "", true},
		{"Empty string in non-empty", "", "non-empty", false},
		{"Japanese", "こんにちは世界", "世界", true},
		{"Emoji", "😊emoji😊", "emoji", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.Contains(tt.str, tt.sub)
			if result != tt.expected {
				t.Errorf("Contains(%q, %q) = %v; want %v", tt.str, tt.sub, result, tt.expected)
			}
		})
	}
}

func TestContainsIgnoreCase(t *testing.T) {
	tests := []struct {
		name        string
		str, substr string
		expected    bool
	}{
		{"Case insensitive", "hello world", "WORLD", true},
		{"Case sensitive - lowercase first", "golang", "golang", true},
		{"Case sensitive - uppercase first", "GOLANG", "golang", true},
		{"Case sensitive - mixed case", "CaseSensitive", "casesensitive", true},
		{"Empty strings", "", "", true},
		{"Non-empty string in empty string", "", "non-empty", false},
		{"Japanese", "こんにちは世界", "世界", true},
		{"Emoji", "😊emoji😊", "emoji", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.ContainsIgnoreCase(tt.str, tt.substr)
			if result != tt.expected {
				t.Errorf("ContainsIgnoreCase(%q, %q) = %v; want %v", tt.str, tt.substr, result, tt.expected)
			}
		})
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		chars    []rune
		expected bool
	}{
		{"Multiple single characters", "hello", []rune{'a', 'e', 'i'}, true},
		{"No match single characters", "world", []rune{'x', 'y', 'z'}, false},
		{"Match first character", "golang", []rune{'g', 'o'}, true},
		{"No match - number", "test", []rune{'1', '2', '3'}, false},
		{"Empty string", "", []rune{'a', 'b'}, false},
		{"None empty string - no chars", "non-empty", []rune{}, false},
		{"Japanese", "こんにちは", []rune{'に', 'は'}, true},
		{"Emoji", "😊emoji😊", []rune{'😊'}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.ContainsAny(tt.str, tt.chars...)
			if result != tt.expected {
				t.Errorf("ContainsAny(%q, %v) = %v; want %v", tt.str, tt.chars, result, tt.expected)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty string", "", true},
		{"Not empty", "not empty", false},
		{"Space", " ", false},
		{"Newline", "\n", false},
		{"Tab", "\t", false},
		{"Hello", "hello", false},
		{"こんにちは", "こんにちは", false},
		{"😊", "😊", false},
		{"\r", "\r", false},
		{"\u200B", "\u200B", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.IsEmpty(tt.input)
			if result != tt.expected {
				t.Errorf("IsEmpty(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsNotEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty string", "", false},
		{"Not empty", "not empty", true},
		{"Space", " ", true},
		{"Newline", "\n", true},
		{"Tab", "\t", true},
		{"Hello", "hello", true},
		{"こんにちは", "こんにちは", true},
		{"😊", "😊", true},
		{"\r", "\r", true},
		{"\u200B", "\u200B", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.IsNotEmpty(tt.input)
			if result != tt.expected {
				t.Errorf("IsNotEmpty(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty string", "", true},
		{"Only spaces", "   ", true},
		{"Only whitespace", "\n\t\r", true},
		{"Not blank", " not blank ", false},
		{"Not blank - lowercase", "hello", false},
		{"Not blank - Japanese", "こんにちは", false},
		{"Not blank - Emoji", "😊", false},
		{"Not blank - mixed whitespace", " \n hello \t ", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.IsBlank(tt.input)
			if result != tt.expected {
				t.Errorf("IsBlank(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsNotBlank(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty string", "", false},
		{"Only spaces", "   ", false},
		{"Only whitespace", "\n\t\r", false},
		{"Not blank", " not blank ", true},
		{"Not blank - lowercase", "hello", true},
		{"Not blank - Japanese", "こんにちは", true},
		{"Not blank - Emoji", "😊", true},
		{"Not blank - mixed whitespace", " \n hello \t ", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.IsNotBlank(tt.input)
			if result != tt.expected {
				t.Errorf("IsNotBlank(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Not blank - lowercase", "hello", true},
		{"Not blank - Mixcase", "HelloWorld", true},
		{"Not blank - alphanumeric", "hello123", false},
		{"Not blank - digits only", "123", false},
		{"Empty string", "", false},
		{"Not blank - Japanese", "こんにちは", true},
		{"Not blank - Emoji", "😊", false},
	}
	for _, tt := range tests {
		result := lxstrings.IsAlpha(tt.input)
		if result != tt.expected {
			t.Errorf("IsAlpha(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Only numberic", "12345", true},
		{"Leading zeros", "00123", true},
		{"Mixed alphanumeric", "123abc", false},
		{"Only letters", "abc", false},
		{"Empty string", "", false},
		{"Full-width digits", "１２３４５", true}, // Full-width digits
		{"Emoji with numbers", "😊123", false},
	}
	for _, tt := range tests {
		result := lxstrings.IsNumeric(tt.input)
		if result != tt.expected {
			t.Errorf("IsNumeric(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Only alphanumeric", "HelloWorld", true},
		{"Only numbers", "12345", true},
		{"Mixed alphanumeric", "hello!", false},
		{"Empty string", "", false},
		{"Japanese with numbers", "こんにちは123", true},
		{"Emoji with numbers", "😊emoji123", false},
	}
	for _, tt := range tests {
		result := lxstrings.IsAlphaNumeric(tt.input)
		if result != tt.expected {
			t.Errorf("IsAlphaNumeric(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsSpace(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"Empty string", "", false},
		{"Single space", " ", true},
		{"Multiple spaces", "   ", true},
		{"Tab", "\t", true},
		{"Newline", "\n", true},
		{"Mixed whitespace", " \t\n", true},
		{"Unicode space", "\u00A0", true}, // non-breaking space
		{"Text only", "abc", false},
		{"Text with space", "a b", false},
		{"Leading space", " abc", false},
		{"Trailing space", "abc ", false},
		{"Unicode text", "xin chào", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxstrings.IsSpace(tt.input)
			if got != tt.want {
				t.Errorf("IsSpace(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
func TestEquals(t *testing.T) {
	tests := []struct {
		name     string
		s1, s2   string
		expected bool
	}{
		{"Equal", "hello", "hello", true},
		{"Different case", "hello", "Hello", false},
		{"Both empty", "", "", true},
		{"One empty", "hello", "", false},
		{"Japanese", "こんにちは", "こんにちは", true},
		{"Emoji", "😊", "😊", true},
		{"Equal but extra space", "test", "test ", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.Equals(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("Equals(%q, %q) = %v; want %v", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}

func TestNotEquals(t *testing.T) {
	tests := []struct {
		name     string
		s1, s2   string
		expected bool
	}{
		{"Equal", "hello", "hello", false},
		{"Different case", "hello", "Hello", true},
		{"Both empty", "", "", false},
		{"One empty", "hello", "", true},
		{"Japanese", "こんにちは", "こんにちは", false},
		{"Emoji", "😊", "😊", false},
		{"Equal but extra space", "test", "test ", true},
	}
	for _, tt := range tests {
		result := lxstrings.NotEquals(tt.s1, tt.s2)
		if result != tt.expected {
			t.Errorf("NotEquals(%q, %q) = %v; want %v", tt.s1, tt.s2, result, tt.expected)
		}
	}
}

func TestEqualsIgnoreCase(t *testing.T) {
	tests := []struct {
		name     string
		s1, s2   string
		expected bool
	}{
		{"Equal", "hello", "HELLO", true},
		{"GoLang", "golang", "GOLANG", true},
		{"Test", "test", "TEST", true},
		{"こんにちは", "こんにちは", "こんにちは", true},
		{"😊", "😊", "😊", true},
		{"NotEqual extra space", "Different ", "Different", false},
	}
	for _, tt := range tests {
		result := lxstrings.EqualsIgnoreCase(tt.s1, tt.s2)
		if result != tt.expected {
			t.Errorf("EqualsIgnoreCase(%q, %q) = %v; want %v", tt.s1, tt.s2, result, tt.expected)
		}
	}
}

func TestNotEqualsIgnoreCase(t *testing.T) {
	tests := []struct {
		name     string
		s1, s2   string
		expected bool
	}{
		{"Equal", "hello", "HELLO", false},
		{"GoLang", "golang", "GOLANG", false},
		{"Test", "test", "TEST", false},
		{"こんにちは", "こんにちは", "こんにちは", false},
		{"😊", "😊", "😊", false},
		{"NotEqual extra space", "Different ", "Different", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.NotEqualsIgnoreCase(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("NotEqualsIgnoreCase(%q, %q) = %v; want %v", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	tests := []struct {
		name     string
		str, sub string
		expected int
	}{
		{"Match the second word", "hello world", "world", 6},
		{"Case insensitive", "hello world", "WORLD", -1},
		{"Match the first word", "golang", "go", 0},
		{"Match the last word", "golang", "lang", 2},
		{"Case insensitive", "test", "TEST", -1},
		{"Both empty", "", "", 0},
		{"Empty substring", "non-empty", "", 0},
		{"Empty string", "", "non-empty", -1},
		{"Japanese", "こんにちは世界", "世界", 15},
		{"Emoji", "😊emoji😊", "emoji", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxstrings.Index(tt.str, tt.sub)
			if got != tt.expected {
				t.Errorf(
					"Index(%q, %q) = %d; want %d",
					tt.str, tt.sub, got, tt.expected,
				)
			}
		})
	}
}

func TestIndexFrom(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		sub      string
		fromInd  int
		expected int
	}{
		{
			name:     "found after index",
			str:      "hello world",
			sub:      "world",
			fromInd:  0,
			expected: 6,
		},
		{
			name:     "found after middle",
			str:      "hello world world",
			sub:      "world",
			fromInd:  7,
			expected: 12,
		},
		{
			name:     "not found",
			str:      "hello world",
			sub:      "abc",
			fromInd:  0,
			expected: -1,
		},
		{
			name:     "from index beyond match",
			str:      "hello world",
			sub:      "hello",
			fromInd:  1,
			expected: -1,
		},
		{
			name:     "from index at exact match",
			str:      "hello world",
			sub:      "world",
			fromInd:  6,
			expected: 6,
		},
		{
			name:     "empty substring",
			str:      "abc",
			sub:      "",
			fromInd:  1,
			expected: 1,
		},
		{
			name:     "fromInd equals len(str)",
			str:      "abc",
			sub:      "a",
			fromInd:  3,
			expected: -1,
		},
		{
			name:     "fromInd greater than len(str)",
			str:      "abc",
			sub:      "a",
			fromInd:  4,
			expected: -1,
		},
		{
			name:     "fromInd is negative",
			str:      "abc",
			sub:      "a",
			fromInd:  -1,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxstrings.IndexFrom(tt.str, tt.sub, tt.fromInd)
			if got != tt.expected {
				t.Errorf(
					"IndexFrom(%q, %q, %d) = %d; want %d",
					tt.str, tt.sub, tt.fromInd, got, tt.expected,
				)
			}
		})
	}
}

func TestIndexIgnoreCase(t *testing.T) {
	{
		tests := []struct {
			s, substr string
			expected  int
		}{
			{"hello world", "WORLD", 6},
			{"GoLang", "golang", 0},
			{"TestString", "teststring", 0},
			{"CaseSensitive", "casesensitive", 0},
			{"NoMatch", "MATCH", 2},
			{"", "", 0},
			{"non-empty", "", 0},
			{"", "non-empty", -1},
			{"こんにちは世界", "世界", 15},
			{"😊Emoji😊", "emoji", 4},
		}
		for _, test := range tests {
			result := lxstrings.IndexIgnoreCase(test.s, test.substr)
			if result != test.expected {
				t.Errorf("IndexIgnoreCase(%q, %q) = %d; want %d", test.s, test.substr, result, test.expected)
			}
		}
	}
}

func TestLastIndex(t *testing.T) {
	tests := []struct {
		s, substr string
		expected  int
	}{
		{"hello world world", "world", 12},
		{"hello world", "WORLD", -1},
		{"golang golang", "go", 7},
		{"test test test", "test", 10},
		{"", "", 0},
		{"non-empty", "", 9},
		{"", "non-empty", -1},
		{"こんにちは世界こんにちは", "こんにちは", 21},
		{"😊emoji😊emoji😊", "emoji", 13},
	}
	for _, test := range tests {
		result := lxstrings.LastIndex(test.s, test.substr)
		if result != test.expected {
			t.Errorf("LastIndex(%q, %q) = %d; want %d", test.s, test.substr, result, test.expected)
		}
	}
}

func TestLastIndexIgnoreCase(t *testing.T) {
	{
		tests := []struct {
			s, substr string
			expected  int
		}{
			{"hello world WORLD", "WORLD", 12},
			{"GoLang goLANG", "golang", 7},
			{"TestString teststring TESTSTRING", "teststring", 22},
			{"CaseSensitive CASESENSITIVE", "casesensitive", 14},
			{"NoMatch MATCH NOMATCH", "MATCH", 16},
			{"", "", 0},
			{"non-empty", "", 9},
			{"", "non-empty", -1},
			{"こんにちは世界こんにちは", "こんにちは", 21},
			{"😊Emoji😊emoji😊", "emoji", 13},
		}
		for _, test := range tests {
			result := lxstrings.LastIndexIgnoreCase(test.s, test.substr)
			if result != test.expected {
				t.Errorf("LastIndexIgnoreCase(%q, %q) = %d; want %d", test.s, test.substr, result, test.expected)
			}
		}
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 5},
		{"", 0},
		{"こんにちは", 15}, // Each Japanese character is 3 bytes in UTF-8
		{"😊emoji", 9}, // Emoji is 4 bytes, 'emoji' is 5 bytes
		{" ", 1},
		{"\n", 1},
		{"\t", 1},
	}
	for _, test := range tests {
		result := lxstrings.Length(test.input)
		if result != test.expected {
			t.Errorf("Length(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestLowerCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HELLO", "hello"},
		{"Hello World", "hello world"},
		{"golang", "golang"},
		{"", ""},
		{"123ABC", "123abc"},
		{"こんにちは", "こんにちは"},
		{"😊EMOJI", "😊emoji"},
	}
	for _, test := range tests {
		result := lxstrings.LowerCase(test.input)
		if result != test.expected {
			t.Errorf("LowerCase(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestUpperCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "HELLO"},
		{"Hello World", "HELLO WORLD"},
		{"GOLANG", "GOLANG"},
		{"", ""},
		{"123abc", "123ABC"},
		{"こんにちは", "こんにちは"},
		{"😊emoji", "😊EMOJI"},
	}
	for _, test := range tests {
		result := lxstrings.UpperCase(test.input)
		if result != test.expected {
			t.Errorf("UpperCase(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestTrimSpace(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"   hello world   ", "hello world"},
		{"\n\tgolang\r\n", "golang"},
		{" no leading or trailing spaces ", "no leading or trailing spaces"},
		{"\u200Bzero width space\u200B", "\u200Bzero width space\u200B"},
		{"😊 emoji 😊", "😊 emoji 😊"},
		{"      ", ""},
		{"", ""},
	}
	for _, test := range tests {
		result := lxstrings.TrimSpace(test.input)
		if result != test.expected {
			t.Errorf("TrimSpace(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestTrim(t *testing.T) {
	tests := []struct {
		input    string
		cutset   string
		expected string
	}{
		{"---hello---", "-", "hello"},
		{"***golang***", "*", "golang"},
		{"   spaced   ", " ", "spaced"},
		{"!!!exciting!!!", "!", "exciting"},
		{"😊emoji😊", "😊", "emoji"},
		{"no trim needed", "xyz", "no trim needed"},
		{"aaaaaa", "a", ""},
		{"", "abc", ""},
	}
	for _, test := range tests {
		result := lxstrings.Trim(test.input, test.cutset)
		if result != test.expected {
			t.Errorf("Trim(%q, %q) = %q; want %q", test.input, test.cutset, result, test.expected)
		}
	}
}

func TestTrimLeft(t *testing.T) {
	tests := []struct {
		input    string
		cutset   string
		expected string
	}{
		{"---hello---", "-", "hello---"},
		{"***golang***", "*", "golang***"},
		{"   spaced   ", " ", "spaced   "},
		{"!!!exciting!!!", "!", "exciting!!!"},
		{"😊emoji😊", "😊", "emoji😊"},
		{"no trim needed", "xyz", "no trim needed"},
		{"aaaaaa", "a", ""},
		{"", "abc", ""},
	}
	for _, test := range tests {
		result := lxstrings.TrimLeft(test.input, test.cutset)
		if result != test.expected {
			t.Errorf("TrimLeft(%q, %q) = %q; want %q", test.input, test.cutset, result, test.expected)
		}
	}
}

func TestTrimRight(t *testing.T) {
	tests := []struct {
		input    string
		cutset   string
		expected string
	}{
		{"---hello---", "-", "---hello"},
		{"***golang***", "*", "***golang"},
		{"   spaced   ", " ", "   spaced"},
		{"!!!exciting!!!", "!", "!!!exciting"},
		{"😊emoji😊", "😊", "😊emoji"},
		{"no trim needed", "xyz", "no trim needed"},
		{"aaaaaa", "a", ""},
		{"", "abc", ""},
	}
	for _, test := range tests {
		result := lxstrings.TrimRight(test.input, test.cutset)
		if result != test.expected {
			t.Errorf("TrimRight(%q, %q) = %q; want %q", test.input, test.cutset, result, test.expected)
		}
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		input    string
		maxWidth int
		expected string
	}{
		{"Hello, World!", 5, "Hello"},
		{"Hello", 10, "Hello"},
		{"GoLang", 3, "GoL"},
		{"Short", 0, ""},
		{"Exact", 5, "Exact"},
		{"This is a longer string", 8, "This is "},
	}
	for _, test := range tests {
		result := lxstrings.Truncate(test.input, test.maxWidth)
		if result != test.expected {
			t.Errorf("Truncate(%q, %d) = %q; want %q", test.input, test.maxWidth, result, test.expected)
		}
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		input    string
		sep      string
		expected []string
	}{
		{"a,b,c", ",", []string{"a", "b", "c"}},
		{"hello world", " ", []string{"hello", "world"}},
		{"one;two;three", ";", []string{"one", "two", "three"}},
		{"no separator", ",", []string{"no separator"}},
		{"", ",", []string{""}},
		{"a--b--c", "--", []string{"a", "b", "c"}},
		{"😊-emoji-😊", "-", []string{"😊", "emoji", "😊"}},
	}
	for _, test := range tests {
		result := lxstrings.Split(test.input, test.sep)
		if len(result) != len(test.expected) {
			t.Errorf("Split(%q, %q) = %v; want %v", test.input, test.sep, result, test.expected)
			continue
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("Split(%q, %q) = %v; want %v", test.input, test.sep, result, test.expected)
				break
			}
		}
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		input    []string
		sep      string
		expected string
	}{
		{[]string{"a", "b", "c"}, ",", "a,b,c"},
		{[]string{"hello", "world"}, " ", "hello world"},
		{[]string{"one", "two", "three"}, ";", "one;two;three"},
		{[]string{"no", "separator"}, "", "noseparator"},
		{[]string{}, ",", ""},
		{[]string{"😊", "emoji", "😊"}, "-", "😊-emoji-😊"},
	}
	for _, test := range tests {
		result := lxstrings.Join(test.input, test.sep)
		if result != test.expected {
			t.Errorf("Join(%v, %q) = %q; want %q", test.input, test.sep, result, test.expected)
		}
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		input    string
		count    int
		expected string
	}{
		{"ha", 3, "hahaha"},
		{"go", 0, ""},
		{"!", 5, "!!!!!"},
		{"😊", 4, "😊😊😊😊"},
		{"abc", 1, "abc"},
		{"", 10, ""},
	}
	for _, test := range tests {
		result := lxstrings.Repeat(test.input, test.count)
		if result != test.expected {
			t.Errorf("Repeat(%q, %d) = %q; want %q", test.input, test.count, result, test.expected)
		}
	}
}

func TestStartBy(t *testing.T) {
	tests := []struct {
		s        string
		prefix   string
		expected bool
	}{
		{"hello world", "hello", true},
		{"hello world", "world", false},
		{"golang", "go", true},
		{"test", "TEST", false},
		{"", "", true},
		{"non-empty", "", true},
		{"", "non-empty", false},
		{"こんにちは世界", "こんにちは", true},
		{"😊emoji😊", "😊", true},
	}
	for _, test := range tests {
		result := lxstrings.StartBy(test.s, test.prefix)
		if result != test.expected {
			t.Errorf("StartBy(%q, %q) = %v; want %v", test.s, test.prefix, result, test.expected)
		}
	}
}

func TestStartByIgnoreCase(t *testing.T) {
	tests := []struct {
		s        string
		prefix   string
		expected bool
	}{
		{"hello world", "HELLO", true},
		{"GoLang", "golang", true},
		{"TestString", "teststring", true},
		{"CaseSensitive", "casesensitive", true},
		{"NoMatch", "MATCH", false},
		{"", "", true},
		{"non-empty", "", true},
		{"", "non-empty", false},
		{"こんにちは世界", "こんにちは", true},
		{"😊Emoji😊", "emoji", false},
	}
	for _, test := range tests {
		result := lxstrings.StartByIgnoreCase(test.s, test.prefix)
		if result != test.expected {
			t.Errorf("StartByIgnoreCase(%q, %q) = %v; want %v", test.s, test.prefix, result, test.expected)
		}
	}
}

func TestStartByAny(t *testing.T) {
	tests := []struct {
		s        string
		prefixes []string
		expected bool
	}{
		{"hello world", []string{"hi", "hello"}, true},
		{"hello world", []string{"world", "planet"}, false},
		{"golang", []string{"go", "lang"}, true},
		{"test", []string{"TEST", "exam"}, false},
		{"", []string{""}, true},
		{"non-empty", []string{}, false},
		{"こんにちは世界", []string{"こんにちは", "こんばんは"}, true},
		{"😊emoji😊", []string{"😊", "😢"}, true},
	}
	for _, test := range tests {
		result := lxstrings.StartByAny(test.s, test.prefixes...)
		if result != test.expected {
			t.Errorf("StartByAny(%q, %v) = %v; want %v", test.s, test.prefixes, result, test.expected)
		}
	}
}

func TestStartByAnyIgnoreCase(t *testing.T) {
	tests := []struct {
		s        string
		prefixes []string
		expected bool
	}{
		{"hello world", []string{"HI", "HELLO"}, true},
		{"GoLang", []string{"WORLD", "PLANET"}, false},
		{"TestString", []string{"teststring", "exam"}, true},
		{"CaseSensitive", []string{"CASESENSITIVE", "other"}, true},
		{"NoMatch", []string{"MATCH", "different"}, false},
		{"", []string{""}, true},
		{"non-empty", []string{}, false},
		{"こんにちは世界", []string{"こんにちは", "こんばんは"}, true},
		{"😊Emoji😊", []string{"emoji", "😢"}, false},
	}
	for _, test := range tests {
		result := lxstrings.StartByAnyIgnoreCase(test.s, test.prefixes...)
		if result != test.expected {
			t.Errorf("StartByAnyIgnoreCase(%q, %v) = %v; want %v", test.s, test.prefixes, result, test.expected)
		}
	}
}

func TestEndBy(t *testing.T) {
	tests := []struct {
		s        string
		suffix   string
		expected bool
	}{
		{"hello world", "world", true},
		{"hello world", "hello", false},
		{"golang", "lang", true},
		{"test", "TEST", false},
		{"", "", true},
		{"non-empty", "", true},
		{"", "non-empty", false},
		{"こんにちは世界", "世界", true},
		{"😊emoji😊", "😊", true},
	}
	for _, test := range tests {
		result := lxstrings.EndBy(test.s, test.suffix)
		if result != test.expected {
			t.Errorf("EndBy(%q, %q) = %v; want %v", test.s, test.suffix, result, test.expected)
		}
	}
}

func TestEndByIgnoreCase(t *testing.T) {
	tests := []struct {
		s        string
		suffix   string
		expected bool
	}{
		{"hello world", "WORLD", true},
		{"GoLang", "golang", true},
		{"TestString", "teststring", true},
		{"CaseSensitive", "casesensitive", true},
		{"NoMatch", "MATCH", true},
		{"", "", true},
		{"non-empty", "", true},
		{"", "non-empty", false},
		{"こんにちは世界", "世界", true},
		{"😊Emoji😊", "emoji", false},
	}
	for _, test := range tests {
		result := lxstrings.EndByIgnoreCase(test.s, test.suffix)
		if result != test.expected {
			t.Errorf("EndByIgnoreCase(%q, %q) = %v; want %v", test.s, test.suffix, result, test.expected)
		}
	}
}

func TestEndByAny(t *testing.T) {
	tests := []struct {
		s        string
		suffixes []string
		expected bool
	}{
		{"hello world", []string{"planet", "world"}, true},
		{"hello world", []string{"hello", "hi"}, false},
		{"golang", []string{"lang", "go"}, true},
		{"test", []string{"TEST", "exam"}, false},
		{"", []string{""}, true},
		{"non-empty", []string{}, false},
		{"こんにちは世界", []string{"こんばんは", "世界"}, true},
		{"😊emoji😊", []string{"😢", "😊"}, true},
	}
	for _, test := range tests {
		result := lxstrings.EndByAny(test.s, test.suffixes...)
		if result != test.expected {
			t.Errorf("EndByAny(%q, %v) = %v; want %v", test.s, test.suffixes, result, test.expected)
		}
	}
}

func TestEndByAnyIgnoreCase(t *testing.T) {
	tests := []struct {
		s        string
		suffixes []string
		expected bool
	}{
		{"hello world", []string{"PLANET", "WORLD"}, true},
		{"GoLang", []string{"HELLO", "HI"}, false},
		{"TestString", []string{"teststring", "exam"}, true},
		{"CaseSensitive", []string{"CASESENSITIVE", "other"}, true},
		{"NoMatch", []string{"MATCH", "different"}, true},
		{"", []string{""}, true},
		{"non-empty", []string{}, false},
		{"こんにちは世界", []string{"こんばんは", "世界"}, true},
		{"😊Emoji😊", []string{"emoji", "😢"}, false},
	}
	for _, test := range tests {
		result := lxstrings.EndByAnyIgnoreCase(test.s, test.suffixes...)
		if result != test.expected {
			t.Errorf("EndByAnyIgnoreCase(%q, %v) = %v; want %v", test.s, test.suffixes, result, test.expected)
		}
	}
}

func TestReplace(t *testing.T) {
	tests := []struct {
		s        string
		old      string
		new      string
		n        int
		expected string
	}{
		{"hello world", "world", "gopher", -1, "hello gopher"},
		{"golang golang", "go", "GO", 1, "GOlang golang"},
		{"test test test", "test", "exam", 2, "exam exam test"},
		{"no match here", "xyz", "abc", -1, "no match here"},
		{"", "anything", "something", -1, ""},
		{"😊emoji😊emoji😊", "emoji", "EMOJI", 2, "😊EMOJI😊EMOJI😊"},
	}
	for _, test := range tests {
		result := lxstrings.Replace(test.s, test.old, test.new, test.n)
		if result != test.expected {
			t.Errorf("Replace(%q, %q, %q, %d) = %q; want %q", test.s, test.old, test.new, test.n, result, test.expected)
		}
	}
}

func TestReplaceAll(t *testing.T) {
	tests := []struct {
		s        string
		old      string
		new      string
		expected string
	}{
		{"hello world", "world", "gopher", "hello gopher"},
		{"golang golang", "go", "GO", "GOlang GOlang"},
		{"test test test", "test", "exam", "exam exam exam"},
		{"no match here", "xyz", "abc", "no match here"},
		{"", "anything", "something", ""},
		{"😊emoji😊emoji😊", "emoji", "EMOJI", "😊EMOJI😊EMOJI😊"},
	}
	for _, test := range tests {
		result := lxstrings.ReplaceAll(test.s, test.old, test.new)
		if result != test.expected {
			t.Errorf("ReplaceAll(%q, %q, %q) = %q; want %q", test.s, test.old, test.new, result, test.expected)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		s        string
		substr   string
		expected string
	}{
		{"hello world", "world", "hello "},
		{"golang golang", "go", "lang lang"},
		{"test test test", "test", "  "},
		{"no match here", "xyz", "no match here"},
		{"", "anything", ""},
		{"😊emoji😊emoji😊", "emoji", "😊😊😊"},
	}
	for _, test := range tests {
		result := lxstrings.Remove(test.s, test.substr)
		if result != test.expected {
			t.Errorf("Remove(%q, %q) = %q; want %q", test.s, test.substr, result, test.expected)
		}
	}
}

func TestRemoveIgnoreCase(t *testing.T) {
	tests := []struct {
		s        string
		substr   string
		expected string
	}{
		{"hello WORLD", "world", "hello "},
		{"GoLang golang", "GOLANG", " "},
		{"Test test TEST", "test", "  "},
		{"no match here", "XYZ", "no match here"},
		{"", "ANYTHING", ""},
		{"😊Emoji😊emoji😊", "EMOJI", "😊😊😊"},
	}
	for _, test := range tests {
		result := lxstrings.RemoveIgnoreCase(test.s, test.substr)
		if result != test.expected {
			t.Errorf("RemoveIgnoreCase(%q, %q) = %q; want %q", test.s, test.substr, result, test.expected)
		}
	}
}

func TestRemoveAny(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		substrs  []string
		expected string
	}{
		{"Check happy case", "hello world", []string{"hello", "world"}, " "},
		{"Check split word", "golang golang", []string{"go", "lang"}, " "},
		{"Check test repeat", "test test test", []string{"test", "exam"}, "  "},
		{"Check no match here", "no match here", []string{"xyz", "abc"}, "no match here"},
		{"Check empty string", "", []string{"anything", "something"}, ""},
		{"Check emoji string", "😊emoji😊emoji😊", []string{"emoji", "EMOJI"}, "😊😊😊"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.RemoveAny(tt.str, tt.substrs...)
			if result != tt.expected {
				t.Errorf("RemoveAny(%q, %v) = %q; want %q", tt.str, tt.substrs, result, tt.expected)
			}
		})
	}
}

func TestRemoveAnyIgnoreCase(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		substrs  []string
		expected string
	}{
		{"Check happy case", "hello WORLD", []string{"HELLO", "world"}, " "},
		{"Check GoLang golang", "GoLang golang", []string{"GOLANG", "go"}, " "},
		{"Check Test test TEST", "Test test TEST", []string{"TEST", "test"}, "  "},
		{"Check no match here", "no match here", []string{"XYZ", "ABC"}, "no match here"},
		{"Check empty string", "", []string{"ANYTHING", "SOMETHING"}, ""},
		{"Check emoji string", "😊Emoji😊emoji😊", []string{"EMOJI", "emoji"}, "😊😊😊"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.RemoveAnyIgnoreCase(tt.str, tt.substrs...)
			if result != tt.expected {
				t.Errorf("RemoveAnyIgnoreCase(%q, %v) = %q; want %q", tt.str, tt.substrs, result, tt.expected)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Check happy case", "hello", "olleh"},
		{"Check empty string", "", ""},
		{"Check golang", "golang", "gnalog"},
		{"Check emoji", "😊emoji", "ijome😊"},
		{"Check japanese", "こんにちは", "はちにんこ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSubString(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		start    int
		end      int
		expected string
	}{
		{"Check happy case", "hello world", 0, 5, "hello"},
		{"Check golang", "golang", 3, 6, "ang"},
		{"Check test string", "test string", 5, 11, "string"},
		{"Check emoji", "😊emoji😊", 1, 6, "emoji"},
		{"Check short", "short", 2, 10, "ort"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.SubString(tt.str, tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("SubString(%q, %d, %d) = %q; want %q", tt.str, tt.start, tt.end, result, tt.expected)
			}
		})
	}
}

func TestSubStringBefore(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		substr   string
		expected string
	}{
		{"Check happy case", "hello world", "world", "hello "},
		{"Check match second word", "golang programming", "programming", "golang "},
		{"Check middle word", "test string example", "string", "test "},
		{"Check no match", "no match here", "xyz", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.SubStringBefore(tt.str, tt.substr)
			if result != tt.expected {
				t.Errorf("SubStringBefore(%q, %q) = %q; want %q", tt.str, tt.substr, result, tt.expected)
			}
		})
	}
}
func TestSubStringBeforeIgnoreCase(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		substr   string
		expected string
	}{
		{"Check lowercase", "hello WORLD", "world", "hello "},
		{"Check uppercase", "GoLang PROGRAMMING", "programming", "GoLang "},
		{"Check middle uppercase", "Test STRING example", "STRING", "Test "},
		{"Check no match", "no match here", "XYZ", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.SubStringBeforeIgnoreCase(tt.str, tt.substr)
			if result != tt.expected {
				t.Errorf("SubStringBeforeIgnoreCase(%q, %q) = %q; want %q", tt.str, tt.substr, result, tt.expected)
			}
		})
	}
}

func TestSubStringAfter(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		substr   string
		expected string
	}{
		{"Check hello world", "hello world", "hello", " world"},
		{"Check golang programming", "golang programming", "golang", " programming"},
		{"Check test string example", "test string example", "string", " example"},
		{"Check no match", "no match here", "xyz", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.SubStringAfter(tt.str, tt.substr)
			if result != tt.expected {
				t.Errorf("SubStringAfter(%q, %q) = %q; want %q", tt.str, tt.substr, result, tt.expected)
			}
		})
	}
}

func TestSubStringAfterIgnoreCase(t *testing.T) {
	{
		tests := []struct {
			name     string
			str      string
			substr   string
			expected string
		}{
			{"Check uppercase", "hello WORLD", "HELLO", " WORLD"},
			{"Check mixcase", "GoLang PROGRAMMING", "GOLANG", " PROGRAMMING"},
			{"Check middle uppercase", "Test STRING example", "STRING", " example"},
			{"Check no match", "no match here", "XYZ", ""},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := lxstrings.SubStringAfterIgnoreCase(tt.str, tt.substr)
				if result != tt.expected {
					t.Errorf("SubStringAfterIgnoreCase(%q, %q) = %q; want %q", tt.str, tt.substr, result, tt.expected)
				}
			})
		}
	}
}

func TestPadLeft(t *testing.T) {
	{
		tests := []struct {
			name     string
			str      string
			length   int
			padChar  string
			expected string
		}{
			{"Test with char *", "hello", 10, "*", "*****hello"},
			{"Test with char 0", "golang", 8, "0", "00golang"},
			{"Test with char -", "test", 6, "-", "--test"},
			{"Test with char x, but not enough padding", "short", 3, "x", "short"},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := lxstrings.PadLeft(tt.str, tt.length, tt.padChar)
				if result != tt.expected {
					t.Errorf("PadLeft(%q, %d, %q) = %q; want %q", tt.str, tt.length, tt.padChar, result, tt.expected)
				}
			})
		}
	}
}

func TestPadRight(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		length   int
		padChar  string
		expected string
	}{
		{"Test with char *", "hello", 10, "*", "hello*****"},
		{"Test with char 0", "golang", 8, "0", "golang00"},
		{"Test with char -", "test", 6, "-", "test--"},
		{"Test with char x, but not enough padding", "short", 3, "x", "short"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.PadRight(tt.str, tt.length, tt.padChar)
			if result != tt.expected {
				t.Errorf("PadRight(%q, %d, %q) = %q; want %q", tt.str, tt.length, tt.padChar, result, tt.expected)
			}
		})
	}
}

func TestPadCenter(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		length   int
		padChar  string
		expected string
	}{
		{"Test with char *", "hello", 11, "*", "***hello***"},
		{"Test with char 0", "golang", 10, "0", "00golang00"},
		{"Test with char -", "test", 8, "-", "--test--"},
		{"Test with char x, but not enough padding", "short", 3, "x", "short"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.PadCenter(tt.str, tt.length, tt.padChar)
			if result != tt.expected {
				t.Errorf("PadCenter(%q, %d, %q) = %q; want %q", tt.str, tt.length, tt.padChar, result, tt.expected)
			}
		})
	}
}

func TestCountMatches(t *testing.T) {
	tests := []struct {
		name string
		str  string
		sub  string
		want int
	}{
		{
			name: "single match",
			str:  "hello world",
			sub:  "world",
			want: 1,
		},
		{
			name: "multiple matches",
			str:  "go go go",
			sub:  "go",
			want: 3,
		},
		{
			name: "no match",
			str:  "hello",
			sub:  "abc",
			want: 0,
		},
		{
			name: "overlapping patterns",
			str:  "aaaa",
			sub:  "aa",
			want: 2,
		},
		{
			name: "substring longer than string",
			str:  "hi",
			sub:  "hello",
			want: 0,
		},
		{
			name: "empty string",
			str:  "",
			sub:  "a",
			want: 0,
		},
		{
			name: "empty substring",
			str:  "abc",
			sub:  "",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxstrings.CountMatches(tt.str, tt.sub)
			if got != tt.want {
				t.Errorf(
					"CountMatches(%q, %q) = %d; want %d",
					tt.str, tt.sub, got, tt.want,
				)
			}
		})
	}
}

func TestDefaultIfEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		def      string
		expected string
	}{
		{"Non-empty string", "hello", "default", "hello"},
		{"Empty string", "", "default", "default"},
		{"Whitespace string", "   ", "default", "   "},
		{"Unicode string", "こんにちは", "default", "こんにちは"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.DefaultIfEmpty(tt.input, tt.def)
			if result != tt.expected {
				t.Errorf("DefaultIfEmpty(%q, %q) = %q; want %q", tt.input, tt.def, result, tt.expected)
			}
		})
	}
}

func TestDefaultIfBlank(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		def      string
		expected string
	}{
		{"Non-empty string", "hello", "default", "hello"},
		{"Empty string", "", "default", "default"},
		{"Whitespace string", "   ", "default", "default"},
		{"Unicode string", "こんにちは", "default", "こんにちは"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.DefaultIfBlank(tt.input, tt.def)
			if result != tt.expected {
				t.Errorf("DefaultIfBlank(%q, %q) = %q; want %q", tt.input, tt.def, result, tt.expected)
			}
		})
	}
}

func TestStartWith(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		prefix   string
		expected bool
	}{
		{"Starts with prefix", "hello world", "hello", true},
		{"Does not start with prefix", "hello world", "world", false},
		{"Empty prefix", "hello world", "", true},
		{"Empty string and empty prefix", "", "", true},
		{"Empty string and non-empty prefix", "", "hello", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.StartWith(tt.input, tt.prefix)
			if result != tt.expected {
				t.Errorf("StartWith(%q, %q) = %v; want %v", tt.input, tt.prefix, result, tt.expected)
			}
		})
	}
}

func TestStartWithIgnoreCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		prefix   string
		expected bool
	}{
		{"Starts with prefix (case insensitive)", "Hello World", "hello", true},
		{"Does not start with prefix (case insensitive)", "Hello World", "WORLD", false},
		{"Empty prefix", "Hello World", "", true},
		{"Empty string and empty prefix", "", "", true},
		{"Empty string and non-empty prefix", "", "hello", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.StartWithIgnoreCase(tt.input, tt.prefix)
			if result != tt.expected {
				t.Errorf("StartWithIgnoreCase(%q, %q) = %v; want %v", tt.input, tt.prefix, result, tt.expected)
			}
		})
	}
}

func TestStartWithAny(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		prefixes []string
		expected bool
	}{
		{"Starts with one of the prefixes", "hello world", []string{"hi", "hello"}, true},
		{"Does not start with any of the prefixes", "hello world", []string{"world", "planet"}, false},
		{"Empty prefixes", "hello world", []string{}, false},
		{"Empty string and empty prefixes", "", []string{}, false},
		{"Empty string and non-empty prefixes", "", []string{"hello"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.StartWithAny(tt.input, tt.prefixes...)
			if result != tt.expected {
				t.Errorf("StartWithAny(%q, %v) = %v; want %v", tt.input, tt.prefixes, result, tt.expected)
			}
		})
	}
}

func TestStartWithAnyIgnoreCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		prefixes []string
		expected bool
	}{
		{"Starts with one of the prefixes (case insensitive)", "Hello World", []string{"HI", "hello"}, true},
		{"Does not start with any of the prefixes (case insensitive)", "Hello World", []string{"WORLD", "PLANET"}, false},
		{"Empty prefixes", "Hello World", []string{}, false},
		{"Empty string and empty prefixes", "", []string{}, false},
		{"Empty string and non-empty prefixes", "", []string{"hello"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.StartWithAnyIgnoreCase(tt.input, tt.prefixes...)
			if result != tt.expected {
				t.Errorf("StartWithAnyIgnoreCase(%q, %v) = %v; want %v", tt.input, tt.prefixes, result, tt.expected)
			}
		})
	}
}

func TestEndWith(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		suffix   string
		expected bool
	}{
		{"Ends with suffix", "hello world", "world", true},
		{"Does not end with suffix", "hello world", "hello", false},
		{"Empty suffix", "hello world", "", true},
		{"Empty string and empty suffix", "", "", true},
		{"Empty string and non-empty suffix", "", "world", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.EndWith(tt.input, tt.suffix)
			if result != tt.expected {
				t.Errorf("EndWith(%q, %q) = %v; want %v", tt.input, tt.suffix, result, tt.expected)
			}
		})
	}
}

func TestEndWithIgnoreCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		suffix   string
		expected bool
	}{
		{"Ends with suffix (case insensitive)", "Hello World", "WORLD", true},
		{"Does not end with suffix (case insensitive)", "Hello World", "HELLO", false},
		{"Empty suffix", "Hello World", "", true},
		{"Empty string and empty suffix", "", "", true},
		{"Empty string and non-empty suffix", "", "world", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.EndWithIgnoreCase(tt.input, tt.suffix)
			if result != tt.expected {
				t.Errorf("EndWithIgnoreCase(%q, %q) = %v; want %v", tt.input, tt.suffix, result, tt.expected)
			}
		})
	}
}

func TestEndWithAny(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		suffixes []string
		expected bool
	}{
		{"Ends with one of the suffixes", "hello world", []string{"planet", "world"}, true},
		{"Does not end with any of the suffixes", "hello world", []string{"hello", "hi"}, false},
		{"Empty suffixes", "hello world", []string{}, false},
		{"Empty string and empty suffixes", "", []string{}, false},
		{"Empty string and non-empty suffixes", "", []string{"world"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.EndWithAny(tt.input, tt.suffixes...)
			if result != tt.expected {
				t.Errorf("EndWithAny(%q, %v) = %v; want %v", tt.input, tt.suffixes, result, tt.expected)
			}
		})
	}
}

func TestEndWithAnyIgnoreCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		suffixes []string
		expected bool
	}{
		{"Ends with one of the suffixes (case insensitive)", "Hello World", []string{"PLANET", "world"}, true},
		{"Does not end with any of the suffixes (case insensitive)", "Hello World", []string{"HELLO", "HI"}, false},
		{"Empty suffixes", "Hello World", []string{}, false},
		{"Empty string and empty suffixes", "", []string{}, false},
		{"Empty string and non-empty suffixes", "", []string{"world"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxstrings.EndWithAnyIgnoreCase(tt.input, tt.suffixes...)
			if result != tt.expected {
				t.Errorf("EndWithAnyIgnoreCase(%q, %v) = %v; want %v", tt.input, tt.suffixes, result, tt.expected)
			}
		})
	}
}

func TestRuneCount(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"hello", 5},
		{"こんにちは", 5},
		{"😊emoji", 6}, // emoji + 'emoji' => 1 + 5
		{"\u00A0\u00A0", 2},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := lxstrings.RuneCount(tt.input)
			if got != tt.expected {
				t.Errorf("RuneCount(%q) = %d; want %d", tt.input, got, tt.expected)
			}
		})
	}
}
