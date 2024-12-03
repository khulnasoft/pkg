package main

import (
	"log"

	"github.com/khulnasoft-lab/go-version/pkg/version"
)

func main() {
	if _, err := version.Parse("v0.1.2"); err != nil {
		log.Fatal(err)
	}
}
