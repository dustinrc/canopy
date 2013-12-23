package canopy

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// GovernedProcess represents a process to be overseen.
type GovernedProcess struct {
	alias string
	cmd   *exec.Cmd
}

// NewGovernedProcess creates a new GovernedProcess instance.
// path is the command to be executed along with zero or more args.
// A custom alias can be provided by prefixing the path, e.g.,
// "list:/usr/sbin/ls" would set the alias as "list".
func NewGovernedProcess(path string, arg ...string) *GovernedProcess {
	split := strings.SplitN(path, ":", 2)
	if len(split) == 1 {
		split = []string{filepath.Base(path), path}
	}
	alias := split[0]
	path = split[1]
	return &GovernedProcess{alias, exec.Command(path, arg...)}
}

// String displays a readable format of GovernedProcess.
func (gp *GovernedProcess) String() string {
	args := strings.Join(gp.cmd.Args[1:], " ")
	return fmt.Sprintf("GP[%s]: %s %s", gp.alias, gp.cmd.Path, args)
}
