package system

import (
	"os"
	"path/filepath"
)

func WriteFile(path string, b []byte) error {
	path = filepath.Join("/data", path)

	// Make sure dir exists
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	// Write file
	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
