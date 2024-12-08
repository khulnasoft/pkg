# tunnel-plugin-attest
Publish SBOM attestation

## Install

```
$ tunnel plugin install github.com/khulnasoft/tunnel-plugin-attest
```

## Usage

```
A Tunnel plugin that publish SBOM attestation

Usage:
  attest [flags]

Examples:
  tunnel attest --type PREDICATE_TYPE --predicate PREDICATE_PATH BLOB_PATH
  # Publish SBOM attestation
  tunnel attest --type cyclonedx --predicate ./sbom.cdx.json ./my-executable

Flags:
  -h, --help               help for attest
      --predicate string   specify the predicate file path
      --type string        specify the predicate type(cyclonedx)
```
