package main

import (
	"fmt"
	f "fmt" // permitted in GO
)

// package scoped
const done = false

func main() {
	// local (block) scoped
	var task = "Running"
	f.Println(task, done)

	const done = true                      // local scoped
	f.Printf("done in main(): %v\n", done) // => done in main(): true
	f1()

}

func f1() {
	// This is done from package scope
	fmt.Printf("done in f1(): %v\n", done)
}
