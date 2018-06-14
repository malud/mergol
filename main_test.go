package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMain__Simple(t *testing.T) {
	// arrange
	in = "./testdata/simple"
	out = "./testdata/simple/result/merged.go"
	// action
	main()
	// verify
	result, _ := ioutil.ReadFile(out)
	expected, _ := ioutil.ReadFile("./testdata/simple/result/merged_expected.go")
	if !strings.EqualFold(string(expected), string(result)) {
		t.Error("Merged result is not in expected form")
	}
	// cleanup
	os.Remove(out)
}

func TestMain__Complex(t *testing.T) {
	// arrange
	in = "./testdata/complex"
	out = "./testdata/complex/result/merged.go"
	// action
	main()
	// verify
	result, _ := ioutil.ReadFile(out)
	expected, _ := ioutil.ReadFile("./testdata/complex/result/merged_expected.go")
	if !strings.EqualFold(string(expected), string(result)) {
		t.Error("Merged result is not in expected form")
	}
	// cleanup
	os.Remove(out)
}
