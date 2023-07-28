package system

import (
	"os"
	"path/filepath"
)

func ReadFile(path string) ([]byte, error) {
	path = filepath.Join("/data", path)
	return os.ReadFile(path)
}
