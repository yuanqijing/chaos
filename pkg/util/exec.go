package util

import "os/exec"

func ExecCmd(cmd string) ([]byte, error) {
	// exec cmd and return output
	rawOutput, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return nil, err
	}
	return rawOutput, nil
}
