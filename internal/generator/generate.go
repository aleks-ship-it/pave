package generator

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*
var templateFS embed.FS

type Config struct {
	Name string
	Repo string
	Bin  string
}

func Generate(cfg Config, outDir string, dryRun, verbose bool) error {
	if outDir == "" {
		outDir = "."
	}

	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	// Generate shell script
	shContent, err := generateShellScript(cfg)
	if err != nil {
		return fmt.Errorf("failed to generate shell script: %w", err)
	}

	shPath := filepath.Join(outDir, "install.sh")
	if dryRun {
		fmt.Println("--- install.sh ---")
		fmt.Println(shContent)
	} else {
		if err := os.WriteFile(shPath, []byte(shContent), 0755); err != nil {
			return fmt.Errorf("failed to write install.sh: %w", err)
		}
		if verbose {
			fmt.Printf("Generated: %s\n", shPath)
		}
	}

	// Generate PowerShell script
	psContent, err := generatePowerShellScript(cfg)
	if err != nil {
		return fmt.Errorf("failed to generate PowerShell script: %w", err)
	}

	psPath := filepath.Join(outDir, "install.ps1")
	if dryRun {
		fmt.Println("--- install.ps1 ---")
		fmt.Println(psContent)
	} else {
		if err := os.WriteFile(psPath, []byte(psContent), 0755); err != nil {
			return fmt.Errorf("failed to write install.ps1: %w", err)
		}
		if verbose {
			fmt.Printf("Generated: %s\n", psPath)
		}
	}

	return nil
}

func generateShellScript(cfg Config) (string, error) {
	return generateFromTemplate("install.sh.tmpl", cfg)
}

func generatePowerShellScript(cfg Config) (string, error) {
	return generateFromTemplate("install.ps1.tmpl", cfg)
}

func generateFromTemplate(tmplName string, cfg Config) (string, error) {
	tmplPath := "templates/" + tmplName

	tmplData, err := templateFS.ReadFile(tmplPath)
	if err != nil {
		return "", fmt.Errorf("failed to read template %s: %w", tmplName, err)
	}

	tmpl, err := template.New(tmplName).Parse(string(tmplData))
	if err != nil {
		return "", fmt.Errorf("failed to parse template %s: %w", tmplName, err)
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, cfg); err != nil {
		return "", fmt.Errorf("failed to execute template %s: %w", tmplName, err)
	}

	return buf.String(), nil
}
