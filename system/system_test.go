package system

import (
	"os"
	"runtime"
	"testing"
)

func TestOS(t *testing.T) {
	got := OS()
	if got != runtime.GOOS {
		t.Errorf("OS() = %v, want %v", got, runtime.GOOS)
	}
}

func TestArch(t *testing.T) {
	got := Arch()
	if got != runtime.GOARCH {
		t.Errorf("Arch() = %v, want %v", got, runtime.GOARCH)
	}
}

func TestNumCPU(t *testing.T) {
	got := NumCPU()
	if got != runtime.NumCPU() {
		t.Errorf("NumCPU() = %v, want %v", got, runtime.NumCPU())
	}
	if got <= 0 {
		t.Errorf("NumCPU() = %v, want > 0", got)
	}
}

func TestGoVersion(t *testing.T) {
	got := GoVersion()
	if got != runtime.Version() {
		t.Errorf("GoVersion() = %v, want %v", got, runtime.Version())
	}
}

func TestHostname(t *testing.T) {
	got, err := Hostname()
	if err != nil {
		t.Fatalf("Hostname() error: %v", err)
	}
	if got == "" {
		t.Error("Hostname() returned empty string")
	}

	expected, _ := os.Hostname()
	if got != expected {
		t.Errorf("Hostname() = %v, want %v", got, expected)
	}
}

func TestUsername(t *testing.T) {
	got, err := Username()
	if err != nil {
		t.Fatalf("Username() error: %v", err)
	}
	if got == "" {
		t.Error("Username() returned empty string")
	}
}

func TestUserHomeDir(t *testing.T) {
	got, err := UserHomeDir()
	if err != nil {
		t.Fatalf("UserHomeDir() error: %v", err)
	}
	if got == "" {
		t.Error("UserHomeDir() returned empty string")
	}

	expected, _ := os.UserHomeDir()
	if got != expected {
		t.Errorf("UserHomeDir() = %v, want %v", got, expected)
	}
}

func TestWorkingDir(t *testing.T) {
	got, err := WorkingDir()
	if err != nil {
		t.Fatalf("WorkingDir() error: %v", err)
	}
	if got == "" {
		t.Error("WorkingDir() returned empty string")
	}

	expected, _ := os.Getwd()
	if got != expected {
		t.Errorf("WorkingDir() = %v, want %v", got, expected)
	}
}

func TestEnv(t *testing.T) {
	// Test with existing environment variable
	path := Env("PATH")
	if path == "" {
		t.Error("Env(PATH) returned empty string")
	}

	// Test with nonexistent environment variable
	nonexistent := Env("NONEXISTENT_VAR_12345")
	if nonexistent != "" {
		t.Errorf("Env(NONEXISTENT_VAR_12345) = %v, want empty string", nonexistent)
	}
}

func TestEnvWithDefault(t *testing.T) {
	// Test with existing environment variable
	path := EnvWithDefault("PATH", "/default")
	if path == "" {
		t.Error("EnvWithDefault(PATH) returned empty string")
	}

	// Test with nonexistent environment variable
	got := EnvWithDefault("NONEXISTENT_VAR_12345", "/default")
	if got != "/default" {
		t.Errorf("EnvWithDefault(NONEXISTENT_VAR_12345) = %v, want /default", got)
	}
}

func TestSetUnsetEnv(t *testing.T) {
	key := "TEST_VAR_12345"
	value := "test_value"

	// Set environment variable
	if err := SetEnv(key, value); err != nil {
		t.Fatalf("SetEnv() error: %v", err)
	}

	// Verify it was set
	got := Env(key)
	if got != value {
		t.Errorf("Env(%s) = %v, want %v", key, got, value)
	}

	// Unset environment variable
	if err := UnsetEnv(key); err != nil {
		t.Fatalf("UnsetEnv() error: %v", err)
	}

	// Verify it was unset
	got = Env(key)
	if got != "" {
		t.Errorf("Env(%s) = %v, want empty string", key, got)
	}
}

func TestEnviron(t *testing.T) {
	envs := Environ()
	if len(envs) == 0 {
		t.Error("Environ() returned empty map")
	}

	// Check if PATH is present
	if _, ok := envs["PATH"]; !ok {
		t.Error("Environ() does not contain PATH")
	}

	// Check if HOME or USERPROFILE is present (depending on OS)
	if IsWindows() {
		if _, ok := envs["USERPROFILE"]; !ok {
			t.Error("Environ() does not contain USERPROFILE on Windows")
		}
	} else {
		if _, ok := envs["HOME"]; !ok {
			t.Error("Environ() does not contain HOME on Unix")
		}
	}
}

