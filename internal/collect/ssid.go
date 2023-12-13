package collect

import (
	"os/exec"
	"strings"
)

func GetSssid() (string, error) {
	cmd := exec.Command("iwgetid")

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Trim string
	split := strings.Split(string(out), `:`)
	replace := strings.Replace(split[1], `"`, "", -1)
	trim := strings.TrimSuffix(replace, "\n")

	return trim, nil
}
