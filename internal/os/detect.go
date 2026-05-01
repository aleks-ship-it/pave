package osutil

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// GetConfigDir returns the pave configuration directory
func GetConfigDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("failed to get config dir: %w", err)
	}
	dir := filepath.Join(base, "pave")
	return ensureDir(dir)
}

// GetDataDir returns the pave data directory for storing links registry
func GetDataDir() (string, error) {
	if runtime.GOOS == "linux" {
		dataHome := os.Getenv("XDG_DATA_HOME")
		if dataHome == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			dataHome = filepath.Join(home, ".local", "share")
		}
		dir := filepath.Join(dataHome, "pave")
		return ensureDir(dir)
	}
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(base, "pave", "data")
	return ensureDir(dir)
}

// GetLinksFilePath returns the path to the links registry JSON file
func GetLinksFilePath() (string, error) {
	dir, err := GetDataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "links.json"), nil
}

// GetBinDir returns the directory where symlinks are stored (should be in PATH)
func GetBinDir() (string, error) {
	var dir string
	switch runtime.GOOS {
	case "linux":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		dir = filepath.Join(home, ".local", "bin")
	case "darwin":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		dir = filepath.Join(home, "bin")
	case "windows":
		base, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		dir = filepath.Join(base, "pave", "bin")
	default:
		return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
	return ensureDir(dir)
}

// IsPathInPATH checks if the given directory is in the system PATH
func IsPathInPATH(dir string) bool {
	pathEnv := os.Getenv("PATH")
	dirs := filepath.SplitList(pathEnv)
	for _, d := range dirs {
		if filepath.Clean(d) == filepath.Clean(dir) {
			return true
		}
	}
	return false
}

// ensureDir creates the directory if it doesn't exist
func ensureDir(dir string) (string, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create dir %s: %w", dir, err)
	}
	return dir, nil
}
