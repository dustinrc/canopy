package canopy

import (
	"fmt"
	. "launchpad.net/gocheck"
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

func (s *S) TestNewGovernedProcess(c *C) {
	for _, tt := range newGovernedProcessTests {
		gp := NewGovernedProcess(tt.args[0], tt.args[1:]...)
		c.Assert(gp.alias, Equals, tt.expectedAlias)
	}
}

func (s *S) TestGovernedProcessString(c *C) {
	gp := NewGovernedProcess("my-process:../path/to/executable", "arg1", "arg2", "arg3")
	expected := "GP[my-process]: ../path/to/executable arg1 arg2 arg3"
	actual := fmt.Sprintf("%s", gp)
	c.Assert(expected, Equals, actual)
}
