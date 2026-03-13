package lxenv

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Load reads one or more .env files and sets environment variables from them.
// Files are loaded in order — later files override earlier ones.
//
// Example:
//
//	lxenv.Load(".env")
//	lxenv.Load(".env", ".env.local")
func Load(paths ...string) error {
	for _, p := range paths {
		if err := loadEnvFile(p); err != nil {
			return err
		}
	}
	return nil
}

// LoadProperties reads one or more .properties files and sets environment variables from them.
// Files are loaded in order — later files override earlier ones.
//
// Example:
//
//	lxenv.LoadProperties("app.properties")
//	lxenv.LoadProperties("app.properties", "app.local.properties")
func LoadProperties(paths ...string) error {
	for _, p := range paths {
		if err := loadEnvFile(p); err != nil {
			return err
		}
	}
	return nil
}

// LoadYML reads one or more .yml/.yaml files and sets environment variables from them.
// Nested keys are flattened using dot-notation with unlimited depth, e.g.:
//
//	database:
//	  pool:
//	    size: 10   →  database.pool.size=10
//
// Files are loaded in order — later files override earlier ones.
//
// Example:
//
//	lxenv.LoadYML("config.yml")
//	lxenv.LoadYML("config.yml", "config.local.yml")
func LoadYML(paths ...string) error {
	for _, p := range paths {
		if err := loadYAMLFile(p); err != nil {
			return err
		}
	}
	return nil
}

func loadEnvFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("lxenv: cannot open file %q: %w", filename, err)
	}
	defer f.Close()

	pairs, err := parseEnv(f)
	if err != nil {
		return fmt.Errorf("lxenv: failed to parse %q: %w", filename, err)
	}

	return applyPairs(pairs)
}

func loadYAMLFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("lxenv: cannot open file %q: %w", filename, err)
	}
	defer f.Close()

	pairs, err := parseYAML(f)
	if err != nil {
		return fmt.Errorf("lxenv: failed to parse %q: %w", filename, err)
	}

	return applyPairs(pairs)
}

func applyPairs(pairs map[string]string) error {
	for k, v := range pairs {
		if err := os.Setenv(k, v); err != nil {
			return fmt.Errorf("lxenv: failed to set %q: %w", k, err)
		}
	}
	return nil
}

// parseEnv parses KEY=VALUE format (.env / .properties).
// - Lines starting with # are comments
// - Blank lines are ignored
// - Values may be quoted with " or '
// - Inline comments after # are stripped (outside quotes)
func parseEnv(r io.Reader) (map[string]string, error) {
	pairs := make(map[string]string)
	scanner := bufio.NewScanner(r)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// skip blank lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		idx := strings.IndexByte(line, '=')
		if idx < 0 {
			return nil, fmt.Errorf("line %d: missing '=' in %q", lineNum, line)
		}

		key := strings.TrimSpace(line[:idx])
		val := strings.TrimSpace(line[idx+1:])

		if key == "" {
			return nil, fmt.Errorf("line %d: empty key", lineNum)
		}

		val = stripInlineComment(val)
		val = unquote(val)
		pairs[key] = val
	}

	return pairs, scanner.Err()
}

// parseYAML parses nested YAML format with unlimited nesting depth.
// Nested keys are flattened using dot-notation, e.g.:
//
//	database:
//	  pool:
//	    size: 10   →  database.pool.size=10
//
// Rules:
// - Lines starting with # are comments and are ignored
// - Blank lines are ignored
// - List items starting with - are ignored
func parseYAML(r io.Reader) (map[string]string, error) {
	pairs := make(map[string]string)
	scanner := bufio.NewScanner(r)

	type frame struct {
		indent int
		key    string
	}
	stack := make([]frame, 0, 8)

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		// skip blank lines, comments and list items
		if trimmed == "" || strings.HasPrefix(trimmed, "#") || strings.HasPrefix(trimmed, "-") {
			continue
		}

		// measure indent (tab counts as 2 spaces)
		indent := 0
		for _, ch := range line {
			if ch == ' ' {
				indent++
			} else if ch == '\t' {
				indent += 2
			} else {
				break
			}
		}

		idx := strings.IndexByte(trimmed, ':')
		if idx < 0 {
			continue
		}

		key := strings.TrimSpace(trimmed[:idx])
		val := strings.TrimSpace(trimmed[idx+1:])

		if key == "" {
			continue
		}

		// pop stack frames at same or deeper indent
		for len(stack) > 0 && stack[len(stack)-1].indent >= indent {
			stack = stack[:len(stack)-1]
		}

		// build full dot-notation key
		fullKey := key
		if len(stack) > 0 {
			parts := make([]string, 0, len(stack)+1)
			for _, f := range stack {
				parts = append(parts, f.key)
			}
			parts = append(parts, key)
			fullKey = strings.Join(parts, ".")
		}

		if val == "" {
			// mapping parent — push onto stack only, do NOT emit an env var
			stack = append(stack, frame{indent: indent, key: key})
		} else {
			val = stripInlineComment(val)
			val = unquote(val)
			pairs[fullKey] = val
		}
	}

	return pairs, scanner.Err()
}

// stripInlineComment removes everything after an unquoted #.
func stripInlineComment(s string) string {
	inSingle, inDouble := false, false
	for i, ch := range s {
		switch ch {
		case '\'':
			if !inDouble {
				inSingle = !inSingle
			}
		case '"':
			if !inSingle {
				inDouble = !inDouble
			}
		case '#':
			if !inSingle && !inDouble {
				return strings.TrimSpace(s[:i])
			}
		}
	}
	return s
}

// unquote removes surrounding single or double quotes from a value.
func unquote(s string) string {
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') ||
			(s[0] == '\'' && s[len(s)-1] == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}
