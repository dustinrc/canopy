package canopy_test

import (
	"fmt"
	"github.com/dustinrc/canopy"
)

func ExampleNewGovernedProcess() {
	gp1 := canopy.NewGovernedProcess("show:ls", "-lA", "-h")
	gp2 := canopy.NewGovernedProcess("which", "ls")
	fmt.Println(gp1)
	fmt.Println(gp2)
}
