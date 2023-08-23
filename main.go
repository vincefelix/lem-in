package main

import (
	"fmt"
	ant "lem_in/tools"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Review arguments")
		return
	}
	ant.Lem_in_prog()
}
