package system

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"
)

// OS returns operating system
func OS() string {
	return runtime.GOOS
}

// Arch returns architecture
func Arch() string {
	return runtime.GOARCH
}

// NumCPU returns number of CPUs
func NumCPU() int {
	return runtime.NumCPU()
}

// GoVersion returns Go version
func GoVersion() string {
	return runtime.Version()
}

// Hostname returns hostname
func Hostname() (string, error) {
	return os.Hostname()
}

// Username returns current username
func Username() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("failed to get current user: %w", err)
	}
	return u.Username, nil
}

// UserHomeDir returns user home directory
func UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

// WorkingDir returns working directory
func WorkingDir() (string, error) {
	return os.Getwd()
}

// Env gets environment variable
func Env(key string) string {
	return os.Getenv(key)
}

// EnvWithDefault gets environment variable with default
func EnvWithDefault(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

// SetEnv sets environment variable
func SetEnv(key, value string) error {
	return os.Setenv(key, value)
}

// UnsetEnv unsets environment variable
func UnsetEnv(key string) error {
	return os.Unsetenv(key)
}

// Environ returns all environment variables
func Environ() map[string]string {
	result := make(map[string]string)
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 {
			result[parts[0]] = parts[1]
		}
	}
	return result
}

// PID returns process ID
func PID() int {
	return os.Getpid()
}

// PPID returns parent process ID
func PPID() int {
	return os.Getppid()
}

// IsWindows checks if OS is Windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux checks if OS is Linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsMac checks if OS is macOS
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// Is64Bit checks if architecture is 64-bit
func Is64Bit() bool {
	return runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64"
}

// MemoryInfo returns memory information
type MemoryInfo struct {
	Total     uint64
	Free      uint64
	Available uint64
}

// GetMemoryInfo returns memory information
func GetMemoryInfo() (*MemoryInfo, error) {
	return getMemoryInfo()
}

// DiskInfo returns disk information
type DiskInfo struct {
	Total uint64
	Free  uint64
	Used  uint64
}

// GetDiskInfo returns disk information for path
func GetDiskInfo(path string) (*DiskInfo, error) {
	return getDiskInfo(path)
}

// UserInfo returns current user information
type UserInfo struct {
	Username string
	UID      string
	GID      string
	HomeDir  string
	Name     string
}

// GetUserInfo returns current user information
func GetUserInfo() (*UserInfo, error) {
	u, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}

	return &UserInfo{
		Username: u.Username,
		UID:      u.Uid,
		GID:      u.Gid,
		HomeDir:  u.HomeDir,
		Name:     u.Name,
	}, nil
}

// TempDir returns temporary directory
func TempDir() string {
	return os.TempDir()
}

// LineSeparator returns line separator
func LineSeparator() string {
	if IsWindows() {
		return "\r\n"
	}
	return "\n"
}

// FileSeparator returns file separator
func FileSeparator() string {
	if IsWindows() {
		return "\\"
	}
	return "/"
}
