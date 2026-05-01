package generator

// GenerateShellScript generates a shell installer script
// This is a convenience function that wraps generateShellScript
func GenerateShellScript(cfg Config) (string, error) {
	return generateShellScript(cfg)
}
