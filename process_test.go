package canopy

import (
	"fmt"
	"testing"
)

// TODO(dustinrc): tests for Windows
var newGovernedProcessTests = []struct {
	args          []string
	expectedAlias string
}{
	{[]string{"my-process:/path/to/executable", "arg1", "arg2", "arg3"}, "my-process"},
	{[]string{"../path/to/executable", "arg1", "arg2"}, "executable"},
	{[]string{"my-process:path/to/executable"}, "my-process"},
	{[]string{":../path/to/executable", "arg3"}, ""},
}

func TestNewGovernedProcess(t *testing.T) {
	for _, tt := range newGovernedProcessTests {
		gp := NewGovernedProcess(tt.args[0], tt.args[1:]...)
		if gp.alias != tt.expectedAlias {
			t.Errorf("NewGovernedProcess alias: expected %s, actual %s", tt.expectedAlias, gp.alias)
		}
	}
}

func TestGovernedProcessString(t *testing.T) {
	gp := NewGovernedProcess("my-process:../path/to/executable", "arg1", "arg2", "arg3")
	expected := "GP[my-process]: ../path/to/executable arg1 arg2 arg3"
	if actual := fmt.Sprintf("%s", gp); actual != expected {
		t.Errorf("(*GovernedProcess) String: expected %s, actual %s", expected, actual)
	}
}
