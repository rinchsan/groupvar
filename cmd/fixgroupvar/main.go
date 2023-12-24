package main

import (
	"os"

	"github.com/rinchsan/groupvar"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	os.Args = append([]string{os.Args[0], "-fix"}, os.Args[1:]...)
	singlechecker.Main(groupvar.Analyzer)
}
