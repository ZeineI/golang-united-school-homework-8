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
	idFlag        = flag.String("id", "", "identificator")
	operationFlag = flag.String("operation", "", "action in json")
	filenameFlag  = flag.String("filename", "", "output path")
	itemFlag      = flag.String("item", "", "data info")
)

var (
	errorFilename  = errors.New("-fileName flag has to be specified")
	errorOperation = errors.New("-operation flag has to be specified")
	errorItem      = errors.New("-item flag has to be specified")
	errorID        = errors.New("-id flag has to be specified")
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
		"id":        *idFlag,
		"operation": *operationFlag,
		"filename":  *filenameFlag,
		"item":      *itemFlag,
	}
	return mpUser
}
