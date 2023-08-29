package main

import (
	"fmt"
	ant "lem_in/tools"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error: There is one error in the file")
		}
	}()
	if len(os.Args) != 2 {

		fmt.Println("Error: Review arguments")
		return
	}
	ant.Lem_in_prog()
}
