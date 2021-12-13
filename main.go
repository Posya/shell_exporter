package main

import (
	"fmt"
	log "github.com/go-pkgz/lgr"
	"github.com/jessevdk/go-flags"
	"net/http"
	"os"
	"shell_exporter/shell"
	"strconv"
	"strings"
)

type Opts struct {
	Debug       bool   `long:"debug" env:"DEBUG" description:"debug mode"`
	ScriptsPath string `short:"f" long:"scripts" env:"SCRIPTS" default:"/etc/shell_exporter/" description:"path to scripts"`
	Port        int    `short:"p" long:"port" env:"PORT" default:"9212" description:"path to scripts"`
}

var version = "unknown"
var opts Opts

func main() {
	fmt.Printf("shell_exporter %s\n", version)
	_, err := flags.Parse(&opts)
	setupLog(opts.Debug)

	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/metrics", MetricsHandler)

	domain := ":" + strconv.Itoa(opts.Port)
	log.Printf("[INFO] Starting server on %s", domain)

	err = http.ListenAndServe(domain, nil)
	if err != nil {
		log.Printf("[ERROR] failed with %#v", err)
	}
}

func RootHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, `<html>
<head><title>Shell Exporter</title></head>
<body>
<h1>Shell Exporter</h1>
<p><a href="/metrics">Metrics</a>
</body></html>
`)
	if err != nil {
		log.Printf("[ERROR] fail to send main page: %#v", err)
	}
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := fmt.Fprintf(w, strings.Join(GetMetrics(opts.ScriptsPath), "\n"))
	if err != nil {
		log.Printf("[ERROR] fail to send main page: %#v", err)
	}
}

func GetMetrics(ScriptsPath string) []string {
	scripts, err := shell.GetScriptsList(ScriptsPath)
	if err != nil {
		log.Printf("[ERROR] fail in getting scripts list: %#v", err)
		return []string{}
	}

	var result []string

	for _, s := range scripts {
		out, err := shell.RunShellCommand(s)
		if err != nil {
			log.Printf("[ERROR] fail in running script %s: %#v", s, err)
		}

		result = append(result, out...)
	}
	return result
}

func setupLog(dbg bool) {
	if dbg {
		log.Setup(log.Debug, log.CallerFile, log.CallerFunc, log.Msec, log.LevelBraces)
		return
	}
	log.Setup(log.Msec, log.LevelBraces)
}
