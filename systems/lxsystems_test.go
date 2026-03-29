package lxsystems_test

import (
	"os"
	"runtime"
	"testing"

	"github.com/hgapdvn/lx/systems"
)

func TestConstants(t *testing.T) {
	t.Run("OSName", func(t *testing.T) {
		if lxsystems.OSName != runtime.GOOS {
			t.Errorf("OSName = %q; want %q", lxsystems.OSName, runtime.GOOS)
		}
	})

	t.Run("OSArch", func(t *testing.T) {
		if lxsystems.OSArch != runtime.GOARCH {
			t.Errorf("OSArch = %q; want %q", lxsystems.OSArch, runtime.GOARCH)
		}
	})
}

func TestGoVersion(t *testing.T) {
	got := lxsystems.GoVersion()
	want := runtime.Version()

	if got != want {
		t.Errorf("GoVersion() = %q; want %q", got, want)
	}
}

func TestUserHomeDir(t *testing.T) {
	got, gotErr := lxsystems.UserHomeDir()
	want, wantErr := os.UserHomeDir()

	if (gotErr != nil) != (wantErr != nil) {
		t.Fatalf("UserHomeE() error = %v; want error = %v", gotErr, wantErr)
	}

	if wantErr != nil {
		// Both errored: nothing more to compare deterministically.
		return
	}

	if got != want {
		t.Errorf("UserHomeE() = %q; want %q", got, want)
	}
}

func TestUserHomeDirOrEmpty(t *testing.T) {
	got := lxsystems.UserHomeDirOrEmpty()
	want, _ := os.UserHomeDir()
	if got != want {
		t.Errorf("UserHomeDirOrEmpty() = %q; want %q", got, want)
	}
}

func TestUserHomeDirOr(t *testing.T) {
	def := "DEFAULT_HOME"

	got := lxsystems.UserHomeDirOr(def)
	want, err := os.UserHomeDir()

	if err != nil {
		if got != def {
			t.Errorf("UserHomeOr(%q) = %q; want %q when os.UserHomeDir() errors", def, got, def)
		}
		return
	}

	if got != want {
		t.Errorf("UserHomeOr(%q) = %q; want %q", def, got, want)
	}
}

func TestTempDir(t *testing.T) {
	got := lxsystems.TempDir()
	want := os.TempDir()

	if got != want {
		t.Errorf("TempDir() = %q; want %q", got, want)
	}
}

func TestWorkingDir(t *testing.T) {
	got, gotErr := lxsystems.WorkingDir()
	want, wantErr := os.Getwd()

	if (gotErr != nil) != (wantErr != nil) {
		t.Fatalf("WorkingDir() error = %v; want error = %v", gotErr, wantErr)
	}

	if wantErr != nil {
		// Both errored: nothing deterministic to compare.
		return
	}

	if got != want {
		t.Errorf("WorkingDir() = %q; want %q", got, want)
	}
}

func TestWorkingDirOrEmpty(t *testing.T) {
	got := lxsystems.WorkingDirOrEmpty()

	want, err := os.Getwd()
	if err != nil {
		if got != "" {
			t.Errorf("WorkingDirOrEmpty() = %q; want empty string when os.Getwd() errors", got)
		}
		return
	}

	if got != want {
		t.Errorf("WorkingDirOrEmpty() = %q; want %q", got, want)
	}
}

func TestWorkingDirOr(t *testing.T) {
	def := "DEFAULT_WD"

	got := lxsystems.WorkingDirOr(def)
	want, err := os.Getwd()

	if err != nil {
		if got != def {
			t.Errorf("WorkingDirOr(%q) = %q; want %q when os.Getwd() errors", def, got, def)
		}
		return
	}

	if got != want {
		t.Errorf("WorkingDirOr(%q) = %q; want %q", def, got, want)
	}
}

func TestNumCPU(t *testing.T) {
	got := lxsystems.NumCPU()
	want := runtime.NumCPU()

	if got != want {
		t.Errorf("NumCPU() = %d; want %d", got, want)
	}
	if got <= 0 {
		t.Errorf("NumCPU() = %d; want > 0", got)
	}
}

func TestIsOSHelpers(t *testing.T) {
	t.Run("IsWindows matches runtime.GOOS", func(t *testing.T) {
		got := lxsystems.IsWindows()
		want := runtime.GOOS == "windows"
		if got != want {
			t.Errorf("IsWindows() = %v; want %v", got, want)
		}
	})

	t.Run("IsLinux matches runtime.GOOS", func(t *testing.T) {
		got := lxsystems.IsLinux()
		want := runtime.GOOS == "linux"
		if got != want {
			t.Errorf("IsLinux() = %v; want %v", got, want)
		}
	})

	t.Run("IsMacOS matches runtime.GOOS", func(t *testing.T) {
		got := lxsystems.IsMacOS()
		want := runtime.GOOS == "darwin"
		if got != want {
			t.Errorf("IsMacOS() = %v; want %v", got, want)
		}
	})
}

func TestHostName(t *testing.T) {
	got, gotErr := lxsystems.HostName()
	want, wantErr := os.Hostname()

	if (gotErr != nil) != (wantErr != nil) {
		t.Fatalf("HostName() error = %v; want error = %v", gotErr, wantErr)
	}

	if wantErr != nil {
		// Both errored: nothing deterministic to compare beyond this.
		return
	}

	if got != want {
		t.Errorf("HostName() = %q; want %q", got, want)
	}
}

func TestHostNameOrEmpty(t *testing.T) {
	got := lxsystems.HostNameOrEmpty()

	want, err := os.Hostname()
	if err != nil {
		if got != "" {
			t.Errorf("HostNameOrEmpty() = %q; want empty string when os.Hostname() errors", got)
		}
		return
	}

	if got != want {
		t.Errorf("HostNameOrEmpty() = %q; want %q", got, want)
	}
}

func TestHostNameOr(t *testing.T) {
	def := "DEFAULT_HOSTNAME"

	got := lxsystems.HostNameOr(def)
	want, err := os.Hostname()

	if err != nil {
		if got != def {
			t.Errorf("HostNameOr(%q) = %q; want %q when os.Hostname() errors", def, got, def)
		}
		return
	}

	if got != want {
		t.Errorf("HostNameOr(%q) = %q; want %q", def, got, want)
	}
}

func TestPID(t *testing.T) {
	got := lxsystems.PID()
	want := os.Getpid()

	if got != want {
		t.Errorf("PID() = %d; want %d", got, want)
	}
	if got <= 0 {
		t.Errorf("PID() = %d; want > 0", got)
	}
}
