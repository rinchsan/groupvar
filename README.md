# groupvar

[![](https://pkg.go.dev/badge/github.com/rinchsan/groupvar.svg)](https://pkg.go.dev/github.com/rinchsan/groupvar/cmd/groupvar)

`groupvar` finds low-readability variable/constant declarations.

## Installation

```shell
go install github.com/rinchsan/groupvar/cmd/groupvar@latest
```

## Usage

```shell
go vet -vettool=`which groupvar` testdata/src/a/a.go
```
