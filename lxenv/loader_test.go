package lxenv_test

import (
	"os"
	"testing"

	"github.com/nthanhhai2909/lx/lxenv"
)

const (
	baseEnv            = "testdata/base.env"
	overrideEnv        = "testdata/override.env"
	baseProperties     = "testdata/base.properties"
	overrideProperties = "testdata/override.properties"
	baseYML            = "testdata/base.yml"
	overrideYML        = "testdata/override.yml"
)

var allBaseEnvKeys = []string{
	"APP_NAME", "APP_VERSION", "APP_ENV", "APP_DEBUG", "APP_PORT",
	"DB_HOST", "DB_PORT", "DB_NAME", "DB_USER", "DB_PASSWORD",
	"GREETING", "SINGLE_QUOTED", "STRIPPED", "TOKEN", "SPECIAL",
	"EMPTY_VAR", "DESCRIPTION",
}

var allBasePropertiesKeys = []string{
	"app.name", "app.version", "app.env", "app.debug", "app.description",
	"server.host", "server.port", "server.context-path",
	"database.host", "database.port", "database.name",
	"database.username", "database.password", "database.driver",
	"cache.type", "cache.host", "cache.port", "cache.password",
	"mail.host", "mail.port", "mail.username", "mail.auth",
	"logging.level", "logging.file",
	"app.greeting", "app.single_quoted", "app.stripped",
	"app.token", "app.secret", "app.empty",
}

var allBaseYMLKeys = []string{
	"app.name", "app.version", "app.env", "app.debug", "app.description",
	"app.greeting", "app.secret", "app.token", "app.empty", "app.stripped",
	"server.host", "server.port", "server.context-path",
	"server.ssl.enabled", "server.ssl.key-store", "server.ssl.key-store-password",
	"database.host", "database.port", "database.name",
	"database.username", "database.password", "database.driver",
	"database.pool.min-size", "database.pool.max-size", "database.pool.timeout",
	"database.pool.idle-timeout", "database.pool.connection-test-query",
	"cache.type", "cache.host", "cache.port", "cache.password", "cache.ttl",
	"cache.pool.max-active", "cache.pool.max-idle", "cache.pool.min-idle",
	"mail.host", "mail.port", "mail.username", "mail.auth",
	"mail.tls.enabled", "mail.tls.required",
	"logging.level", "logging.file",
	"logging.pattern.console", "logging.pattern.file",
	"security.jwt.secret", "security.jwt.expiration",
	"security.jwt.refresh.secret", "security.jwt.refresh.expiration",
	"security.cors.allowed-origins", "security.cors.allowed-methods",
	// parent/intermediate nodes that must NOT be set as env vars
	"app", "server", "database", "cache", "mail", "logging", "security",
	"server.ssl", "database.pool", "cache.pool", "mail.tls",
	"logging.pattern", "security.jwt", "security.jwt.refresh", "security.cors",
}

func cleanupKeys(t *testing.T, keys []string) {
	t.Helper()
	t.Cleanup(func() {
		for _, k := range keys {
			os.Unsetenv(k)
		}
	})
}

// -----------------------------------------------
// LoadEnv (.env)
// -----------------------------------------------

func TestLoad_Base(t *testing.T) {
	cleanupKeys(t, allBaseEnvKeys)

	if err := lxenv.LoadEnv(baseEnv); err != nil {
		t.Fatalf("LoadEnv() unexpected error: %v", err)
	}

	tests := []struct{ key, want string }{
		{"APP_NAME", "lx"},
		{"APP_VERSION", "1.0.0"},
		{"APP_ENV", "development"},
		{"APP_DEBUG", "true"},
		{"APP_PORT", "8080"},
		{"DB_HOST", "localhost"},
		{"DB_PORT", "5432"},
		{"DB_NAME", "lxdb"},
		{"DB_USER", "admin"},
		{"DB_PASSWORD", "secret"},
		{"GREETING", "hello world"},
		{"SINGLE_QUOTED", "single quoted value"},
		{"STRIPPED", "value"},
		{"TOKEN", "abc#123"},
		{"SPECIAL", "p@ssw0rd!"},
		{"EMPTY_VAR", ""},
		{"DESCRIPTION", "this is a description"},
	}
	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			if got := os.Getenv(tt.key); got != tt.want {
				t.Errorf("env[%q] = %q, want %q", tt.key, got, tt.want)
			}
		})
	}
}

