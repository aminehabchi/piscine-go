package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(arg) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	name := "./" + arg[0]
	content, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", string(content))
}
