// Package test checks that the code generated by bpf2go conforms to a
// specific API.
package test

//go:generate go run github.com/khulnasoft/gbpf/cmd/bpf2go test ../testdata/minimal.c