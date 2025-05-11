package main

import (
	"testing"

	"github.com/karetskiiVO/DoCompiler/multilinector"
	"github.com/stretchr/testify/assert"
)

func TestCompile(t *testing.T) {
	type TestConfig struct {
		name             string
		files            []string
		expectedErrorMsg string
	}

	testConfigs := []TestConfig{
		{
			name:  "scopes",
			files: []string{"./testdata/scopes/main.do"},
			expectedErrorMsg: multilinector.New().
				AddLine("./testdata/scopes/main.do:13:4: variable `c` is not declared in this scope").
				AddLine("./testdata/scopes/main.do:13:4: incorrect number of expressions in the assignment expected: 1 actual: 0").
				String(),
		},
		{
			name:  "assignments",
			files: []string{"./testdata/assignments/main.do"},
			expectedErrorMsg: multilinector.New().
				AddLine("./testdata/assignments/main.do:6:7: incorrect number of expressions in the assignment expected: 2 actual: 1").
				AddLine("./testdata/assignments/main.do:7:4: incorrect number of expressions in the assignment expected: 1 actual: 2").
				AddLine("./testdata/assignments/main.do:12:7: incorrect number of expressions in the assignment expected: 2 actual: 1").
				AddLine("./testdata/assignments/main.do:15:4: incorrect number of expressions in the assignment expected: 1 actual: 2").
				AddLine("./testdata/assignments/main.do:18:4: mistmatch in argument values: expected: 3 actual: 0").
				AddLine("./testdata/assignments/main.do:20:4: mistmatch in argument values: expected: 3 actual: 4").
				String(),
		},
	}

	for _, testCfg := range testConfigs {
		t.Run(testCfg.name, func(t *testing.T) {
			_, err := Compile(testCfg.files...)

			if testCfg.expectedErrorMsg == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, testCfg.expectedErrorMsg)
			}
		})
	}

	Compile("./doexamples/generaltest_simple.do")
}
