//go:build !linux

package executils

import "os/exec"

func LookPath(file string) (string, error) {
	return exec.LookPath(file)
}