func TestLoad_Override(t *testing.T) {
	cleanupKeys(t, allBaseEnvKeys)

	if err := lxenv.LoadEnv(baseEnv, overrideEnv); err != nil {
		t.Fatalf("LoadEnv() unexpected error: %v", err)
	}

	overridden := []struct{ key, want string }{
		{"APP_ENV", "local"},
		{"APP_DEBUG", "false"},
		{"APP_PORT", "9090"},
		{"DB_HOST", "127.0.0.1"},
		{"DB_NAME", "lxdb_local"},
	}
	for _, tt := range overridden {
		t.Run("overridden/"+tt.key, func(t *testing.T) {
			if got := os.Getenv(tt.key); got != tt.want {
				t.Errorf("env[%q] = %q, want %q", tt.key, got, tt.want)
			}
		})
	}

	preserved := []struct{ key, want string }{
		{"APP_NAME", "lx"},
		{"APP_VERSION", "1.0.0"},
		{"DB_PORT", "5432"},
		{"DB_USER", "admin"},
	}
	for _, tt := range preserved {
		t.Run("preserved/"+tt.key, func(t *testing.T) {
			if got := os.Getenv(tt.key); got != tt.want {
				t.Errorf("env[%q] = %q, want %q", tt.key, got, tt.want)
			}
		})
	}
}

func TestLoad_NonExistentFile(t *testing.T) {
	if err := lxenv.LoadEnv("testdata/nonexistent.env"); err == nil {
		t.Error("LoadEnv() expected error for non-existent file, got nil")
	}
}

func TestLoad_NoPaths(t *testing.T) {
	if err := lxenv.LoadEnv(); err != nil {
		t.Errorf("LoadEnv() with no paths returned unexpected error: %v", err)
	}
}

func TestLoad_SecondFileNotFound(t *testing.T) {
	cleanupKeys(t, allBaseEnvKeys)

	if err := lxenv.LoadEnv(baseEnv, "testdata/nonexistent.env"); err == nil {
		t.Error("LoadEnv() expected error when second file not found, got nil")
	}
}

// -----------------------------------------------
// LoadProperties
// -----------------------------------------------

func TestLoadProperties_Base(t *testing.T) {
	cleanupKeys(t, allBasePropertiesKeys)

	if err := lxenv.LoadProperties(baseProperties); err != nil {
		t.Fatalf("LoadProperties() unexpected error: %v", err)
	}

	tests := []struct{ key, want string }{
		{"app.name", "lx"},
		{"app.version", "1.0.0"},
		{"app.env", "development"},
		{"app.debug", "true"},
		{"app.description", "this is a lx application"},
		{"server.host", "localhost"},
		{"server.port", "8080"},
		{"server.context-path", "/api"},
		{"database.host", "localhost"},
		{"database.port", "5432"},
		{"database.name", "lxdb"},
		{"database.username", "admin"},
		{"database.password", "secret"},
		{"database.driver", "org.postgresql.Driver"},
		{"cache.type", "redis"},
		{"cache.host", "localhost"},
		{"cache.port", "6379"},
		{"cache.password", ""},
		{"mail.host", "smtp.gmail.com"},
		{"mail.port", "587"},
		{"mail.username", "no-reply@lx.dev"},
		{"mail.auth", "true"},
		{"logging.level", "INFO"},
		{"logging.file", "logs/app.log"},
		{"app.greeting", "hello world"},
		{"app.single_quoted", "single quoted value"},
		{"app.stripped", "value"},
		{"app.token", "abc#123"},
		{"app.secret", "p@ssw0rd!"},
		{"app.empty", ""},
	}
	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			if got := os.Getenv(tt.key); got != tt.want {
				t.Errorf("env[%q] = %q, want %q", tt.key, got, tt.want)
			}
		})
	}
}

