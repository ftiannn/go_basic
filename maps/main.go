package main

import "fmt"

func main() {
	var emp map[string]string
	fmt.Printf("%#v\n", emp) // => map[string]string(nil)

	fmt.Printf("No of pairs:%d\n", len(emp)) // => No of pairs:0

	fmt.Printf("The value of key %q is %q", "Dan", emp["Dan"]) // => The value of key "Dan" is ""

}
