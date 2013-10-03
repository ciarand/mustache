package mustache

import (
	"encoding/json"
	//"os"
	//"path"
	"io/ioutil"
	"path/filepath"
	//"strings"
	"testing"
)

// Testing your Mustache implementation against this specification should be
// relatively simple.  If you have a readily available testing framework on your
// platform, your task may be even simpler.

// 1. Use a --YAML-- JSON parser to load the file

// 2. For each test in the tests array

// A. Ensure that each element of the 'partials' hash (if it exists) is stored
// in a place where the interpreter will look for it

// B. We'll skip over the lambdas.yml file for now
// C. Render the template with the given "data"
// 4. Compare the results against "expected"

type MustacheSpecTest struct {
	Data     struct{}
	Desc     string
	Name     string
	Template string
	Expected string
}

type MustacheSpecFile struct {
	Overview string
	Tests    []MustacheSpecTest
}

func runSpecFile(file string, t *testing.T) {
	// Read in the spec file
	rawJson, specErr := ioutil.ReadFile(file)
	if specErr != nil {
		t.Fatal("Could not read file" + file + ": " + specErr.Error())
	}

	// Create a new spec struct
	var spec MustacheSpecFile

	// Unmarshal the json
	err := json.Unmarshal(rawJson, &spec)
	if err != nil {
		t.Fatal("Could not decode json" + err.Error())
	}

	for _, value := range spec.Tests {
		// Run the individual spec test
		runSpecTest(value, t)
	}
}

func runSpecTest(test MustacheSpecTest, t *testing.T) {
	output := Render(test.Template, test.Data)
	if output != test.Expected {
		t.Fail()
		t.Logf("%q (%q) expected %q got %q",
			test.Desc, test.Name, test.Expected, output)
	}
}

func TestSpecs(t *testing.T) {
	// 1. Find each test
	matches, err := filepath.Glob("spec/specs/*.json")
	if err != nil {
		t.Fatal("filepath.Glob: " + err.Error())
	}

	for _, file := range matches {
		// 2. Decode the spec file
		go runSpecFile(file, t)
	}
}
