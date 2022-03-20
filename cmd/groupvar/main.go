package main

import (
	"github.com/rinchsan/groupvar"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(groupvar.Analyzer) }
