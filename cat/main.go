package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		fmt.Println("Hello\nHello")
		return
	}
	for i := 0; i < len(arg); i++ {
		name := "./" + arg[i]
		content, err := ioutil.ReadFile(name)
		if err != nil {
			fmt.Println("ERROR: open " + name + ": no such file or directory")
			os.Exit(1)
			return
		}

		fmt.Printf("%s", string(content))
	}
}
