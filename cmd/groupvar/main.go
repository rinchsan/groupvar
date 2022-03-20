package main

import (
	"groupvar"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(groupvar.Analyzer) }