func TestLoadProperties_Override(t *testing.T) {
	cleanupKeys(t, allBasePropertiesKeys)

	if err := lxenv.LoadProperties(baseProperties, overrideProperties); err != nil {
		t.Fatalf("LoadProperties() unexpected error: %v", err)
	}

	overridden := []struct{ key, want string }{
		{"app.env", "local"},
		{"app.debug", "false"},
		{"server.port", "9090"},
		{"database.host", "127.0.0.1"},
		{"database.name", "lxdb_local"},
		{"database.username", "dev"},
		{"database.password", "devpassword"},
		{"cache.host", "127.0.0.1"},
		{"logging.level", "DEBUG"},
	}
	for _, tt := range overridden {
		t.Run("overridden/"+tt.key, func(t *testing.T) {
			if got := os.Getenv(tt.key); got != tt.want {
				t.Errorf("env[%q] = %q, want %q", tt.key, got, tt.want)
			}
		})
	}

	preserved := []struct{ key, want string }{
		{"app.name", "lx"},
		{"server.host", "localhost"},
		{"database.driver", "org.postgresql.Driver"},
		{"mail.host", "smtp.gmail.com"},
	}
	for _, tt := range preserved {
		t.Run("preserved/"+tt.key, func(t *testing.T) {
			if got := os.Getenv(tt.key); got != tt.want {
				t.Errorf("env[%q] = %q, want %q", tt.key, got, tt.want)
			}
		})
	}
}

func TestLoadProperties_NonExistentFile(t *testing.T) {
	if err := lxenv.LoadProperties("testdata/nonexistent.properties"); err == nil {
		t.Error("LoadProperties() expected error for non-existent file, got nil")
	}
}

func TestLoadProperties_NoPaths(t *testing.T) {
	if err := lxenv.LoadProperties(); err != nil {
		t.Errorf("LoadProperties() with no paths returned unexpected error: %v", err)
	}
}

func TestLoadProperties_SecondFileNotFound(t *testing.T) {
	cleanupKeys(t, allBasePropertiesKeys)

	if err := lxenv.LoadProperties(baseProperties, "testdata/nonexistent.properties"); err == nil {
		t.Error("LoadProperties() expected error when second file not found, got nil")
	}
}

// -----------------------------------------------
// LoadYML
// -----------------------------------------------