func TestPID(t *testing.T) {
	got := PID()
	if got <= 0 {
		t.Errorf("PID() = %v, want > 0", got)
	}

	expected := os.Getpid()
	if got != expected {
		t.Errorf("PID() = %v, want %v", got, expected)
	}
}

func TestPPID(t *testing.T) {
	got := PPID()
	if got <= 0 {
		t.Errorf("PPID() = %v, want > 0", got)
	}

	expected := os.Getppid()
	if got != expected {
		t.Errorf("PPID() = %v, want %v", got, expected)
	}
}

func TestIsWindows(t *testing.T) {
	got := IsWindows()
	expected := runtime.GOOS == "windows"
	if got != expected {
		t.Errorf("IsWindows() = %v, want %v", got, expected)
	}
}

func TestIsLinux(t *testing.T) {
	got := IsLinux()
	expected := runtime.GOOS == "linux"
	if got != expected {
		t.Errorf("IsLinux() = %v, want %v", got, expected)
	}
}

func TestIsMac(t *testing.T) {
	got := IsMac()
	expected := runtime.GOOS == "darwin"
	if got != expected {
		t.Errorf("IsMac() = %v, want %v", got, expected)
	}
}

func TestIs64Bit(t *testing.T) {
	got := Is64Bit()
	expected := runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64"
	if got != expected {
		t.Errorf("Is64Bit() = %v, want %v", got, expected)
	}
}

func TestGetMemoryInfo(t *testing.T) {
	info, err := GetMemoryInfo()
	if err != nil {
		t.Fatalf("GetMemoryInfo() error: %v", err)
	}

	if info.Total == 0 {
		t.Error("GetMemoryInfo().Total = 0, want > 0")
	}
	if info.Free == 0 {
		t.Error("GetMemoryInfo().Free = 0, want > 0")
	}
	if info.Available == 0 {
		t.Error("GetMemoryInfo().Available = 0, want > 0")
	}

	// Free should be less than or equal to total
	if info.Free > info.Total {
		t.Errorf("GetMemoryInfo().Free (%v) > Total (%v)", info.Free, info.Total)
	}
}

func TestGetDiskInfo(t *testing.T) {
	// Test with current directory
	info, err := GetDiskInfo(".")
	if err != nil {
		t.Fatalf("GetDiskInfo() error: %v", err)
	}

	if info.Total == 0 {
		t.Error("GetDiskInfo().Total = 0, want > 0")
	}
	if info.Free == 0 {
		t.Error("GetDiskInfo().Free = 0, want > 0")
	}
	if info.Used == 0 {
		t.Error("GetDiskInfo().Used = 0, want > 0")
	}

	// Used should be less than or equal to total
	if info.Used > info.Total {
		t.Errorf("GetDiskInfo().Used (%v) > Total (%v)", info.Used, info.Total)
	}
}

func TestGetUserInfo(t *testing.T) {
	info, err := GetUserInfo()
	if err != nil {
		t.Fatalf("GetUserInfo() error: %v", err)
	}

	if info.Username == "" {
		t.Error("GetUserInfo().Username is empty")
	}
	if info.UID == "" {
		t.Error("GetUserInfo().UID is empty")
	}
	if info.GID == "" {
		t.Error("GetUserInfo().GID is empty")
	}
	if info.HomeDir == "" {
		t.Error("GetUserInfo().HomeDir is empty")
	}
}

func TestTempDir(t *testing.T) {
	got := TempDir()
	if got == "" {
		t.Error("TempDir() returned empty string")
	}

	expected := os.TempDir()
	if got != expected {
		t.Errorf("TempDir() = %v, want %v", got, expected)
	}
}

func TestLineSeparator(t *testing.T) {
	got := LineSeparator()
	if IsWindows() {
		if got != "\r\n" {
			t.Errorf("LineSeparator() = %q, want %q on Windows", got, "\r\n")
		}
	} else {
		if got != "\n" {
			t.Errorf("LineSeparator() = %q, want %q on Unix", got, "\n")
		}
	}
}

func TestFileSeparator(t *testing.T) {
	got := FileSeparator()
	if IsWindows() {
		if got != "\\" {
			t.Errorf("FileSeparator() = %q, want %q on Windows", got, "\\")
		}
	} else {
		if got != "/" {
			t.Errorf("FileSeparator() = %q, want %q on Unix", got, "/")
		}
	}
}
