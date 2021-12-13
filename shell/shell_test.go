package shell

import (
	"shell_exporter/util"
	"testing"
)

func TestIsMetric(t *testing.T) {
	lines := []string{
		`# HELP http_requests_total The total number of HTTP requests.`,
		`# TYPE http_requests_total counter`,
		`http_requests_total{method="post",code="200"} 1027 1395066363000`,
		`http_requests_total{method="post",code="400"}    3 1395066363000`,
		`# Escaping in label values:`,
		`msdos_file_access_time_seconds{path="C:/DIR/FILE.TXT",error="Cannot find file:\n\"FILE.TXT\""} 1.458255915e9`,
		`# Minimalistic line:`,
		`metric_without_timestamp_and_labels 12.47`,
		`# A weird metric from before the epoch:`,
		`something_weird{problem="division by zero"} +Inf -3982045`,
		`# A histogram, which has a pretty complex representation in the text format:`,
		`# HELP http_request_duration_seconds A histogram of the request duration.`,
		`# TYPE http_request_duration_seconds histogram`,
		`http_request_duration_seconds_bucket{le="0.05"} 24054`,
		`http_request_duration_seconds_bucket{le="0.1"} 33444`,
		`http_request_duration_seconds_bucket{le="0.2"} 100392`,
		`http_request_duration_seconds_bucket{le="0.5"} 129389`,
		`http_request_duration_seconds_bucket{le="1"} 133988`,
		`http_request_duration_seconds_bucket{le="+Inf"} 144320`,
		`http_request_duration_seconds_sum 53423`,
		`http_request_duration_seconds_count 144320`,
		`# Finally a summary, which has a complex representation, too:`,
		`# HELP rpc_duration_seconds A summary of the RPC duration in seconds.`,
		`# TYPE rpc_duration_seconds summary`,
		`rpc_duration_seconds{quantile="0.01"} 3102`,
		`rpc_duration_seconds{quantile="0.05"} 3272`,
		`rpc_duration_seconds{quantile="0.5"} 4773`,
		`rpc_duration_seconds{quantile="0.9"} 9001`,
		`rpc_duration_seconds{quantile="0.99"} 76656`,
		`rpc_duration_seconds_sum 1.7560473e+07`,
		`rpc_duration_seconds_count 2693`,
	}

	for _, l := range lines {
		if !IsMetric(l) {
			t.Errorf("IsMetric() = false, want true, line = '%s'", l)
		}
	}

	wrongLines := []string{
		`wrong metric 1234`,
	}

	for _, l := range wrongLines {
		if IsMetric(l) {
			t.Errorf("IsMetric() = true, want false, line = '%#v'", l)
		}
	}
}

func TestGetScriptsList(t *testing.T) {
	list, err := GetScriptsList("../test_files")
	if err != nil {
		t.Errorf("function fail with error '%#v'", err)
	}

	want := []string{
		"../test_files/err_test.sh",
		"../test_files/test1.sh",
		"../test_files/test2.sh",
	}
	if !util.Equal(list, want) {
		t.Errorf("\n result '%#v',\n want   '%#v'", list, want)
	}
}

func TestRunShellCommand(t *testing.T) {
	result, err := RunShellCommand("../test_files/test1.sh")
	if err != nil {
		t.Errorf("function fail with error '%#v'", err)
	}

	want := []string{
		`# HELP test1_metric1 First metric`,
		`# TYPE test1_metric1 counter`,
		`test1_metric1{tag1="t1",tag2="t2"} 1234 1395066363000`,
		`# HELP test1_metric2 First metric`,
		`# TYPE test1_metric2 counter`,
		`test1_metric2{tag1="t1",tag2="t2"} 5678 1395066363001`,
	}
	if !util.Equal(result, want) {
		t.Errorf("\n result '%#v',\n want   '%#v'", result, want)
	}

	result, err = RunShellCommand("../test_files/nonexistent.sh")
	if err == nil {
		t.Errorf("function must fail")
	}

	result, err = RunShellCommand("../test_files/err_test.sh")
	if err == nil {
		t.Errorf("function must fail")
	}
}
