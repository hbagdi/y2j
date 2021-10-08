package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	"github.com/hokaccha/go-prettyjson"
)

var help = flag.Bool("help",false,"print help for y2j")

func die(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(1)
}

func main() {
	flag.Parse()
	if *help {
		fmt.Println(`y2j reads a YAML from the standard input (stdin)
and outputs corresponding JSON on standard output (stdout).`)
		flag.PrintDefaults()
		return
	}
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
