package session

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetClaudeConfigDir_Default(t *testing.T) {
	// Unset env var to test default/config behavior
	os.Unsetenv("CLAUDE_CONFIG_DIR")

	dir := GetClaudeConfigDir()
	home, _ := os.UserHomeDir()
	defaultPath := filepath.Join(home, ".claude")

	// If user config exists with claude.config_dir, that takes precedence
	// Otherwise, default to ~/.claude
	userConfig, _ := LoadUserConfig()
	if userConfig != nil && userConfig.Claude.ConfigDir != "" {
		// Config exists, just verify we get a valid path
		if dir == "" {
			t.Error("GetClaudeConfigDir() returned empty string")
		}
	} else {
		// No config, should return default
		if dir != defaultPath {
			t.Errorf("GetClaudeConfigDir() = %s, want %s", dir, defaultPath)
		}
	}
}

func TestGetClaudeConfigDir_EnvOverride(t *testing.T) {
	os.Setenv("CLAUDE_CONFIG_DIR", "/custom/path")
	defer os.Unsetenv("CLAUDE_CONFIG_DIR")

	dir := GetClaudeConfigDir()
	if dir != "/custom/path" {
		t.Errorf("GetClaudeConfigDir() = %s, want /custom/path", dir)
	}
}

func TestGetClaudeSessionID_NotFound(t *testing.T) {
	id, err := GetClaudeSessionID("/nonexistent/path")
	if err == nil {
		t.Error("Expected error for nonexistent path")
	}
	if id != "" {
		t.Errorf("Expected empty ID, got %s", id)
	}
}
