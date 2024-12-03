package cloudformation

import (
	"go.khulnasoft.com/pkg/fanal/analyzer"
	"go.khulnasoft.com/pkg/fanal/analyzer/config"
	"go.khulnasoft.com/pkg/iac/detection"
)

const (
	analyzerType = analyzer.TypeCloudFormation
	version      = 1
)

func init() {
	analyzer.RegisterPostAnalyzer(analyzerType, newCloudFormationConfigAnalyzer)
}

// cloudFormationConfigAnalyzer is an analyzer for detecting misconfigurations in CloudFormation files.
// It embeds config.Analyzer so it can implement analyzer.PostAnalyzer.
type cloudFormationConfigAnalyzer struct {
	*config.Analyzer
}

func newCloudFormationConfigAnalyzer(opts analyzer.AnalyzerOptions) (analyzer.PostAnalyzer, error) {
	a, err := config.NewAnalyzer(analyzerType, version, detection.FileTypeCloudFormation, opts)
	if err != nil {
		return nil, err
	}
	return &cloudFormationConfigAnalyzer{Analyzer: a}, nil
}
