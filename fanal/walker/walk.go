package walker

import (
	"os"

	"go.khulnasoft.com/pkg/fanal/analyzer"
)

const defaultSizeThreshold = int64(100) << 20 // 200MB

var defaultSkipDirs = []string{
	"**/.git",
	"proc",
	"sys",
	"dev",
}

type Option struct {
	SkipFiles []string
	SkipDirs  []string
}

type WalkFunc func(filePath string, info os.FileInfo, opener analyzer.Opener) error