func TestLoadYML_Base(t *testing.T) {
	cleanupKeys(t, allBaseYMLKeys)

	if err := lxenv.LoadYML(baseYML); err != nil {
		t.Fatalf("LoadYML() unexpected error: %v", err)
	}

	tests := []struct{ key, want string }{
		// app — depth 1
		{"app.name", "lx"},
		{"app.version", "1.0.0"},
		{"app.env", "development"},
		{"app.debug", "true"},
		{"app.description", "this is a lx application"},
		{"app.greeting", "hello world"},
		{"app.secret", "p@ssw0rd!"},
		{"app.token", "abc#123"},
		{"app.empty", ""},
		{"app.stripped", "value"},
		// server — depth 2
		{"server.host", "localhost"},
		{"server.port", "8080"},
		{"server.context-path", "/api"},
		// server.ssl — depth 3
		{"server.ssl.enabled", "false"},
		{"server.ssl.key-store", "classpath:keystore.p12"},
		{"server.ssl.key-store-password", "changeit"},
		// database — depth 2
		{"database.host", "localhost"},
		{"database.port", "5432"},
		{"database.name", "lxdb"},
		{"database.username", "admin"},
		{"database.password", "secret"},
		{"database.driver", "org.postgresql.Driver"},
		// database.pool — depth 3
		{"database.pool.min-size", "2"},
		{"database.pool.max-size", "10"},
		{"database.pool.timeout", "30000"},
		{"database.pool.idle-timeout", "600000"},
		{"database.pool.connection-test-query", "SELECT 1"},
		// cache — depth 2
		{"cache.type", "redis"},
		{"cache.host", "localhost"},
		{"cache.port", "6379"},
		{"cache.password", ""},
		{"cache.ttl", "3600"},
		// cache.pool — depth 3
		{"cache.pool.max-active", "8"},
		{"cache.pool.max-idle", "8"},
		{"cache.pool.min-idle", "0"},
		// mail — depth 2
		{"mail.host", "smtp.gmail.com"},
		{"mail.port", "587"},
		{"mail.username", "no-reply@lx.dev"},
		{"mail.auth", "true"},
		// mail.tls — depth 3
		{"mail.tls.enabled", "true"},
		{"mail.tls.required", "true"},
		// logging — depth 2
		{"logging.level", "INFO"},
		{"logging.file", "logs/app.log"},
		// logging.pattern — depth 3
		{"logging.pattern.console", "%d{yyyy-MM-dd} [%thread] %-5level %logger - %msg%n"},
		{"logging.pattern.file", "%d{yyyy-MM-dd} [%thread] %-5level %logger - %msg%n"},
		// security.jwt — depth 3
		{"security.jwt.secret", "my-jwt-secret"},
		{"security.jwt.expiration", "86400"},
		// security.jwt.refresh — depth 4
		{"security.jwt.refresh.secret", "my-refresh-secret"},
		{"security.jwt.refresh.expiration", "604800"},
		// security.cors — depth 3
		{"security.cors.allowed-origins", "http://localhost:3000"},
		{"security.cors.allowed-methods", "GET,POST,PUT,DELETE"},
	}
	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			if got := os.Getenv(tt.key); got != tt.want {
				t.Errorf("env[%q] = %q, want %q", tt.key, got, tt.want)
			}
		})
	}
}

func TestLoadYML_NestedKeysNotLeakedAsTopLevel(t *testing.T) {
	cleanupKeys(t, allBaseYMLKeys)

	if err := lxenv.LoadYML(baseYML); err != nil {
		t.Fatalf("LoadYML() unexpected error: %v", err)
	}

	for _, k := range []string{"name", "host", "port", "username", "password", "secret", "enabled", "level"} {
		t.Run("not_set/"+k, func(t *testing.T) {
			if got := os.Getenv(k); got != "" {
				t.Errorf("bare key %q should not be set, got %q", k, got)
			}
		})
	}
}

func TestLoadYML_Override(t *testing.T) {
	cleanupKeys(t, allBaseYMLKeys)

	if err := lxenv.LoadYML(baseYML, overrideYML); err != nil {
		t.Fatalf("LoadYML() unexpected error: %v", err)
	}

	overridden := []struct{ key, want string }{
		{"app.env", "local"},
		{"app.debug", "false"},
		{"server.port", "9090"},
		{"server.ssl.enabled", "true"},
		{"server.ssl.key-store-password", "localpassword"},
		{"database.host", "127.0.0.1"},
		{"database.name", "lxdb_local"},
		{"database.username", "dev"},
		{"database.password", "devpassword"},
		{"database.pool.max-size", "5"},
		{"cache.host", "127.0.0.1"},
		{"cache.ttl", "300"},
		{"logging.level", "DEBUG"},
		{"security.jwt.expiration", "3600"},
		{"security.jwt.refresh.expiration", "7200"},
		{"security.cors.allowed-origins", "*"},
	}
	for _, tt := range overridden {
		t.Run("overridden/"+tt.key, func(t *testing.T) {
			if got := os.Getenv(tt.key); got != tt.want {
				t.Errorf("env[%q] = %q, want %q", tt.key, got, tt.want)
			}
		})
	}

	preserved := []struct{ key, want string }{
		{"app.name", "lx"},
		{"app.version", "1.0.0"},
		{"server.host", "localhost"},
		{"server.ssl.key-store", "classpath:keystore.p12"},
		{"database.port", "5432"},
		{"database.driver", "org.postgresql.Driver"},
		{"database.pool.min-size", "2"},
		{"cache.type", "redis"},
		{"cache.port", "6379"},
		{"mail.host", "smtp.gmail.com"},
		{"security.jwt.secret", "my-jwt-secret"},
		{"security.jwt.refresh.secret", "my-refresh-secret"},
		{"security.cors.allowed-methods", "GET,POST,PUT,DELETE"},
	}
	for _, tt := range preserved {
		t.Run("preserved/"+tt.key, func(t *testing.T) {
			if got := os.Getenv(tt.key); got != tt.want {
				t.Errorf("env[%q] = %q, want %q", tt.key, got, tt.want)
			}
		})
	}
}

