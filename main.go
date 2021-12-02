package main

import (
	"fmt"
	log "github.com/go-pkgz/lgr"
	"github.com/jessevdk/go-flags"
	"net/http"
	"os"
)

type Opts struct {
	Debug bool `long:"dbg" env:"DEBUG" description:"debug mode"`
}

var version = "unknown"

func main() {
	fmt.Printf("shell_exporter %s\n", version)
	var opts Opts
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
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("[ERROR] failed with %+v", err)
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
		panic(err)
	}
}

func setupLog(dbg bool) {
	if dbg {
		log.Setup(log.Debug, log.CallerFile, log.CallerFunc, log.Msec, log.LevelBraces)
		return
	}
	log.Setup(log.Msec, log.LevelBraces)
}
