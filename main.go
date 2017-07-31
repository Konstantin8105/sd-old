package main

import (
	"flag"
	"fmt"
)

// Version can be requested through the command line
const Version = "0.2.0"

func main() {
	version := flag.Bool("version", false, "version of software")
	flag.Parse()
	if *version {
		fmt.Println("Version: ", Version)
	}
}
