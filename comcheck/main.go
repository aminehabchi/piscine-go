package main

import "os"

func main() {
	arg := os.Args[1:]
	for i := 0; i < len(arg); i++ {
		if arg[i] == "galaxy" || arg[i] == "galaxy 01" || arg[i] == "01" {
			print("Alert!!!")
		}
	}
}
