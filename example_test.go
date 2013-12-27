package canopy

import (
	"fmt"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type S struct{}

var _ = Suite(&S{})

func ExampleNewGovernedProcess() {
	gp1 := NewGovernedProcess("show:ls", "-lA", "-h")
	gp2 := NewGovernedProcess("which", "ls")
	fmt.Println(gp1)
	fmt.Println(gp2)
}
