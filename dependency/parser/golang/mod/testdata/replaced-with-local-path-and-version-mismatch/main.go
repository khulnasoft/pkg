package main

import (
	"log"

	"go.khulnasoft.com/pkg/dependency/parser/golang/mod"
)

func main() {
	if _, err := mod.Parse(nil); err != nil {
		log.Fatal(err)
	}
}
