package shell

import (
	"os/exec"
	"path/filepath"
	"regexp"
	"shell_exporter/util"
	"strings"
)

var (
	reMetric = regexp.MustCompile(`^[\w_-]+ *({[^}]*})? +([\d.-]+ +|[+-]Inf +)?[\d.-]+(e[+-]?\d+)?$|^#.*$`)
)

func RunShellCommand(filename string) ([]string, error) {
	checkMetricsAndTrim := func(l string) string {
		if IsMetric(l) {
			return strings.TrimSpace(l)
		}
		return ""
	}

	cmd := exec.Command("/bin/bash", filename)
	out, err := cmd.Output()
	if err != nil {
		return []string{}, err
	}
	lines := strings.Split(string(out), string('\n'))
	return util.FilterSlice(lines, checkMetricsAndTrim), nil
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
