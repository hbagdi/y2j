package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	"github.com/hokaccha/go-prettyjson"
)

func die(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(1)
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		die("failed to read input: %v", err)
	}
	output, err := yaml.YAMLToJSON(input)
	if err != nil {
		die("YAML to JSON conversion failed: %v", err)
	}
	colorOutput, err := prettyjson.Format(output)
	if err != nil {
		die("failed to colorize output: %v", err)
	}
	fmt.Printf("%s\n", colorOutput)
}
