# tunnel-plugin-referrer

Tunnel plugin for the OCI referrer

## Install

```
$ tunnel plugin install github.com/khulnasoft/tunnel-plugin-referrer
```

## Usage

### Putting the SBOM into the OCI registry

Put the SBOM into the OCI registry as a referrer.
You need the write permission for the target repository.
```
# CycloneDX
$ tunnel image -q -f cyclonedx YOUR_IMAGE | tunnel referrer put

# SPDX
$ tunnel image -q -f spdx-json YOUR_IMAGE | tunnel referrer put
```

You can also upload by specifying a file.
```
$ tunnel image -q -f cyclonedx YOUR_IMAGE > sbom.cdx.json
$ tunnel referrer put -f sbom.cdx.json
```

### Putting the SARIF into the OCI registry
```
# NOTICE: Since Tunnel v0.38.3 doesn’t write subject identification in SARIF, you need to explicitly state the subject.
$ tunnel image -q -f sarif YOUR_IMAGE | tunnel referrer put --subject YOUR_IMAGE
```

### Putting the Vulnerability info into the OCI registry
```
$ tunnel image -q -f cosign-vuln YOUR_IMAGE | tunnel referrer put
```

### Listing referrers
```
$ tunnel referrer list localhost:5002/demo:app --type cyclonedx
Subject:                                     localhost:5002/demo:app
Referrers:                                    
  Digest:                                    sha256:5b0306dbbe8252a9efaa7a5177d5ae313441aa85c3d994b2183726303fe52101
  Reference:                                 localhost:5002/demo@sha256:5b0306dbbe8252a9efaa7a5177d5ae313441aa85c3d994b2183726303fe52101
  MediaType:                                 application/vnd.oci.image.manifest.v1+json
  ArtifactType:                              application/vnd.cyclonedx+json
  Annotations:                               
    created-by:                              tunnel
    org.opencontainers.artifact.created:     2023-04-11T19:21:29+09:00
    org.opencontainers.artifact.description: CycloneDX JSON SBOM


$ tunnel referrer list localhost:5002/demo:app --format table
DIGEST  TYPE      ANNOTATIONS       DESCRIPTION         CREATED
5b0306d cyclonedx created-by=tunnel  CycloneDX JSON SBOM 22 hours ago
771989f spdx-json created-by=tunnel  SPDX JSON SBOM      22 hours ago
83542d1 sarif     created-by=tunnel  SARIF               18 hours ago
8b9f058 sarif     created-by=tunnel  SARIF               15 hours ago
```

### Getting the artifact

```
$ tunnel referrer get localhost:5002/demo:app --digest 5b0306d | head
{
  "bomFormat": "CycloneDX",
  "specVersion": "1.4",
  "serialNumber": "urn:uuid:2512b7b2-ed10-4526-9b64-dc32869eb401",
  "version": 1,
  "metadata": {
    "timestamp": "2023-04-11T10:21:29+00:00",
    "tools": [
      {
        "vendor": "khulnasoft",

```

### Listing referrers recursively

```
$ tunnel referrer tree localhost:500/demo:app
Subject: localhost:5002/demo:app

a5cb013
├── 5b0306d: application/vnd.cyclonedx+json
│   └── c248e29: application/vnd.cncf.notary.signature
├── 771989f: application/spdx+json
├── 83542d1: application/sarif+json
└── 8b9f058: application/sarif+json
```
