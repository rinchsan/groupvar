# groupvar

[![pkg.go.dev][gopkg-badge]][gopkg]

`groupvar` finds low-readability variable/constant declarations.

## Installation

```shell
go install github.com/rinchsan/groupvar/cmd/groupvar@latest
```

## Usage

```shell
go vet -vettool=`which groupvar` testdata/src/a/a.go
```
