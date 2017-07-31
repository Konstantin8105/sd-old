package main

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	var args []string
	args = append(args, "-version")
	os.Args = args
	flag.Parse()
	result := m.Run()
	os.Exit(result)
}
