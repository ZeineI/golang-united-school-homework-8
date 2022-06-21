package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type Arguments map[string]string

var operationFlaf = flag.String("operation", "", "describe the action")

func Perform(args Arguments, writer io.Writer) error {
	fmt.Println(args)
	return nil
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}

func parseArgs() Arguments {
	flag.Parse()
	mpUser := Arguments{
		"operation": *operationFlaf,
	}
	return mpUser
}
