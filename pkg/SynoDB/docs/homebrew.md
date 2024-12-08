## Updating the Homebrew package

Homebrew package is automatically updated by GitHub Actions on synodb repository.
Those GitHub Actions are triggered by a new tag being pushed to the repository.

For example:

```console
git tag v0.1.3
git push --tags
```

## Setup details

Synodb Homebrew Tap is stored in [homebrew-tap](https://github.com/khulnasoft/homebrew-tap) repository.

There's a `ACCESS_TOKEN_TO_TAP` GitHub personal access token that has read/write access to Content and Actions homebrew-tap repositories.
It will expire on Jul 15 2024.
It is used by GitHub Actions in synodb repository to give them access to both synodb and homebrew-tap repositories.

[GoReleaser](https://github.com/goreleaser/goreleaser) is used to package everything up.
[GoReleaser GitHub Actions](https://github.com/goreleaser/goreleaser-action) are used for CI.

To install run:
```console
brew install khulnasoft/tap/synodb
```
