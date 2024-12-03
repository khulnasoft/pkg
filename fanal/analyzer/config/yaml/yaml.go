package yaml

import (
	"os"
	"path/filepath"

	"go.khulnasoft.com/pkg/fanal/analyzer"
	"go.khulnasoft.com/pkg/fanal/analyzer/config"
	"go.khulnasoft.com/pkg/iac/detection"
)

const (
	analyzerType = analyzer.TypeYAML
	version      = 1
)

func init() {
	analyzer.RegisterPostAnalyzer(analyzerType, newYAMLConfigAnalyzer)
}

// yamlConfigAnalyzer analyzes YAML files
type yamlConfigAnalyzer struct {
	*config.Analyzer
}

func newYAMLConfigAnalyzer(opts analyzer.AnalyzerOptions) (analyzer.PostAnalyzer, error) {
	a, err := config.NewAnalyzer(analyzerType, version, detection.FileTypeYAML, opts)
	if err != nil {
		return nil, err
	}
	return &yamlConfigAnalyzer{Analyzer: a}, nil
}

func (*yamlConfigAnalyzer) Required(filePath string, _ os.FileInfo) bool {
	return filepath.Ext(filePath) == ".yaml" || filepath.Ext(filePath) == ".yml"
}
