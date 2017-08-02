// +build ignore

// It can be invoked by running go generate

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

/*
* find tempate of test [*.gotmpl]
* find all variables in test
* create combinations
* generate tests
	* present sequence of blocks - for success result
	* randomize sequence of blocks - for success result
	* remove one or more blocks - for fail result
* write into file [*_test.go]
*/

type lines []string

type template struct {
	begin lines
	tests []test
}

type test struct {
	header        headerBlock
	blocks        []block
	footerSuccess lines
	footerFail    lines
}

type headerBlock struct {
	testName string
	body     lines
}

type block struct {
	name string
	body []lines
}

func main() {
	files, _ := filepath.Glob("*.gotmpl")
	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Errorf("Cannot open file : %v", fileName)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}

		if err := scanner.Err(); err != nil {
			fmt.Errorf("err = %v", err)
			return
		}
	}
}
