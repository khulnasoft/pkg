<div align="center">
  <a href="https://github.com/khulnasoft/codebase">
    <img alt="codebase" src="https://raw.githubusercontent.com/haya14busa/i/d598ed7dc49fefb0018e422e4c43e5ab8f207a6b/khulnasoft/codebase.logo.png">
  </a>
</div>

<h2 align="center">
  codebase - A code review dog who keeps your codebase healthy.
</h2>

<div align="center">
  <a href="./LICENSE">
    <img alt="LICENSE" src="https://img.shields.io/badge/license-MIT-blue.svg?maxAge=43200">
  </a>
  <a href="https://godoc.org/github.com/khulnasoft/codebase">
    <img alt="GoDoc" src="https://img.shields.io/badge/godoc-reference-4F73B3.svg?label=godoc.org&maxAge=43200&logo=go">
  </a>
  <a href="./CHANGELOG.md">
    <img alt="releases" src="https://img.shields.io/github/release/khulnasoft/codebase.svg?logo=github">
  </a>
  <a href="https://github.com/khulnasoft/nightly">
    <img alt="nightly releases" src="https://img.shields.io/github/v/release/codebase/nightly.svg?logo=github">
  </a>
</div>

<div align="center">
  <a href="https://github.com/khulnasoft/codebase/actions?query=workflow%3AGo+event%3Apush+branch%3Amaster">
    <img alt="GitHub Actions" src="https://github.com/khulnasoft/codebase/workflows/Go/badge.svg">
  </a>
  <a href="https://github.com/khulnasoft/codebase/actions?query=workflow%3Acodebase+event%3Apush+branch%3Amaster">
    <img alt="codebase" src="https://github.com/khulnasoft/codebase/workflows/codebase/badge.svg?branch=master&event=push">
  </a>
  <a href="https://github.com/khulnasoft/codebase/actions?query=workflow%3Arelease">
    <img alt="release" src="https://github.com/khulnasoft/codebase/workflows/release/badge.svg">
  </a>
  <a href="https://travis-ci.org/khulnasoft/codebase"><img alt="Travis Status" src="https://img.shields.io/travis/khulnasoft/codebase/master.svg?label=Travis&logo=travis"></a>
  <a href="https://circleci.com/gh/khulnasoft/codebase"><img alt="CircleCI Status" src="http://img.shields.io/circleci/build/github/khulnasoft/codebase/master.svg?label=CircleCI&logo=circleci"></a>
  <a href="https://codecov.io/github/khulnasoft/codebase"><img alt="Coverage Status" src="https://img.shields.io/codecov/c/github/khulnasoft/codebase/master.svg?logo=codecov"></a>
</div>

<div align="center">
  <a href="https://gitlab.com/khulnasoft/codebase/pipelines">
    <img alt="GitLab Supported" src="https://img.shields.io/badge/GitLab%20-Supported-fc6d26?logo=gitlab">
  </a>
  <a href="https://github.com/haya14busa/action-bumpr">
    <img alt="action-bumpr supported" src="https://img.shields.io/badge/bumpr-supported-ff69b4?logo=github&link=https://github.com/haya14busa/action-bumpr">
  </a>
  <a href="https://github.com/khulnasoft/.github/blob/master/CODE_OF_CONDUCT.md">
    <img alt="Contributor Covenant" src="https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg">
  </a>
  <a href="https://somsubhra.github.io/github-release-stats/?username=codebase&repository=codebase&per_page=30">
    <img alt="GitHub Releases Stats" src="https://img.shields.io/github/downloads/khulnasoft/codebase/total.svg?logo=github">
  </a>
  <a href="https://starchart.cc/khulnasoft/codebase"><img alt="Stars" src="https://img.shields.io/github/stars/khulnasoft/codebase.svg?style=social"></a>
</div>
<br />

codebase provides a way to post review comments to code hosting services,
such as GitHub, automatically by integrating with any linter tools with ease.
It uses an output of lint tools and posts them as a comment if findings are in
the diff of patches to review.

codebase also supports running in the local environment to filter the output of lint tools
by diff.

