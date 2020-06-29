package main

import (
	"github.com/tommyknows/funcheck/prettyprint"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(prettyprint.Analyzer)
}
