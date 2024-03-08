package goscreen

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/kevinpita/goscreen/internal/execute"
)

const (
	command        = "screen"
	versionPattern = `^Screen version .* \(GNU\)`
)

var (
	ErrCommandNotFound = errors.New("GNU Screen execute not found")
	ErrCommandVersion  = fmt.Errorf("unexpected GNU Screen version, expecting pattern: %s", versionPattern)

	versionRegex = regexp.MustCompile(versionPattern)
)

type Screen struct {
	Executor execute.Execute
}

// Ok checks if GNU Screen is available, returns true on success, false otherwise.
//
// ErrCommandNotFound is returned as error if `screen` command is not present.
//
// ErrCommandVersion is returned as error if `screen -v` returns an unexpected output.
//
// An unspecified error is returned if `screen -v` returns an error.
func Ok() (bool, error) {
	const versionArg = "-v"

	c := execute.New()

	_, errPath := c.LookPath(command)
	if errPath != nil {
		return false, ErrCommandNotFound
	}

	cmdOut, errCmd := c.Run(command, versionArg)
	if errCmd != nil {
		return false, fmt.Errorf("'%s %s': %w", command, versionArg, errCmd)
	}

	if !versionRegex.MatchString(cmdOut) {
		return false, ErrCommandVersion
	}

	return true, nil
}
