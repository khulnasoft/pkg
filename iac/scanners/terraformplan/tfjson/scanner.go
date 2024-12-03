package tfjson

import (
	"context"
	"fmt"
	"io"
	"io/fs"

	"go.khulnasoft.com/pkg/iac/scan"
	"go.khulnasoft.com/pkg/iac/scanners/options"
	"go.khulnasoft.com/pkg/iac/scanners/terraform"
	"go.khulnasoft.com/pkg/iac/scanners/terraformplan/tfjson/parser"
	"go.khulnasoft.com/pkg/log"
)

type Scanner struct {
	parser    *parser.Parser
	logger    *log.Logger
	options   []options.ScannerOption
	tfScanner *terraform.Scanner
}

func (s *Scanner) Name() string {
	return "Terraform Plan JSON"
}

func (s *Scanner) ScanFS(ctx context.Context, fsys fs.FS, dir string) (scan.Results, error) {

	var results scan.Results

	walkFn := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		res, err := s.ScanFile(path, fsys)
		if err != nil {
			return fmt.Errorf("failed to scan %s: %w", path, err)
		}

		results = append(results, res...)
		return nil
	}

	if err := fs.WalkDir(fsys, dir, walkFn); err != nil {
		return nil, err
	}

	return results, nil
}

func New(opts ...options.ScannerOption) *Scanner {
	scanner := &Scanner{
		options:   opts,
		logger:    log.WithPrefix("tfjson scanner"),
		parser:    parser.New(),
		tfScanner: terraform.New(opts...),
	}

	return scanner
}

func (s *Scanner) ScanFile(filepath string, fsys fs.FS) (scan.Results, error) {

	s.logger.Debug("Scanning file", log.FilePath(filepath))
	file, err := fsys.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return s.Scan(file)
}

func (s *Scanner) Scan(reader io.Reader) (scan.Results, error) {

	planFile, err := s.parser.Parse(reader)
	if err != nil {
		return nil, err
	}

	planFS, err := planFile.ToFS()
	if err != nil {
		return nil, fmt.Errorf("failed to convert plan to FS: %w", err)
	}

	return s.tfScanner.ScanFS(context.TODO(), planFS, ".")
}
