package lxenv

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	// Value normalizers
	defaultValueNormalizer = &simpleValueNormalizer{}
	// Env file parsers
	defaultEnvFileParser    = &envFileParser{valueNormalizer: defaultValueNormalizer}
	defaultPropertiesParser = &propertiesFileParser{envFileParser: defaultEnvFileParser}
	defaultYAMLParser       = &yamlFileParser{valueNormalizer: defaultValueNormalizer}
	// Env Loaders
	defaultEnvLoader        = &envLoader{parser: defaultEnvFileParser}
	defaultPropertiesLoader = &envLoader{parser: defaultPropertiesParser}
	defaultYAMLLoader       = &envLoader{parser: defaultYAMLParser}
)

// LoadEnv reads one or more .env files and sets environment variables from them.
// Files are loaded in order — later files override earlier ones.
//
// Example:
//
//	lxenv.LoadEnv(".env")
//	lxenv.LoadEnv(".env", ".env.local")
func LoadEnv(paths ...string) error {
	return defaultEnvLoader.load(paths...)
}

// LoadProperties reads one or more .properties files and sets environment variables from them.
// Files are loaded in order — later files override earlier ones.
//
// Example:
//
//	lxenv.LoadProperties("app.properties")
//	lxenv.LoadProperties("app.properties", "app.local.properties")
func LoadProperties(paths ...string) error {
	return defaultPropertiesLoader.load(paths...)
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
	return defaultYAMLLoader.load(paths...)
}

// envLoader is an interface for parsing environment variables from a file.
type envParser interface {
	parse(io.Reader) (map[string]string, error)
}

// valueNormalizer is an interface for cleaning up parsed values.
type valueNormalizer interface {
	normalize(string) string
}

type envLoader struct {
	parser envParser
}

func (ep *envLoader) load(paths ...string) error {
	for _, path := range paths {
		if err := func() error {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()

			pairs, err := ep.parser.parse(f)
			if err != nil {
				return fmt.Errorf("parse %q: %w", path, err)
			}
			for k, v := range pairs {
				if err := Set(k, v); err != nil {
					return fmt.Errorf("set %q: %w", k, err)
				}
			}
			return nil
		}(); err != nil {
			return err
		}
	}
	return nil
}

// envFileParser is a struct that implements the envParser interface.
// It reads environment variables from .env files.
type envFileParser struct {
	valueNormalizer valueNormalizer
}

// parse parses KEY=VALUE format (.env / .properties).
// - Lines starting with # are comments
// - Blank lines are ignored
// - Values may be quoted with " or '
// - Inline comments after # are stripped (outside quotes)
func (elp *envFileParser) parse(r io.Reader) (map[string]string, error) {
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

		val = elp.valueNormalizer.normalize(val)
		pairs[key] = val
	}

	return pairs, scanner.Err()
}

// yamlFileParser is a struct that implements the Loader interface.
// It reads environment variables from .yml/.yaml files.
type yamlFileParser struct {
	valueNormalizer valueNormalizer
}

// parse parses nested YAML format with unlimited nesting depth.
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
func (yfp *yamlFileParser) parse(r io.Reader) (map[string]string, error) {
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
			pairs[fullKey] = yfp.valueNormalizer.normalize(val)
		}
	}

	return pairs, scanner.Err()
}

// propertiesFileParser is a wrapper around envParser.
type propertiesFileParser struct {
	envFileParser envParser
}

func (pfp *propertiesFileParser) parse(r io.Reader) (map[string]string, error) {
	return pfp.envFileParser.parse(r)
}

// simpleValueNormalizer handles raw string cleanup for parsed values.
type simpleValueNormalizer struct{}

// stripInlineComment removes everything after an unquoted #.
func (*simpleValueNormalizer) stripInlineComment(s string) string {
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
func (*simpleValueNormalizer) unquote(s string) string {
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') ||
			(s[0] == '\'' && s[len(s)-1] == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// normalize strips inline comments, trims surrounding whitespace, then unquotes the value.
// This is the standard pipeline for any parsed scalar value.
func (p *simpleValueNormalizer) normalize(s string) string {
	s = p.stripInlineComment(s)
	s = strings.TrimSpace(s)
	return p.unquote(s)
}
