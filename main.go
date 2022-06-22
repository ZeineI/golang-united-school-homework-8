package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type Arguments map[string]string

type user struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

var (
	idFlag        = flag.String("id", "", "identificator")
	operationFlag = flag.String("operation", "", "action in json")
	fileNameFlag  = flag.String("fileName", "", "output path")
	itemFlag      = flag.String("item", "", "data info")
)

var (
	errorfileName  = errors.New("-fileName flag has to be specified")
	errorOperation = errors.New("-operation flag has to be specified")
	errorItem      = errors.New("-item flag has to be specified")
	errorID        = errors.New("-id flag has to be specified")
)

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
		"fileName":  *fileNameFlag,
		"item":      *itemFlag,
	}
	fmt.Println(mpUser)
	return mpUser
}

func Perform(args Arguments, writer io.Writer) error {
	if args["fileName"] == "" {
		fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAA")
		fmt.Println(args)
		return errorfileName
	}

	opName := args["operation"]
	switch opName {
	case "add":
		return addF(args, writer)
	case "list":

	case "findById":

	case "remove":

	case "":
		return errorOperation
	default:
		return fmt.Errorf("Operation %s not allowed!", opName)
	}

	return nil
}

func readUsers(file string) ([]user, error) {
	var users []user
	//read file
	dat, err := os.ReadFile(file)
	if err != nil {
		return users, err
	}

	if len(dat) == 0 {
		return users, nil
	}
	err = json.Unmarshal(dat, &users)
	if err != nil {
		return users, err
	}
	return users, nil
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func addF(args Arguments, writer io.Writer) error {
	itemS := args["item"]
	if itemS == "" {
		return errorItem
	}

	exist, err := Exists(args["fileName"])
	if err != nil {
		return err
	}

	if !exist {
		_, err := os.Create(args["fileName"])
		if err != nil {
			return err
		}
	}
	users, err := readUsers(args["fileName"])
	if err != nil {
		return err
	}

	//convert string to user
	item, err := readItem(itemS)
	if err != nil {
		return err
	}

	if alreadyExist(users, item) {
		fmt.Println("AAAAAAAAAAAAAAAA")
		_, err := writer.Write([]byte(fmt.Sprintf("Item with id %s already exists", item.Id)))
		return err
	}
	users = append(users, item)

	newData, err := updateData(users)
	if err != nil {
		return err
	}

	if err = updateFile(args["fileName"], newData); err != nil {
		return err
	}
	return nil
}

func alreadyExist(users []user, item user) bool {
	for _, v := range users {
		if v.Id == item.Id {
			return true
		}
	}
	return false
}

func updateFile(fileN string, data []byte) error {
	file, err := os.OpenFile(fileN, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err = file.Write(data); err != nil {
		return err
	}
	return nil
}

func updateData(users []user) ([]byte, error) {
	res, err := json.Marshal(users)
	if err != nil {
		return res, err
	}
	return res, nil
}

func readItem(item string) (user, error) {
	var itm user
	err := json.Unmarshal([]byte(item), &itm)
	if err != nil {
		return itm, err
	}
	return itm, nil
}
