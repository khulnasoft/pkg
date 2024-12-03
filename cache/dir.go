package cache

import (
	"os"
	"path/filepath"
)

// DefaultDir returns/creates the cache-dir to be used for tunnel operations
func DefaultDir() string {
	tmpDir, err := os.UserCacheDir()
	if err != nil {
		tmpDir = os.TempDir()
	}
	return filepath.Join(tmpDir, "tunnel")
}
