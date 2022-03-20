package groupvar_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/rinchsan/groupvar"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, groupvar.Analyzer, "a")
}
