// Package save persists the character to a JSON file. The framework ledger
// replaces it once that exists.
package save

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/unstoppablemango/ouranosis/pkg/character"
)

// DefaultPath is the save file under the user's config directory.
func DefaultPath() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		dir = "."
	}
	return filepath.Join(dir, "ouranosis", "save.json")
}

// Load reads the character at path.
// It returns fs.ErrNotExist when there is no save.
func Load(path string) (*character.Character, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var c character.Character
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

// Save writes the character to path, creating directories as needed.
func Save(path string, c *character.Character) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o600)
}
