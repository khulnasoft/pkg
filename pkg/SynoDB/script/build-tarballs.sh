#!/bin/bash

version=0.0.0

build_tarball () {
  os=$1
  arch=$2
  program=$3
  target=$os-$arch
  GOARCH=$arch GOOS=$os go build -o build/synodb-$os-$arch/$program main.go
  tar -C build -czvf "synodb-$version-$target.tar.gz" "synodb-$target/"
}

build_tarball darwin amd64 synodb
build_tarball darwin arm64 synodb
build_tarball linux amd64 synodb
build_tarball windows amd64 synodb.exe