func TestLoadYML_NonExistentFile(t *testing.T) {
	if err := lxenv.LoadYML("testdata/nonexistent.yml"); err == nil {
		t.Error("LoadYML() expected error for non-existent file, got nil")
	}
}

func TestLoadYML_NoPaths(t *testing.T) {
	if err := lxenv.LoadYML(); err != nil {
		t.Errorf("LoadYML() with no paths returned unexpected error: %v", err)
	}
}

func TestLoadYML_SecondFileNotFound(t *testing.T) {
	cleanupKeys(t, allBaseYMLKeys)

	if err := lxenv.LoadYML(baseYML, "testdata/nonexistent.yml"); err == nil {
		t.Error("LoadYML() expected error when second file not found, got nil")
	}
}

func TestLoadYML_ParentKeysNotSetInEnv(t *testing.T) {
	cleanupKeys(t, allBaseYMLKeys)

	if err := lxenv.LoadYML(baseYML); err != nil {
		t.Fatalf("LoadYML() unexpected error: %v", err)
	}

	parentKeys := []string{
		"app", "server", "database", "cache", "mail", "logging", "security",
		"server.ssl", "database.pool", "cache.pool", "mail.tls",
		"logging.pattern", "security.jwt", "security.jwt.refresh", "security.cors",
	}
	for _, k := range parentKeys {
		t.Run("not_set/"+k, func(t *testing.T) {
			if _, exists := os.LookupEnv(k); exists {
				t.Errorf("parent key %q should not be set in env, but it is (value=%q)", k, os.Getenv(k))
			}
		})
	}
}

func TestLoadYML_ListItemsSkipped(t *testing.T) {
	content := []byte(`server:
  host: localhost
  allowed:
    - item1
    - item2
  port: 8080
`)

	f, err := os.CreateTemp("", "lxenv_list_test_*.yml")
	if err != nil {
		t.Fatalf("could not create temp file: %v", err)
	}
	t.Cleanup(func() { os.Remove(f.Name()) })

	if _, err := f.Write(content); err != nil {
		f.Close()
		t.Fatalf("could not write temp file: %v", err)
	}
	f.Close()

	keysToClean := []string{"server", "server.host", "server.allowed", "server.port"}
	cleanupKeys(t, keysToClean)

	if err := lxenv.LoadYML(f.Name()); err != nil {
		t.Fatalf("LoadYML() unexpected error: %v", err)
	}

	if got := os.Getenv("server.host"); got != "localhost" {
		t.Errorf("env[%q] = %q, want %q", "server.host", got, "localhost")
	}
	if got := os.Getenv("server.port"); got != "8080" {
		t.Errorf("env[%q] = %q, want %q", "server.port", got, "8080")
	}
	if _, exists := os.LookupEnv("server.allowed"); exists {
		t.Errorf("list parent key %q should not be set in env", "server.allowed")
	}
	if _, exists := os.LookupEnv("server"); exists {
		t.Errorf("parent key %q should not be set in env", "server")
	}
}
