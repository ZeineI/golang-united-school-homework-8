package main

import (
	"flag"
	"io"
	"os"
)

type Arguments map[string]string

var operationFlaf = flag.String("operation", "", "describe the action")

func Perform(args Arguments, writer io.Writer) error {
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}

func parseArgs() Arguments {
	mpUser := Arguments{
		"operation": *operationFlaf,
	}
	return mpUser
}
