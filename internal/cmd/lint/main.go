package main

import (
	"github.com/open-panoptes/opni/internal/linter"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(linter.AnalyzerPlugin.GetAnalyzers()...)
}
