package main

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.T) {
	var args []string
	args = append(args, "-version")
	os.Args = args
	flag.Parse()
	main()
}
