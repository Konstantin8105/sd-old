// +build ignore

// It can be invoked by running go generate

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
	t     test
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

func (h headerBlock) getBody(name string) lines {
	var d []string
	for _, l := range h.body {
		d = append(d, l)
	}
	testName := "{{TestName}}"
	for i := range d {
		d[i] = strings.Replace(d[i], testName, name, -1)
	}
	return lines(d)
}

type block struct {
	name string
	body lines
}

type testStage int

const (
	noneTestStage testStage = iota
	headerStage
	blockStage
	footerSuccessStage
	footerFailStage
)

func (tmpl *template) parseTest(txt lines) {
	stage := noneTestStage
	var blockName string
	var headerLines lines
	for _, line := range []string(txt) {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		short := strings.TrimSpace(line)
		switch short {
		case "<<Header>>":
			if stage == headerStage {
				stage = noneTestStage
			} else {
				stage = headerStage
			}
			continue
		case "<<FooterSuccess>>":
			if stage == footerSuccessStage {
				stage = noneTestStage
			} else {
				stage = footerSuccessStage
			}
			continue
		case "<<FooterFail>>":
			if stage == footerFailStage {
				stage = noneTestStage
			} else {
				stage = footerFailStage
			}
			continue
		}
		blockPrefix := "<<Block"
		if strings.HasPrefix(short, blockPrefix) {
			blockName = short[len(blockPrefix) : len(short)-2]
			if stage == blockStage {
				stage = noneTestStage
			} else {
				stage = blockStage
				tmpl.t.blocks = append(tmpl.t.blocks, block{name: blockName})
			}
			continue
		}
		switch stage {
		case headerStage:
			headerLines = append(headerLines, line)
		case blockStage:
			size := len([]block(tmpl.t.blocks))
			tmpl.t.blocks[size-1].body = append(tmpl.t.blocks[size-1].body, line)
		case footerSuccessStage:
			tmpl.t.footerSuccess = append(tmpl.t.footerSuccess, line)
		case footerFailStage:
			tmpl.t.footerFail = append(tmpl.t.footerFail, line)
		default:
			panic(fmt.Errorf("Strange line = ", line))
		}
	}
	tmpl.t.header.body = headerLines
}

type stageParse int

const (
	noneTemplateStage stageParse = iota
	beginTemplateStage
	testTemplateStage
)

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
		stage := noneTemplateStage
		var tmpl template
		var testLines lines
		for scanner.Scan() {
			line := scanner.Text()
			if len(strings.TrimSpace(line)) == 0 {
				continue
			}
			short := strings.TrimSpace(line)
			switch short {
			case "<<Begin>>":
				if stage == beginTemplateStage {
					stage = noneTemplateStage
				} else {
					stage = beginTemplateStage
				}
				continue
			case "<<Test>>":
				if stage == testTemplateStage {
					stage = noneTemplateStage
				} else {
					stage = testTemplateStage
				}
				continue
			}
			switch stage {
			case beginTemplateStage:
				tmpl.begin = append(tmpl.begin, line)
			case testTemplateStage:
				testLines = append(testLines, line)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Errorf("err = %v", err)
			return
		}

		tmpl.parseTest(testLines)

		fileName := fileName[0:len(fileName)-len(filepath.Ext(fileName))] + "_test.go"

		// remove file
		_ = os.Remove(fileName)

		_, err = os.Stat(fileName)

		// create file if not exists
		if os.IsNotExist(err) {
			fileT, _ := os.Create(fileName)
			fileT.Close()
		}

		fileOut, err := os.OpenFile(fileName, os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Cannot open file ")
			return
		}
		defer fileOut.Close()

		// write data
		generateTest(tmpl, fileOut)
		// end write data

		// save changes
		err = fileOut.Sync()
		if err != nil {
			return
		}
	}
}

func generateTest(tmpl template, file *os.File) {
	for _, b := range []string(tmpl.begin) {
		fmt.Fprintln(file, b)
	}

	generateSuccess(tmpl, file)
	generateSuccessInvert(tmpl, file)
	generateWithoutOneForFail(tmpl, file)
	generateWithoutOneExpectForFail(tmpl, file)
	generateDublicateForFail(tmpl, file)
}

func generateSuccess(tmpl template, file *os.File) {
	h := tmpl.t.header.getBody("SuccessFull")
	for _, t := range []string(h) {
		fmt.Fprintln(file, t)
	}
	for _, t := range tmpl.t.blocks {
		for _, tt := range []string(t.body) {
			fmt.Fprintln(file, tt)
		}
	}
	for _, t := range []string(tmpl.t.footerSuccess) {
		fmt.Fprintln(file, t)
	}
}

func generateSuccessInvert(tmpl template, file *os.File) {
	h := tmpl.t.header.getBody("SuccessInvertFull")
	for _, t := range []string(h) {
		fmt.Fprintln(file, t)
	}
	for i := range tmpl.t.blocks {
		t := tmpl.t.blocks[len(tmpl.t.blocks)-1-i]
		for _, tt := range []string(t.body) {
			fmt.Fprintln(file, tt)
		}
	}
	for _, t := range []string(tmpl.t.footerSuccess) {
		fmt.Fprintln(file, t)
	}
}

func generateWithoutOneForFail(tmpl template, file *os.File) {
	for i := range tmpl.t.blocks {
		h := tmpl.t.header.getBody(fmt.Sprintf("WithoutOneForFail%v", i))
		for _, t := range []string(h) {
			fmt.Fprintln(file, t)
		}
		for j, t := range tmpl.t.blocks {
			if i == j {
				continue
			}
			for _, tt := range []string(t.body) {
				fmt.Fprintln(file, tt)
			}
		}
		for _, t := range []string(tmpl.t.footerFail) {
			fmt.Fprintln(file, t)
		}
	}
}

func generateWithoutOneExpectForFail(tmpl template, file *os.File) {
	for i := range tmpl.t.blocks {
		h := tmpl.t.header.getBody(fmt.Sprintf("WithoutOneExpectForFail%v", i))
		for _, t := range []string(h) {
			fmt.Fprintln(file, t)
		}
		for j, t := range tmpl.t.blocks {
			if i != j {
				continue
			}
			for _, tt := range []string(t.body) {
				fmt.Fprintln(file, tt)
			}
		}
		for _, t := range []string(tmpl.t.footerFail) {
			fmt.Fprintln(file, t)
		}
	}
}

func generateDublicateForFail(tmpl template, file *os.File) {
	for i := range tmpl.t.blocks {
		h := tmpl.t.header.getBody(fmt.Sprintf("DublicateForFail%v", i))
		for _, t := range []string(h) {
			fmt.Fprintln(file, t)
		}
		for j, t := range tmpl.t.blocks {
			if i == j {
				for _, tt := range []string(t.body) {
					fmt.Fprintln(file, tt)
				}
			}
			for _, tt := range []string(t.body) {
				fmt.Fprintln(file, tt)
			}
		}
		for _, t := range []string(tmpl.t.footerFail) {
			fmt.Fprintln(file, t)
		}
	}
}
