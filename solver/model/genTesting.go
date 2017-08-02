// +build ignore

// It can be invoked by running go generate

package main

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

type template struct {
	begin string
	test  testFunc
}

type testFunc struct {
	header        headerBlock
	blocks        []block
	footerSuccess string
	footerFail    string
}

type headerBlock struct {
	testName string
	body     string
}

type block struct {
	name string
	body []string
}

func main() {

}
