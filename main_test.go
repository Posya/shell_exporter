package main

import (
	"shell_exporter/util"
	"testing"
)

func TestGetMetrics(t *testing.T) {
	want := []string{
		`# HELP test1_metric1 First metric`,
		`# TYPE test1_metric1 counter`,
		`test1_metric1{tag1="t1",tag2="t2"} 1234 1395066363000`,
		`# HELP test1_metric2 First metric`,
		`# TYPE test1_metric2 counter`,
		`test1_metric2{tag1="t1",tag2="t2"} 5678 1395066363001`,
		`# HELP test2_metric1 First metric`,
		`# TYPE test2_metric1 counter`,
		`test2_metric1{tag1="t1",tag2="t2"} 1234 1395066363000`,
		`# HELP test2_metric2 First metric`,
		`# TYPE test2_metric2 counter`,
		`test2_metric2{tag1="t1",tag2="t2"} 5678 1395066363001`,
	}

	result := GetMetrics("./test_files/")
	if !util.Equal(result, want) {
		t.Errorf("\n result '%#v',\n want   '%#v'", result, want)
	}
}
