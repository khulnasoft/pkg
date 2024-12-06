package all

import (
	_ "go.khulnasoft.com/pkg/fanal/analyzer/buildinfo"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/config/all"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/executable"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/imgconf/apk"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/imgconf/dockerfile"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/imgconf/secret"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/c/conan"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/conda/environment"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/conda/meta"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/dart/pub"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/dotnet/deps"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/dotnet/nuget"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/dotnet/packagesprops"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/elixir/mix"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/golang/binary"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/golang/mod"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/java/gradle"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/java/jar"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/java/pom"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/java/sbt"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/julia/pkg"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/nodejs/npm"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/nodejs/pkg"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/nodejs/pnpm"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/nodejs/yarn"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/php/composer"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/python/packaging"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/python/pip"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/python/pipenv"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/python/poetry"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/ruby/bundler"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/ruby/gemspec"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/rust/binary"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/rust/cargo"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/swift/cocoapods"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/language/swift/swift"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/licensing"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/os/alpine"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/os/amazonlinux"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/os/debian"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/os/redhatbase"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/os/release"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/os/ubuntu"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/pkg/apk"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/pkg/dpkg"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/pkg/rpm"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/repo/apk"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/sbom"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/secret"
)
