package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// opening the file in write-only mode if the file exists and then it truncates the file.
	// if the file doesn't exist it creates the file with 0644 permissions
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	// defer closing the file
	defer file.Close()

	// WRITING A NEW FILE USING `os` package
	byteSlice := []byte("Hello World!")       // converting a string to a bytes slice
	byteWritten, err := file.Write(byteSlice) // writing bytes to file.
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes Written: %d\n", byteWritten)

	// WRITING A NEW FILE USING `ioutil` package
	// ioutil.WriteFile() handles creating, opening, writing a slice of bytes and closing the file.
	// if the file doesn't exist WriteFile() creates it
	// and if it already exists the function will truncate it before writing to file.

	bs := []byte("Hello World again!")
	err = ioutil.WriteFile("c.txt", bs, 0644)
}
