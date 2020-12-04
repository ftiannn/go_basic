package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("os.Args", os.Args)
	fmt.Println("Path: ", os.Args[0])
	fmt.Println("1st Arg: ", os.Args[1])
	fmt.Println("2nd Arg: ", os.Args[2])
	fmt.Println("No of item inside os.Args: ", len(os.Args))

	// try run "go run main.go python go java php"
}
