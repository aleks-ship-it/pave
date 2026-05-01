package linker

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Link struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Target string `json:"target,omitempty"`
}

type LinkStatus struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Target string `json:"target,omitempty"`
	Status string `json:"status"`
}

// CreateLink creates a symlink or wrapper script for the given name and path
func CreateLink(name, path string, verbose bool) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	if _, err := os.Stat(absPath); err != nil {
		return fmt.Errorf("target path does not exist: %s", absPath)
	}

	binDir, err := getBinDir()
	if err != nil {
		return err
	}

	var target string
	if runtime.GOOS == "windows" {
		target, err = createWindowsWrapper(binDir, name, absPath)
	} else {
		target, err = createUnixSymlink(binDir, name, absPath)
	}
	if err != nil {
		return err
	}

	if verbose {
		fmt.Printf("Created: %s -> %s\n", filepath.Join(binDir, name), absPath)
	}

	registry, err := LoadRegistry()
	if err != nil {
		return err
	}

	link := Link{
		Name:   name,
		Path:   absPath,
		Target: target,
	}
	registry.AddLink(link)
	return registry.Save()
}

// RemoveLink removes a symlink or wrapper script by name
func RemoveLink(name string, verbose bool) error {
	binDir, err := getBinDir()
	if err != nil {
		return err
	}

	var linkPath string
	if runtime.GOOS == "windows" {
		linkPath = filepath.Join(binDir, name+".cmd")
	} else {
		linkPath = filepath.Join(binDir, name)
	}

	if err := os.Remove(linkPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove link: %w", err)
	}

	if verbose {
		fmt.Printf("Removed: %s\n", linkPath)
	}

	registry, err := LoadRegistry()
	if err != nil {
		return err
	}

	registry.RemoveLink(name)
	return registry.Save()
}

// ListLinks returns all managed links with their status
func ListLinks(verbose bool) ([]LinkStatus, error) {
	registry, err := LoadRegistry()
	if err != nil {
		return nil, err
	}

	var results []LinkStatus
	for _, link := range registry.Links {
		status := "valid"
		if _, err := os.Stat(link.Path); err != nil {
			status = "broken"
		}
		results = append(results, LinkStatus{
			Name:   link.Name,
			Path:   link.Path,
			Target: link.Target,
			Status: status,
		})
	}

	if verbose {
		fmt.Printf("Found %d link(s)\n", len(results))
	}

	return results, nil
}

// StatusLink returns the status of a specific link
func StatusLink(name string, verbose bool) (*LinkStatus, error) {
	registry, err := LoadRegistry()
	if err != nil {
		return nil, err
	}

	link := registry.FindLink(name)
	if link == nil {
		return nil, nil
	}

	status := "valid"
	if _, err := os.Stat(link.Path); err != nil {
		status = "broken"
	}

	if verbose {
		fmt.Printf("Checking status of %s\n", name)
	}

	return &LinkStatus{
		Name:   link.Name,
		Path:   link.Path,
		Target: link.Target,
		Status: status,
	}, nil
}

func createUnixSymlink(binDir, name, target string) (string, error) {
	linkPath := filepath.Join(binDir, name)
	if err := os.Symlink(target, linkPath); err != nil {
		return "", fmt.Errorf("failed to create symlink: %w", err)
	}
	return linkPath, nil
}

func createWindowsWrapper(binDir, name, target string) (string, error) {
	wrapperPath := filepath.Join(binDir, name+".cmd")
	// Get proper Windows path (use short names if needed)
	target = strings.ReplaceAll(target, "/", "\\")

	content := fmt.Sprintf("@echo off\r\n\"%s\" %%*\r\n", target)
	if err := os.WriteFile(wrapperPath, []byte(content), 0755); err != nil {
		return "", fmt.Errorf("failed to create wrapper: %w", err)
	}
	return wrapperPath, nil
}

func getBinDir() (string, error) {
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
		// On Windows, use a subfolder in UserConfigDir
		base, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		dir = filepath.Join(base, "pave", "bin")
	default:
		return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create bin dir: %w", err)
	}

	// Check if dir is in PATH
	if !isInPath(dir) {
		fmt.Printf("Warning: %s is not in your PATH. Add it to use commands globally.\n", dir)
		if runtime.GOOS == "linux" {
			fmt.Printf("Add with: echo 'export PATH=%s:$PATH' >> ~/.bashrc\n", dir)
		} else if runtime.GOOS == "darwin" {
			fmt.Printf("Add with: echo 'export PATH=%s:$PATH' >> ~/.zshrc\n", dir)
		}
	}

	return dir, nil
}

func isInPath(dir string) bool {
	pathEnv := os.Getenv("PATH")
	dirs := filepath.SplitList(pathEnv)
	absDir, _ := filepath.Abs(dir)
	for _, d := range dirs {
		absD, _ := filepath.Abs(d)
		if absD == absDir {
			return true
		}
	}
	return false
}
