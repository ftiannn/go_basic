package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var newFile *os.File //(* = pointer => store memory address of another variable)
	fmt.Printf("%T\n", newFile)

	var err error

	// CREATE A FILE
	newFile, err = os.Create("a.txt")

	// TRUNCATING A FILE
	err = os.Truncate("a.txt", 0)

	// error handling
	if err != nil {
		// method 1
		// fmt.Println(err)
		// os.Exit(1)

		// method 2 (recommended)
		log.Fatal(err)
	}

	newFile.Close()

	// OPEN/CLOSE AN EXISTING FILE

	//method 1
	// file, err := os.Open("a.txt") // => read-only mode

	// method 2: create file if it doesnt exists, or open file in append mode if exists
	// can make use of other OS attributes (os.O_RDONLY, etc)
	// os.O_RDONLY // Read only
	// os.O_WRONLY // Write only
	// os.O_RDWR // Read and write
	// os.O_APPEND // Append to end of file
	// os.O_CREATE // Create is none exist
	// os.O_TRUNC // Truncate file when opening
	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// GETTING FILE INFORMATION
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")
	p := fmt.Println

	p("File Name:", fileInfo.Name())        // => File Name: a.txt
	p("Size in bytes:", fileInfo.Size())    // => Size in bytes: 0
	p("Last modified:", fileInfo.ModTime()) // => Last modified: 2019-10-21 16:16:00.325037748 +0300 EEST
	p("Is Directory? ", fileInfo.IsDir())   // => Is Directory?  false
	p("Pemissions:", fileInfo.Mode())       // => Pemissions: -rw-r-----

	// CHECKING IF FILE EXISTS
	fileInfo, err = os.Stat("b.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File file does not exist!")
		}
	}

	// RENAMING AND MOVING A FILE
	oldPath := "a.txt"
	newPath := "aaa.txt"
	err = os.Rename(oldPath, newPath)
	// error handling
	if err != nil {
		log.Fatal(err)
	}

	// REMOVING A FILE
	err = os.Remove("aaa.txt")
	if err != nil {
		log.Fatal(err)
	}

}
