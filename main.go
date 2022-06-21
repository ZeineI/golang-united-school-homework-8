package main

import (
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
