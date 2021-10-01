package uuid

import "os/exec"

func GenerateUUID() (string, error) {
	// generate UUID
	rawUUID, err := exec.Command("uuidgen").Output()
	uuid := string(rawUUID[:len(rawUUID)-1]) // strip the trailing newline

	return uuid, err
}