[design doc](https://docs.google.com/document/d/1mGOX19SSqRowWGbXieBfGPtLnM0BdTkIc9JelTiu6wA/edit?usp=sharing)

## Table of Contents

- [Installation](#installation)
- [Input Format](#input-format)
  * ['errorformat'](#errorformat)
  * [Available pre-defined 'errorformat'](#available-pre-defined-errorformat)
  * [Codebase Diagnostic Format (RDFormat)](#codebase-diagnostic-format-rdformat)
  * [Diff](#diff)
  * [checkstyle format](#checkstyle-format)
  * [SARIF format](#sarif-format)
- [Code Suggestions](#code-suggestions)
- [codebase config file](#codebase-config-file)
- [Reporters](#reporters)
  * [Reporter: Local (-reporter=local) [default]](#reporter-local--reporterlocal-default)
  * [Reporter: GitHub Checks (-reporter=github-pr-check)](#reporter-github-checks--reportergithub-pr-check)
  * [Reporter: GitHub Checks (-reporter=github-check)](#reporter-github-checks--reportergithub-check)
  * [Reporter: GitHub PullRequest review comment (-reporter=github-pr-reviewe)](#reporter-github-pullrequest-review-comment--reportergithub-pr-reviewe)
  * [Reporter: GitLab MergeRequest discussions (-reporter=gitlab-mr-discussion)](#reporter-gitlab-mergerequest-discussions--reportergitlab-mr-discussion)
  * [Reporter: GitLab MergeRequest commit (-reporter=gitlab-mr-commit)](#reporter-gitlab-mergerequest-commit--reportergitlab-mr-commit)
  * [Reporter: Bitbucket Code Insights Reports (-reporter=bitbucket-code-report)](#reporter-bitbucket-code-insights-reports--reporterbitbucket-code-report)
- [Supported CI services](#supported-ci-services)
  * [GitHub Actions](#github-actions)
  * [Travis CI](#travis-ci)
  * [Circle CI](#circle-ci)
  * [GitLab CI](#gitlab-ci)
  * [Bitbucket Pipelines](#bitbucket-pipelines)
  * [Common (Jenkins, local, etc...)](#common-jenkins-local-etc)
    + [Jenkins with GitHub pull request builder plugin](#jenkins-with-github-pull-request-builder-plugin)
- [Exit codes](#exit-codes)
- [Filter mode](#filter-mode)
- [Articles](#articles)

[![github-pr-check sample](https://user-images.githubusercontent.com/3797062/40884858-6efd82a0-6756-11e8-9f1a-c6af4f920fb0.png)](https://github.com/khulnasoft/codebase/pull/131/checks)
![comment in pull-request](https://user-images.githubusercontent.com/3797062/40941822-1d775064-6887-11e8-98e9-4775d37d47f8.png)
![commit status](https://user-images.githubusercontent.com/3797062/40941738-d62acb0a-6886-11e8-858d-7b97aded2a42.png)
[![sample-comment.png](https://raw.githubusercontent.com/haya14busa/i/dc0ccb1e110515ea407c146d99b749018db05c45/codebase/sample-comment.png)](https://github.com/khulnasoft/codebase/pull/24#discussion_r84599728)
![codebase-local-demo.gif](https://raw.githubusercontent.com/haya14busa/i/dc0ccb1e110515ea407c146d99b749018db05c45/khulnasoft/codebase-local-demo.gif)

## Installation

```shell
# Install the latest version. (Install it into ./bin/ by default).
$ curl -sfL https://raw.githubusercontent.com/khulnasoft/codebase/master/install.sh | sh -s

# Specify installation directory ($(go env GOPATH)/bin/) and version.
$ curl -sfL https://raw.githubusercontent.com/khulnasoft/codebase/master/install.sh | sh -s -- -b $(go env GOPATH)/bin [vX.Y.Z]

# In alpine linux (as it does not come with curl by default)
$ wget -O - -q https://raw.githubusercontent.com/khulnasoft/codebase/master/install.sh | sh -s [vX.Y.Z]
```

### Nightly releases

You can also use [nightly codebase release](https://github.com/khulnasoft/nightly)
to try the latest codebase improvements every day!

```shell
$ curl -sfL https://raw.githubusercontent.com/khulnasoft/nightly/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```

### GitHub Action: [codebase/action-setup](https://github.com/khulnasoft/action-setup)

```yaml
steps:
- uses: codebase/action-setup@v1
  with:
    codebase_version: latest # Optional. [latest,nightly,v.X.Y.Z]
```

### homebrew / linuxbrew
You can also install codebase using brew:

```shell
$ brew install codebase/tap/codebase
$ brew upgrade codebase/tap/codebase
```

### [Scoop](https://scoop.sh/) on Windows

```
> scoop install codebase
```

### Build with go install

```shell
$ go install github.com/khulnasoft/codebase/cmd/codebase@latest
```

## Input Format

### 'errorformat'

codebase accepts any compiler or linter result from stdin and parses it with
scan-f like [**'errorformat'**](https://github.com/khulnasoft/codebase/errorformat),
which is the port of Vim's [errorformat](https://vim-jp.org/vimdoc-en/quickfix.html#error-file-format)
feature.

For example, if the result format is `{file}:{line number}:{column number}: {message}`,
errorformat should be `%f:%l:%c: %m` and you can pass it as `-efm` arguments.

```shell
$ golint ./...
comment_iowriter.go:11:6: exported type CommentWriter should have comment or be unexported
$ golint ./... | codebase -efm="%f:%l:%c: %m" -diff="git diff FETCH_HEAD"
```

| name | description |
| ---- | ----------- |
| %f | file name |
| %l | line number |
| %c | column number |
| %m | error message |
| %% | the single '%' character |
| ... | ... |

Please see [codebase/errorformat](https://github.com/khulnasoft/codebase/errorformat)
and [:h errorformat](https://vim-jp.org/vimdoc-en/quickfix.html#error-file-format)
if you want to deal with a more complex output. 'errorformat' can handle more
complex output like a multi-line error message.

You can also try errorformat on [the Playground](https://codebase.github.io/errorformat-playground/)!

With this 'errorformat' feature, codebase can support any tool output with ease.

### Available pre-defined 'errorformat'

But, you don't have to write 'errorformat' in many cases. codebase supports
pre-defined errorformat for major tools.

You can find available errorformat name by `codebase -list` and you can use it
with `-f={name}`.

```shell
$ codebase -list
golint          linter for Go source code                                       - https://github.com/golang/lint
govet           Vet examines Go source code and reports suspicious problems     - https://golang.org/cmd/vet/
sbt             the interactive build tool                                      - http://www.scala-sbt.org/
...
```

```shell
$ golint ./... | codebase -f=golint -diff="git diff FETCH_HEAD"
```

You can add supported pre-defined 'errorformat' by contributing to [codebase/errorformat](https://github.com/khulnasoft/codebase/errorformat)

### Codebase Diagnostic Format (RDFormat)

codebase supports [Codebase Diagnostic Format (RDFormat)](./proto/rdf/) as a
generic diagnostic format and it supports both [rdjson](./proto/rdf/#rdjson) and
[rdjsonl](./proto/rdf/#rdjsonl) formats.

This rdformat supports rich features like multiline ranged comments, severity,
rule code with URL, and [code suggestions](#code-suggestions).

```shell
$ <linter> | <convert-to-rdjson> | codebase -f=rdjson -reporter=github-pr-reviewe
# or
$ <linter> | <convert-to-rdjsonl> | codebase -f=rdjsonl -reporter=github-pr-reviewe
```

#### Example: ESLint with RDFormat 

![eslint codebase rdjson demo](https://user-images.githubusercontent.com/3797062/97085944-87233a80-165b-11eb-94a8-0a47d5e24905.png)

You can use [eslint-formatter-rdjson](https://www.npmjs.com/package/eslint-formatter-rdjson)
to output `rdjson` as eslint output format.

```shell
$ npm install --save-dev eslint-formatter-rdjson
$ eslint -f rdjson . | codebase -f=rdjson -reporter=github-pr-reviewe
```

Or you can also use [codebase/action-eslint](https://github.com/khulnasoft/action-eslint) for GitHub Actions.

### Diff

![codebase with gofmt example](https://user-images.githubusercontent.com/3797062/89168305-a3ad5a80-d5b7-11ea-8939-be7ac1976d30.png)

codebase supports diff (unified format) as an input format especially useful
for [code suggestions](#code-suggestions).
codebase can integrate with any code suggestions tools or formatters to report suggestions.

`-f.diff.strip`: option for `-f=diff`: strip NUM leading components from diff file names (equivalent to 'patch -p') (default is 1 for git diff) (default 1)

```shell
$ <any-code-fixer/formatter> # e.g. eslint --fix, gofmt
$ TMPFILE=$(mktemp)
$ git diff >"${TMPFILE}"
$ git stash -u && git stash drop
$ codebase -f=diff -f.diff.strip=1 -reporter=github-pr-reviewe < "${TMPFILE}"
```

Or you can also use [codebase/action-suggester](https://github.com/khulnasoft/action-suggester) for GitHub Actions.

If diagnostic tools support diff output format, you can pipe the diff directly.

```shell
$ gofmt -s -d . | codebase -name="gofmt" -f=diff -f.diff.strip=0 -reporter=github-pr-reviewe
$ shellcheck -f diff $(shfmt -f .) | codebase -f=diff
```

### checkstyle format

codebase also accepts [checkstyle XML format](http://checkstyle.sourceforge.net/) as well.
If the linter supports checkstyle format as a report format, you can use
-f=checkstyle instead of using 'errorformat'.

```shell
# Local
$ eslint -f checkstyle . | codebase -f=checkstyle -diff="git diff"

# CI (overwrite tool name which is shown in review comment by -name arg)
$ eslint -f checkstyle . | codebase -f=checkstyle -name="eslint" -reporter=github-check
```

Also, if you want to pass other Json/XML/etc... format to codebase, you can write a converter.

```shell
$ <linter> | <convert-to-checkstyle> | codebase -f=checkstyle -name="<linter>" -reporter=github-pr-check
```

### SARIF format

codebase supports [SARIF 2.1.0 JSON format](https://sarifweb.azurewebsites.net/).
You can use codebase with -f=sarif option.

```shell
# Local
$ eslint -f @microsoft/eslint-formatter-sarif . | codebase -f=sarif -diff="git diff"
````

## Code Suggestions

![eslint codebase suggestion demo](https://user-images.githubusercontent.com/3797062/97085944-87233a80-165b-11eb-94a8-0a47d5e24905.png)
![codebase with gofmt example](https://user-images.githubusercontent.com/3797062/89168305-a3ad5a80-d5b7-11ea-8939-be7ac1976d30.png)

codebase supports *code suggestions* feature with [rdformat](#codebase-diagnostic-format-rdformat) or [diff](#diff) input.
You can also use [codebase/action-suggester](https://github.com/khulnasoft/action-suggester) for GitHub Actions.

codebase can suggest code changes along with diagnostic results if a diagnostic tool supports code suggestions data.
You can integrate codebase with any code fixing tools and any code formatter with [diff](#diff) input as well.

### Code Suggestions Support Table
Note that not all reporters provide support for code suggestions.

| `-reporter`     | Suggestion support |
| ---------------------------- | ------- |
| **`local`**                  | NO [1]  |
| **`github-check`**           | NO [2]  |
| **`github-pr-check`**        | NO [2]  |
| **`github-pr-reviewe`**       | OK      |
| **`gitlab-mr-discussion`**   | NO [1]  |
| **`gitlab-mr-commit`**       | NO [2]  |
| **`gerrit-change-review`**   | NO [1]  |
| **`bitbucket-code-report`**  | NO [2]  |
| **`gitea-pr-review`**        | NO [2]  |

- [1] The reporter service supports the code suggestion feature, but codebase does not support it yet. See [#678](https://github.com/khulnasoft/codebase/issues/678) for the status.
- [2] The reporter service itself doesn't support the code suggestion feature.

## codebase config file

codebase can also be controlled via the .codebase.yml configuration file instead of "-f" or "-efm" arguments.

With .codebase.yml, you can run the same commands for both CI service and local
environment including editor integration with ease.

#### .codebase.yml

```yaml
runner:
  <tool-name>:
    cmd: <command> # (required)
    errorformat: # (optional if you use `format`)
      - <list of errorformat>
    format: <format-name> # (optional if you use `errorformat`. e.g. golint,rdjson,rdjsonl)
    name: <tool-name> # (optional. you can overwrite <tool-name> defined by runner key)
    level: <level> # (optional. same as -level flag. [info,warning,error])

  # examples
  golint:
    cmd: golint ./...
    errorformat:
      - "%f:%l:%c: %m"
    level: warning
  govet:
    cmd: go vet -all .
    format: govet
  your-awesome-linter:
    cmd: awesome-linter run
    format: rdjson
    name: AwesomeLinter
```

```shell
$ codebase -diff="git diff FETCH_HEAD"
project/run_test.go:61:28: [golint] error strings should not end with punctuation
project/run.go:57:18: [errcheck]        defer os.Setenv(name, os.Getenv(name))
project/run.go:58:12: [errcheck]        os.Setenv(name, "")
# You can use -runners to run only specified runners.
$ codebase -diff="git diff FETCH_HEAD" -runners=golint,govet
project/run_test.go:61:28: [golint] error strings should not end with punctuation
# You can use -conf to specify config file path.
$ codebase -conf=./.codebase.yml -reporter=github-pr-check
```

Output format for project config based run is one of the following formats.

- `<file>: [<tool name>] <message>`
- `<file>:<lnum>: [<tool name>] <message>`
- `<file>:<lnum>:<col>: [<tool name>] <message>`

## Reporters

codebase can report results both in the local environment and review services as
continuous integration.

### Reporter: Local (-reporter=local) [default]

codebase can find newly introduced findings by filtering linter results
using diff. You can pass the diff command as `-diff` arg.

```shell
$ golint ./... | codebase -f=golint -diff="git diff FETCH_HEAD"
```

### Reporter: GitHub Checks (-reporter=github-pr-check)

[![github-pr-check sample annotation with option 1](https://user-images.githubusercontent.com/3797062/64875597-65016f80-d688-11e9-843f-4679fb666f0d.png)](https://github.com/khulnasoft/codebase/pull/275/files#annotation_6177941961779419)
[![github-pr-check sample](https://user-images.githubusercontent.com/3797062/40884858-6efd82a0-6756-11e8-9f1a-c6af4f920fb0.png)](https://github.com/khulnasoft/codebase/pull/131/checks)

github-pr-check reporter reports results to [GitHub Checks](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/collaborating-on-repositories-with-code-quality-features/about-status-checks).

You can change the report level for this reporter by `level` field in [config
file](#codebase-config-file) or `-level` flag. You can control GitHub status
check results with this feature. (default: error)

| Level     | GitHub Status |
| --------- | ------------- |
| `info`    | neutral       |
| `warning` | neutral       |
| `error`   | failure       |

There are two options to use this reporter.

#### Option 1) Run codebase from GitHub Actions w/ secrets.GITHUB_TOKEN

Example: [.github/workflows/codebase.yml](.github/workflows/codebase.yml)

```yaml
- name: Run codebase
  env:
    CODEBASE_GITHUB_API_TOKEN: ${{ secrets.GITHUB_TOKEN }}
  run: |
    golint ./... | codebase -f=golint -reporter=github-pr-check
```

See [GitHub Actions](#github-actions) section too. You can also use public
codebase GitHub Actions.

#### Option 2) Install codebase GitHub Apps
codebase CLI sends a request to codebase GitHub App server and the server post
results as GitHub Checks, because Check API is only supported for GitHub App and
GitHub Actions.

1. Install codebase Apps. https://github.com/apps/codebase
2. Set `CODEBASE_TOKEN` or run codebase CLI in trusted CI providers.
  - Get token from `https://codebase.app/gh/{owner}/{repo-name}`.

```shell
$ export CODEBASE_TOKEN="<token>"
$ codebase -reporter=github-pr-check
```

Note: Token is not required if you run codebase in Travis or AppVeyor.

*Caution*

As described above, github-pr-check reporter with Option 2 depends on
codebase GitHub App server.
The server is running with haya14busa's pocket money for now and I may break
things, so I cannot ensure that the server is running 24h and 365 days.

**UPDATE:** Started getting support by [opencollective](https://opencollective.com/codebase)
and GitHub sponsor.
See [Supporting codebase](#supporting-codebase)

You can use github-pr-reviewe reporter or use run codebase under GitHub Actions
if you don't want to depend on codebase server.

### Reporter: GitHub Checks (-reporter=github-check)

It's basically the same as `-reporter=github-pr-check` except it works not only for
Pull Request but also for commit.

[![sample comment outside diff](https://user-images.githubusercontent.com/3797062/69917921-e0680580-14ae-11ea-9a56-de9e3cbac005.png)](https://github.com/khulnasoft/codebase/pull/364/files)

You can create [codebase badge](#codebase-badge-) for this reporter.

### Reporter: GitHub PullRequest review comment (-reporter=github-pr-reviewe)

[![sample-comment.png](https://raw.githubusercontent.com/haya14busa/i/dc0ccb1e110515ea407c146d99b749018db05c45/codebase/sample-comment.png)](https://github.com/khulnasoft/codebase/pull/24#discussion_r84599728)

github-pr-reviewe reporter reports results to GitHub PullRequest review comments
using GitHub Personal API Access Token.
[GitHub Enterprise](https://github.com/enterprise) is supported too.

- Go to https://github.com/settings/tokens and generate a new API token.
- Check `repo` for private repositories or `public_repo` for public repositories.

```shell
$ export CODEBASE_GITHUB_API_TOKEN="<token>"
$ codebase -reporter=github-pr-reviewe
```

For GitHub Enterprise, set the API endpoint by an environment variable.

```shell
$ export GITHUB_API="https://example.githubenterprise.com/api/v3/"
$ export CODEBASE_INSECURE_SKIP_VERIFY=true # set this as you need to skip verifying SSL
```

See [GitHub Actions](#github-actions) section too if you can use GitHub
Actions. You can also use public codebase GitHub Actions.

### Reporter: GitLab MergeRequest discussions (-reporter=gitlab-mr-discussion)

[![gitlab-mr-discussion sample](https://user-images.githubusercontent.com/3797062/41810718-f91bc540-773d-11e8-8598-fbc09ce9b1c7.png)](https://gitlab.com/khulnasoft/codebase/-/merge_requests/113#note_83411103)

Required GitLab version: >= v10.8.0

gitlab-mr-discussion reporter reports results to GitLab MergeRequest discussions using
GitLab Personal API Access token.
Get the token with `api` scope from https://gitlab.com/profile/personal_access_tokens.

```shell
$ export CODEBASE_GITLAB_API_TOKEN="<token>"
$ codebase -reporter=gitlab-mr-discussion
```

The `CI_API_V4_URL` environment variable, defined automatically by Gitlab CI (v11.7 onwards), will be used to find out the Gitlab API URL.

Alternatively, `GITLAB_API` can also be defined, in which case it will take precedence over `CI_API_V4_URL`.

```shell
$ export GITLAB_API="https://example.gitlab.com/api/v4"
$ export CODEBASE_INSECURE_SKIP_VERIFY=true # set this as you need to skip verifying SSL
```

### Reporter: GitLab MergeRequest commit (-reporter=gitlab-mr-commit)

gitlab-mr-commit is similar to [gitlab-mr-discussion](#reporter-gitlab-mergerequest-discussions--reportergitlab-mr-discussion) reporter but reports results to each commit in GitLab MergeRequest.

gitlab-mr-discussion is recommended, but you can use gitlab-mr-commit reporter
if your GitLab version is under v10.8.0.

```shell
$ export CODEBASE_GITLAB_API_TOKEN="<token>"
$ codebase -reporter=gitlab-mr-commit
```

### Reporter: Gerrit Change review (-reporter=gerrit-change-review)

gerrit-change-review reporter reports results to Gerrit Change using Gerrit Rest APIs.

The reporter supports Basic Authentication and Git-cookie based authentication for reporting results.

Set `GERRIT_USERNAME` and `GERRIT_PASSWORD` environment variables for basic authentication, and put `GIT_GITCOOKIE_PATH` for git cookie-based authentication.

```shell
$ export GERRIT_CHANGE_ID=changeID
$ export GERRIT_REVISION_ID=revisionID
$ export GERRIT_BRANCH=master
$ export GERRIT_ADDRESS=http://<gerrit-host>:<gerrit-port>
$ codebase -reporter=gerrit-change-review
```

### Reporter: Bitbucket Code Insights Reports (-reporter=bitbucket-code-report)

[![bitbucket-code-report](https://user-images.githubusercontent.com/9948629/96770123-c138d600-13e8-11eb-8e46-250b4bb393bd.png)](https://bitbucket.org/Trane9991/codebase-example/pull-requests/1)
[![bitbucket-code-annotations](https://user-images.githubusercontent.com/9948629/97054896-5e813f00-158e-11eb-9ad7-f8d75489b8ba.png)](https://bitbucket.org/Trane9991/codebase-example/pull-requests/1)

bitbucket-code-report generates the annotated
[Bitbucket Code Insights](https://support.atlassian.com/bitbucket-cloud/docs/code-insights/) report.

For now, only the `no-filter` mode is supported, so the whole project is scanned on every run.
Reports are stored per commit and can be viewed per commit from Bitbucket Pipelines UI or
in Pull Request. In the Pull Request UI affected code lines will be annotated in the diff,
as well as you will be able to filter the annotations by **This pull request** or **All**.

If running from [Bitbucket Pipelines](#bitbucket-pipelines), no additional configuration is needed (even credentials).
If running locally or from some other CI system you would need to provide Bitbucket API credentials:

- For Basic Auth you need to set the following env variables:
    `BITBUCKET_USER` and `BITBUCKET_PASSWORD`
- For AccessToken Auth you need to set `BITBUCKET_ACCESS_TOKEN`

```shell
$ export BITBUCKET_USER="my_user"
$ export BITBUCKET_PASSWORD="my_password"
$ codebase -reporter=bitbucket-code-report
```

To post a report to the Bitbucket Server use `BITBUCKET_SERVER_URL` variable:
```shell
$ export BITBUCKET_USER="my_user"
$ export BITBUCKET_PASSWORD="my_password"
$ export BITBUCKET_SERVER_URL="https://bitbucket.my-company.com"
$ codebase -reporter=bitbucket-code-report
```

## Supported CI services

### [GitHub Actions](https://github.com/features/actions)

Example: [.github/workflows/codebase.yml](.github/workflows/codebase.yml)

```yaml
name: codebase
on: [pull_request]
jobs:
  codebase:
    name: codebase
    runs-on: ubuntu-latest
    steps:
      # ...
      - uses: codebase/action-setup@v1
        with:
          codebase_version: latest # Optional. [latest,nightly,v.X.Y.Z]
      - name: Run codebase
        env:
          CODEBASE_GITHUB_API_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        run: |
          codebase -reporter=github-pr-check -runners=golint,govet
          # or
          codebase -reporter=github-pr-reviewe -runners=golint,govet
```

<details>
<summary><strong>Example (github-check reporter):</strong></summary>

[.github/workflows/codebase](.github/workflows/codebase.yml)

Only `github-check` reporter can run on the push event too.

```yaml
name: codebase (github-check)
on:
  push:
    branches:
      - master
  pull_request:

jobs:
  codebase:
    name: codebase
    runs-on: ubuntu-latest
    steps:
      # ...
      - name: Run codebase
        env:
          CODEBASE_GITHUB_API_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        run: |
          codebase -reporter=github-check -runners=golint,govet
```

</details>

#### Public Codebase GitHub Actions
You can use public GitHub Actions to start using codebase with ease! :tada: :arrow_forward: :tada:

- Common
  - [codebase/action-misspell](https://github.com/khulnasoft/action-misspell) - Run [misspell](https://github.com/client9/misspell).
  - [EPMatt/codebase-action-prettier](https://github.com/EPMatt/codebase-action-prettier) - Run [Prettier](https://prettier.io/).
- Text
  - [codebase/action-alex](https://github.com/khulnasoft/action-alex) - Run [alex](https://github.com/get-alex/alex) which catches insensitive, inconsiderate writing. (e.g. master/slave)
  - [codebase/action-languagetool](https://github.com/khulnasoft/action-languagetool) - Run [languagetool](https://github.com/languagetool-org/languagetool).
  - [tsuyoshicho/action-textlint](https://github.com/tsuyoshicho/action-textlint) - Run [textlint](https://github.com/textlint/textlint)
  - [tsuyoshicho/action-redpen](https://github.com/tsuyoshicho/action-redpen) - Run [redpen](https://github.com/redpen-cc/redpen)
- Markdown
  - [codebase/action-markdownlint](https://github.com/khulnasoft/action-markdownlint) - Run [markdownlint](https://github.com/DavidAnson/markdownlint)
- Docker
  - [codebase/action-hadolint](https://github.com/khulnasoft/action-hadolint) - Run [hadolint](https://github.com/hadolint/hadolint) to lint `Dockerfile`.
- Env
  - [dotenv-linter/action-dotenv-linter](https://github.com/dotenv-linter/action-dotenv-linter) - Run [dotenv-linter](https://github.com/dotenv-linter/dotenv-linter) to lint `.env` files.
- Shell script
  - [codebase/action-shellcheck](https://github.com/khulnasoft/action-shellcheck) - Run [shellcheck](https://github.com/koalaman/shellcheck).
  - [codebase/action-shfmt](https://github.com/khulnasoft/action-shfmt) - Run [shfmt](https://github.com/mvdan/sh).
- Go
  - [codebase/action-staticcheck](https://github.com/khulnasoft/action-staticcheck) - Run [staticcheck](https://staticcheck.io/).
  - [codebase/action-golangci-lint](https://github.com/khulnasoft/action-golangci-lint) - Run [golangci-lint](https://github.com/golangci/golangci-lint) and supported linters individually by golangci-lint.
- JavaScript
  - [codebase/action-eslint](https://github.com/khulnasoft/action-eslint) - Run [eslint](https://github.com/eslint/eslint).
- TypeScript
  - [EPMatt/codebase-action-tsc](https://github.com/EPMatt/codebase-action-tsc) - Run [tsc](https://www.typescriptlang.org/docs/handbook/compiler-options.html).
- CSS
  - [codebase/action-stylelint](https://github.com/khulnasoft/action-stylelint) - Run [stylelint](https://github.com/stylelint/stylelint).
- Vim script
  - [codebase/action-vint](https://github.com/khulnasoft/action-vint) - Run [vint](https://github.com/Vimjas/vint).
  - [tsuyoshicho/action-vimlint](https://github.com/tsuyoshicho/action-vimlint) - Run [vim-vimlint](https://github.com/syngan/vim-vimlint)
- Terraform
  - [codebase/action-tflint](https://github.com/khulnasoft/action-tflint) - Run [tflint](https://github.com/terraform-linters/tflint).
- YAML
  - [codebase/action-yamllint](https://github.com/khulnasoft/action-yamllint) - Run [yamllint](https://github.com/adrienverge/yamllint).
- Ruby
  - [codebase/action-brakeman](https://github.com/khulnasoft/action-brakeman) - Run [brakeman](https://github.com/presidentbeef/brakeman).
  - [codebase/action-reek](https://github.com/khulnasoft/action-reek) - Run [reek](https://github.com/troessner/reek).
  - [codebase/action-rubocop](https://github.com/khulnasoft/action-rubocop) - Run [rubocop](https://github.com/rubocop/rubocop).
  - [vk26/action-fasterer](https://github.com/vk26/action-fasterer) - Run [fasterer](https://github.com/DamirSvrtan/fasterer).
  - [PrintReleaf/action-standardrb](https://github.com/PrintReleaf/action-standardrb) - Run [standardrb](https://github.com/standardrb/standard).
  - [tk0miya/action-erblint](https://github.com/tk0miya/action-erblint) - Run [erb-lint](https://github.com/Shopify/erb-lint)
  - [tk0miya/action-steep](https://github.com/tk0miya/action-steep) - Run [steep](https://github.com/soutaro/steep)
  - [blooper05/action-rails_best_practices](https://github.com/blooper05/action-rails_best_practices) - Run [rails_best_practices](https://github.com/flyerhzm/rails_best_practices)
  - [tomferreira/action-bundler-audit](https://github.com/tomferreira/action-bundler-audit) - Run [bundler-audit](https://github.com/rubysec/bundler-audit)
- Python
  - [wemake-python-styleguide](https://github.com/wemake-services/wemake-python-styleguide) - Run wemake-python-styleguide
  - [tsuyoshicho/action-mypy](https://github.com/tsuyoshicho/action-mypy) - Run [mypy](https://pypi.org/project/mypy/)
  - [jordemort/action-pyright](https://github.com/jordemort/action-pyright) - Run [pyright](https://github.com/Microsoft/pyright)
  - [dciborow/action-pylint](https://github.com/dciborow/action-pylint) - Run [pylint](https://github.com/pylint-dev/pylint)
  - [codebase/action-black](https://github.com/khulnasoft/action-black) - Run [black](https://github.com/psf/black)
- Kotlin
  - [ScaCap/action-ktlint](https://github.com/ScaCap/action-ktlint) - Run [ktlint](https://ktlint.github.io/).
- Android Lint
  - [dvdandroid/action-android-lint](https://github.com/DVDAndroid/action-android-lint) - Run [Android Lint](https://developer.android.com/studio/write/lint)
- Ansible
  - [codebase/action-ansiblelint](https://github.com/khulnasoft/action-ansiblelint) - Run [ansible-lint](https://github.com/ansible/ansible-lint)
- GitHub Actions
  - [codebase/action-actionlint](https://github.com/khulnasoft/action-actionlint) - Run [actionlint](https://github.com/rhysd/actionlint)
- Protocol Buffers
  - [yoheimuta/action-protolint](https://github.com/yoheimuta/action-protolint) - Run [protolint](https://github.com/yoheimuta/protolint)
  
... and more on [GitHub Marketplace](https://github.com/marketplace?utf8=✓&type=actions&query=codebase).

Missing actions? Check out [codebase/action-template](https://github.com/khulnasoft/action-template) and create a new codebase action!

Please open a Pull Request to add your created codebase actions here :sparkles:.
I can also put your repositories under codebase org and co-maintain the actions.
Example: [action-tflint](https://github.com/khulnasoft/codebase/issues/322).

#### Graceful Degradation for Pull Requests from forked repositories

![Graceful Degradation example](https://user-images.githubusercontent.com/3797062/71781334-e2266b00-3010-11ea-8a38-dee6e30c8162.png)

`GITHUB_TOKEN` for Pull Requests from a forked repository doesn't have write
access to Check API nor Review API due to [GitHub Actions
restriction](https://docs.github.com/en/actions/security-guides/automatic-token-authentication#permissions-for-the-github_token).

Instead, codebase uses [Logging commands of GitHub
Actions](https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions#set-an-error-message-error)
to post results as
[annotations](https://docs.github.com/en/rest/checks/runs#annotations-object)
similar to `github-pr-check` reporter.

Note that there is a limitation for annotations created by logging commands,
such as [max # of annotations per run](https://github.com/khulnasoft/codebase/issues/411#issuecomment-570893427).
You can check GitHub Actions log to see full results in such cases.

#### codebase badge [![codebase](https://github.com/khulnasoft/codebase/workflows/codebase/badge.svg?branch=master&event=push)](https://github.com/khulnasoft/codebase/actions?query=workflow%3Acodebase+event%3Apush+branch%3Amaster)

As [`github-check` reporter](#reporter-github-checks--reportergithub-pr-check) support running on commit, we can create codebase
[GitHub Action badge](https://docs.github.com/en/actions/using-workflows#adding-a-workflow-status-badge-to-your-repository)
to check the result against master commit for example. :tada:

Example:
```
<!-- Replace <OWNER> and <REPOSITORY>. It assumes workflow name is "codebase" -->
[![codebase](https://github.com/<OWNER>/<REPOSITORY>/workflows/codebase/badge.svg?branch=master&event=push)](https://github.com/<OWNER>/<REPOSITORY>/actions?query=workflow%3Acodebase+event%3Apush+branch%3Amaster)
```

### Travis CI

#### Travis CI (-reporter=github-pr-check)

If you use -reporter=github-pr-check in Travis CI, you don't need to set `CODEBASE_TOKEN`.

Example:

```yaml
install:
  - mkdir -p ~/bin/ && export PATH="~/bin/:$PATH"
  - curl -sfL https://raw.githubusercontent.com/khulnasoft/codebase/master/install.sh| sh -s -- -b ~/bin

script:
  - codebase -conf=.codebase.yml -reporter=github-pr-check
```

#### Travis CI (-reporter=github-pr-reviewe)

Store GitHub API token by [travis encryption keys](https://docs.travis-ci.com/user/encryption-keys/).

```shell
$ gem install travis
$ travis encrypt CODEBASE_GITHUB_API_TOKEN=<token> --add env.global
```
Example:

```yaml
env:
  global:
    - secure: <token>

install:
  - mkdir -p ~/bin/ && export PATH="~/bin/:$PATH"
  - curl -sfL https://raw.githubusercontent.com/khulnasoft/codebase/master/install.sh| sh -s -- -b ~/bin

script:
  - >-
    golint ./... | codebase -f=golint -reporter=github-pr-reviewe
```

Examples
- https://github.com/azu/textlint-codebase-example

### Circle CI

Store `CODEBASE_GITHUB_API_TOKEN` (or `CODEBASE_TOKEN` for github-pr-check) in
[Environment variables - CircleCI](https://circleci.com/docs/environment-variables/#setting-environment-variables-for-all-commands-without-adding-them-to-git)

#### .circleci/config.yml sample

```yaml
version: 2
jobs:
  build:
    docker:
      - image: golang:latest
    steps:
      - checkout
      - run: curl -sfL https://raw.githubusercontent.com/khulnasoft/codebase/master/install.sh| sh -s -- -b ./bin
      - run: go vet ./... 2>&1 | ./bin/codebase -f=govet -reporter=github-pr-reviewe

      # Deprecated: prefer GitHub Actions to use github-pr-check reporter.
      - run: go vet ./... 2>&1 | ./bin/codebase -f=govet -reporter=github-pr-check
```

### GitLab CI

Store `CODEBASE_GITLAB_API_TOKEN` in [GitLab CI variable](https://docs.gitlab.com/ee/ci/variables/#variables).

#### .gitlab-ci.yml sample

```yaml
codebase:
  script:
    - codebase -reporter=gitlab-mr-discussion
    # Or
    - codebase -reporter=gitlab-mr-commit
```

### Bitbucket Pipelines

No additional configuration is needed.

#### bitbucket-pipelines.yml sample

```yaml
pipelines:
  default:
    - step:
        name: Codebase
        image: golangci/golangci-lint:v1.31-alpine
        script:
          - wget -O - -q https://raw.githubusercontent.com/khulnasoft/codebase/master/install.sh | 
              sh -s -- -b $(go env GOPATH)/bin
          - golangci-lint run --out-format=line-number ./... | codebase -f=golangci-lint -reporter=bitbucket-code-report
```

### Common (Jenkins, local, etc...)

You can use codebase to post review comments from anywhere with following
environment variables.

| name | description |
| ---- | ----------- |
| `CI_PULL_REQUEST` | Pull Request number (e.g. 14) |
| `CI_COMMIT`       | SHA1 for the current build |
| `CI_REPO_OWNER`   | repository owner (e.g. "codebase" for https://github.com/khulnasoft/codebase/errorformat) |
| `CI_REPO_NAME`    | repository name (e.g. "errorformat" for https://github.com/khulnasoft/codebase/errorformat) |
| `CI_BRANCH`       | [optional] branch of the commit |

```shell
$ export CI_PULL_REQUEST=14
$ export CI_REPO_OWNER=haya14busa
$ export CI_REPO_NAME=codebase
$ export CI_COMMIT=$(git rev-parse HEAD)
```
and set a token if required.

```shell
$ CODEBASE_TOKEN="<token>"
$ CODEBASE_GITHUB_API_TOKEN="<token>"
$ CODEBASE_GITLAB_API_TOKEN="<token>"
```

If a CI service doesn't provide information such as Pull Request ID - codebase can guess it by a branch name and commit SHA.
Just pass the flag `guess`:

```shell
$ codebase -conf=.codebase.yml -reporter=github-pr-check -guess
```

#### Jenkins with GitHub pull request builder plugin
- [GitHub pull request builder plugin - Jenkins - Jenkins Wiki](https://wiki.jenkins-ci.org/display/JENKINS/GitHub+pull+request+builder+plugin)
- [Configuring a GitHub app account - Jenkins - CloudBees](https://docs.cloudbees.com/docs/cloudbees-ci/latest/cloud-admin-guide/github-app-auth) - required to use github-pr-check formatter without codebase server or GitHub actions.

```shell
$ export CI_PULL_REQUEST=${ghprbPullId}
$ export CI_REPO_OWNER=haya14busa
$ export CI_REPO_NAME=codebase
$ export CI_COMMIT=${ghprbActualCommit}
$ export CODEBASE_INSECURE_SKIP_VERIFY=true # set this as you need

# To submit via codebase server using github-pr-check reporter
$ CODEBASE_TOKEN="<token>" codebase -reporter=github-pr-check
# Or, to submit directly via API using github-pr-reviewe reporter
$ CODEBASE_GITHUB_API_TOKEN="<token>" codebase -reporter=github-pr-reviewe
# Or, to submit directly via API using github-pr-check reporter (requires GitHub App Account configured)
$ CODEBASE_SKIP_DOGHOUSE=true CODEBASE_GITHUB_API_TOKEN="<token>" codebase -reporter=github-pr-check
```

## Exit codes
By default codebase will return `0` as exit code even if it finds errors.
If `-fail-on-error` flag is passed, codebase exits with `1` when at least one error was found/reported.
This can be helpful when you are using it as a step in your CI pipeline and want to mark the step failed if any error found by linter.

See also `-level` flag for [github-pr-check/github-check](#reporter-github-checks--reportergithub-pr-check) reporters.
codebase will exit with `1` if reported check status is `failure` as well if `-fail-on-error=true`.

## Filter mode
codebase filter results by diff and you can control how codebase filter results by `-filter-mode` flag.
Available filter modes are as below.

### `added` (default)
Filter results by added/modified lines.
### `diff_context`
Filter results by diff context. i.e. changed lines +-N lines (N=3 for example).
### `file`
Filter results by added/modified file. i.e. codebase will report results as long as they are in added/modified file even if the results are not in actual diff.
### `nofilter`
Do not filter any results. Useful for posting results as comments as much as possible and check other results in console at the same time.

`-fail-on-error` also works with any filter-mode and can catch all results from any linters with `nofilter` mode.

Example:
```shell
$ codebase -reporter=github-pr-reviewe -filter-mode=nofilter -fail-on-error
```

### Filter Mode Support Table
Note that not all reporters provide full support for filter mode due to API limitation.
e.g. `github-pr-reviewe` reporter uses [GitHub Review
API](https://docs.github.com/en/rest/pulls/reviews) and [GitHub Review
Comment API](https://docs.github.com/en/rest/pulls/comments) but these APIs don't support posting comments outside diff file,
so codebase will use [Check annotation](https://docs.github.com/en/rest/checks/runs) as fallback to post those comments [1]. 

| `-reporter` \ `-filter-mode` | `added` | `diff_context` | `file`                  | `nofilter` |
| ---------------------------- | ------- | -------------- | ----------------------- | ---------- |
| **`local`**                  | OK      | OK             | OK                      | OK |
| **`github-check`**           | OK      | OK             | OK                      | OK |
| **`github-pr-check`**        | OK      | OK             | OK                      | OK |
| **`github-pr-reviewe`**       | OK      | OK             | OK                      | Partially Supported [1] |
| **`gitlab-mr-discussion`**   | OK      | OK             | OK                      | Partially Supported [2] |
| **`gitlab-mr-commit`**       | OK      | Partially Supported [2] | Partially Supported [2] | Partially Supported [2] |
| **`gerrit-change-review`**   | OK      | OK? [3]        | OK? [3]                 | Partially Supported? [2][3] |
| **`bitbucket-code-report`**  | NO [4]  | NO [4]         | NO [4]                  | OK |
| **`gitea-pr-review`**        | OK      | OK             | Partially Supported [2] | Partially Supported [2] |

- [1] Report results that are outside the diff file with Check annotation as fallback if it's running in GitHub actions instead of Review API (comments). All results will be reported to console as well.
- [2] Report results that are outside the diff file to console.
- [3] It should work, but not been verified yet.
- [4] Not implemented at the moment

## Debugging

Use the `-tee` flag to show debug info.

```shell
codebase -filter-mode=nofilter -tee
```

## Articles
- [codebase — A code review dog who keeps your codebase healthy ](https://medium.com/@haya14busa/codebase-a-code-review-dog-who-keeps-your-codebase-healthy-d957c471938b)
- [codebase ♡ GitHub Check — improved automated review experience](https://medium.com/@haya14busa/codebase-github-check-improved-automated-review-experience-58f89e0c95f3)
- [Automated Code Review on GitHub Actions with codebase for any languages/tools](https://medium.com/@haya14busa/automated-code-review-on-github-actions-with-codebase-for-any-languages-tools-20285e04448e)
- [GitHub Actions to guard your workflow](https://evrone.com/blog/github-actions)

## :bird: Author
haya14busa [![GitHub followers](https://img.shields.io/github/followers/haya14busa.svg?style=social&label=Follow)](https://github.com/haya14busa)

## Contributors

[![Contributors](https://opencollective.com/codebase/contributors.svg?width=890)](https://github.com/khulnasoft/codebase/graphs/contributors)

### Supporting codebase

Become GitHub Sponsor for [each contributor](https://github.com/khulnasoft/codebase/graphs/contributors)
or become a backer or sponsor from [opencollective](https://opencollective.com/codebase).

[![Become a backer](https://opencollective.com/codebase/tiers/backer.svg?avatarHeight=64)](https://opencollective.com/codebase#backers)
