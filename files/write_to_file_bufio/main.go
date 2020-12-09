package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	// using buffio package: write with a buffer in memory before writing to disk
	// useful when data needs manipulation before writing to disk

	// Opening the file for writing
	file, err := os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Creating a buffered writer from the file variable using bufio.NewWriter()
	bufferedWriter := bufio.NewWriter(file)
	// declaring a byte slice
	bs := []byte{97, 98, 99}
	// writing the byte slice to the buffer in memory
	bytesWritten, err := bufferedWriter.Write(bs)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bytes written to buffer (not file) %d\n", bytesWritten)
	// => 2019/10/21 16:30:59 Bytes written to buffer (not file): 3

	// checking the available buffer
	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes available in buffer: %d\n", bytesAvailable)
	// => 2019/10/21 16:30:59 Bytes available in buffer: 4093

	// writing a string (not a byte slice) to the buffer in memory
	bytesWritten, err = bufferedWriter.WriteString("\nJust a random string")
	if err != nil {
		log.Fatal(err)
	}

	// checking how much data is stored in buffer, just  waiting to be written to disk
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Byte bufered: %d\n", unflushedBufferSize)
	// -> 24 (3 bytes in the byte slice + 21 runes in the string, each rune is 1 byte)

	// The bytes have been written to buffer, not yet to file.
	// Writing from buffer to file.
	bufferedWriter.Flush()
}
