package linker

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	osutil "github.com/aleks-ship-it/pave/internal/os"
)

type Registry struct {
	Links []Link `json:"links"`
}

func LoadRegistry() (*Registry, error) {
	path, err := osutil.GetLinksFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Registry{Links: []Link{}}, nil
		}
		return nil, fmt.Errorf("failed to read registry: %w", err)
	}

	var registry Registry
	if err := json.Unmarshal(data, &registry); err != nil {
		return nil, fmt.Errorf("failed to parse registry: %w", err)
	}

	if registry.Links == nil {
		registry.Links = []Link{}
	}

	return &registry, nil
}

func (r *Registry) Save() error {
	path, err := osutil.GetLinksFilePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal registry: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func (r *Registry) AddLink(link Link) {
	r.RemoveLink(link.Name)
	r.Links = append(r.Links, link)
}

func (r *Registry) RemoveLink(name string) {
	var filtered []Link
	for _, l := range r.Links {
		if l.Name != name {
			filtered = append(filtered, l)
		}
	}
	r.Links = filtered
}

func (r *Registry) FindLink(name string) *Link {
	for _, l := range r.Links {
		if l.Name == name {
			return &l
		}
	}
	return nil
}
