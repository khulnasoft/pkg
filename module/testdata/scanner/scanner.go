//go:generate tinygo build -o scanner.wasm -scheduler=none -target=wasi --no-debug scanner.go
//go:build tinygo.wasm

package main

import (
	"go.khulnasoft.com/pkg/module/api"
	"go.khulnasoft.com/pkg/module/serialize"
	"go.khulnasoft.com/pkg/module/wasm"
)

const (
	moduleVersion = 2
	moduleName    = "scanner"
)

func main() {
	wasm.RegisterModule(PostScannerModule{})
}

type PostScannerModule struct{}

func (PostScannerModule) Version() int {
	return moduleVersion
}

func (PostScannerModule) Name() string {
	return moduleName
}

func (PostScannerModule) PostScanSpec() serialize.PostScanSpec {
	return serialize.PostScanSpec{
		Action: api.ActionInsert, // Add new vulnerabilities
	}
}

func (PostScannerModule) PostScan(_ serialize.Results) (serialize.Results, error) {
	return nil, nil
}
