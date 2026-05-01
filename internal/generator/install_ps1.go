package generator

// GeneratePowerShellScript generates a PowerShell installer script
// This is a convenience function that wraps generatePowerShellScript
func GeneratePowerShellScript(cfg Config) (string, error) {
	return generatePowerShellScript(cfg)
}
