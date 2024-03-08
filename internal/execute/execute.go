package execute

import (
	"os/exec"
	"strings"
)

type Execute interface {
	LookPath(file string) (string, error)
	Run(name string, arg ...string) (string, error)
}

// New returns an empty struct implementing methods using exec.LookPath and exec.Command with cmd.Run.
func New() Execute {
	return &cmd{}
}

type cmd struct {
}

// LookPath calls exec.LookPath.
func (*cmd) LookPath(file string) (string, error) {
	return exec.LookPath(file)
}

// Run creates an exec.Command and executes the execute, returning its output as a string and its error.
func (*cmd) Run(name string, arg ...string) (string, error) {
	var out strings.Builder

	cmd := exec.Command(name, arg...)
	cmd.Stdout = &out

	err := cmd.Run()
	return out.String(), err
}
