package main

import (
	"fmt"
	"io"
)

var (
	version   string
	commit    string
	buildTime string
	buildUser string
)

func printVersion(w io.Writer) {
	fmt.Fprintf(w, "version: %s", version)
	fmt.Fprintf(w, "git commit: %s", commit)
	fmt.Fprintf(w, "build time: %s", buildTime)
	fmt.Fprintf(w, "build user: %s", buildUser)
}
