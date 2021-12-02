package shell

import (
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	reMetric = regexp.MustCompile(`^\w+({.*})?(\s+\S+)?(\s+\S+)$|^#.*$|^\s*$`)
)

func RunShellCommand(filename string) ([]string, error) {
	cmd := exec.Command("/bin/bash", filename)
	out, err := cmd.Output()
	if err != nil {
		return []string{}, err
	}
	lines := strings.Split(string(out), string('\n'))
	return lines, nil
}

func IsMetric(l string) bool {
	if !reMetric.MatchString(l) {
		return false
	}
	return true
}

func GetScriptsList(path string) ([]string, error) {
	files, err := filepath.Glob(path + "/*.sh")
	if err != nil {
		return []string{}, err
	}
	return files, nil
}
