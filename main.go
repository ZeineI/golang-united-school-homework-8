package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type Arguments map[string]string

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

var (
	operationFlag = flag.String("operation", "", "describe the action")
	filenameFlag  = flag.String("filename", "", "where data save")
	itemFlag      = flag.String("item", "", "output")
)

var (
	errorFilename  = errors.New("-fileName flag has to be specified")
	errorOperation = errors.New("-operation flag has to be specified")
	errorNotExist  = errors.New("Operation abcd not allowed!")
)

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
		"operation": *operationFlag,
		"filename":  *filenameFlag,
		"item":      *itemFlag,
	}
	return mpUser
}
