package pipenv

import (
	"context"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"

	"go.khulnasoft.com/pkg/dependency/parser/python/pipenv"
	"go.khulnasoft.com/pkg/fanal/analyzer"
	"go.khulnasoft.com/pkg/fanal/analyzer/language"
	"go.khulnasoft.com/pkg/fanal/types"
	"go.khulnasoft.com/pkg/fanal/utils"
)

func init() {
	analyzer.RegisterAnalyzer(&pipenvLibraryAnalyzer{})
}

const version = 1

var requiredFiles = []string{types.PipfileLock}

type pipenvLibraryAnalyzer struct{}

func (a pipenvLibraryAnalyzer) Analyze(_ context.Context, input analyzer.AnalysisInput) (*analyzer.AnalysisResult, error) {
	res, err := language.Analyze(types.Pipenv, input.FilePath, input.Content, pipenv.NewParser())
	if err != nil {
		return nil, xerrors.Errorf("unable to parse Pipfile.lock: %w", err)
	}
	return res, nil
}

func (a pipenvLibraryAnalyzer) Required(filePath string, _ os.FileInfo) bool {
	fileName := filepath.Base(filePath)
	return utils.StringInSlice(fileName, requiredFiles)
}

func (a pipenvLibraryAnalyzer) Type() analyzer.Type {
	return analyzer.TypePipenv
}

func (a pipenvLibraryAnalyzer) Version() int {
	return version
}
