package main

import "fmt"

// Unused global variable will not lead to compile with error
var y int = 20

// y := 20 (Cannot use short declaration for package scoped variables)

func main() {
	var age int = 30
	fmt.Println("Age: ", age)

	// Mute compile error returned by unused local variables
	// using a blank identifier "_ = xxx"
	var name string = "Dan"
	_ = name

	// Short declaration operator, works only in block scope
	sentence := "Learning golang!"
	fmt.Println(sentence)

	// To overwrite variable; use = instead of :=
	name = "Danny"

	// Multiple declaration can be used only if
	// one of the variables is not being declared previously
	car, cost := "Audi", 5000
	fmt.Println(car, cost)

	// This will prompt error
	// car, cost := "BMW", 6000

	car, year := "BMW", 2019
	fmt.Println(car, year)

	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	// Declare and initialise multiple variables of DIFFERENT data type
	// Go Zero Value (inital values): numberic = 0; bool = false; string = ""; pointer = nil
	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	// Declare and initialise multiple variables of SAME data type
	var h1, h2, h3 int
	fmt.Println(h1, h2, h3)

	var i, j int
	i, j = 5, 8

	// Swapping variable over
	i, j = j, i
	fmt.Println(i, j)

	// short declaration with summation
	sum := 5 + 2.3
	fmt.Println(sum)

	var v1 = 4   // Int
	var v2 = 5.2 //Float

	// v1 != v2 ==> will throw error if try to set v1 = v2
	v1 = int(v2)
	fmt.Println(v1, v2)

	// VERBS:
	// %d -> decimal
	// %f -> float
	// %s -> string
	// %q -> double-quoted string
	// %v -> value (any)
	// %#v -> a Go-syntax representation of the value
	// %T -> value Type
	// %t -> bool (true or false)
	// %p -> pointer (address in base 16, with leading 0x)
	// %c -> char (rune) represented by the corresponding Unicode code point
	a, b, c := 10, 15.5, "Gophers"
	grades := []int{10, 20, 30}

	fmt.Printf("a is %d, b is %f, c is %s \n", a, b, c) // => a is 10, b is 15.500000, c is Gophers
	fmt.Printf("%q\n", c)                               // => "Gophers"
	fmt.Printf("%v\n", grades)                          // => [10 20 30]
	fmt.Printf("%#v\n", grades)                         // => b is of type float64 and grades is of type []int
	fmt.Printf("b is of type %T and grades is of type %T\n", b, grades)
	// => b is of type float64 and grades is of type []int
	fmt.Printf("The address of a: %p\n", &a) // => The address of a: 0xc000016128
	fmt.Printf("%c and %c\n", 100, 51011)    // =>  d and 읃  (runes for code points 101 and 51011)

	const pi float64 = 3.14159265359
	fmt.Printf("pi is %.4f\n", pi) // => formatting with 4 decimal points

	// %b -> base 2
	// %x -> base 16
	fmt.Printf("255 in base 2 is %b\n", 255)  //  => 255 in base 2 is 11111111
	fmt.Printf("101 in base 16 is %x\n", 101) // => 101 in base 16 is 65

	// fmt.Sprintf() returns a string. Uses the same verbs as fmt.Printf()
	s := fmt.Sprintf("a is %d, b is %f, c is %s \n", a, b, c)
	fmt.Println(s) // => a is 10, b is 15.500000, c is Gophers

	// upon saving, constants will flag out errors (compile time error)
	/* <-- star comments only used for debugging
	const num1 = 5
	const num2 = 6
	fmt.Println(num1/num2)
	*/

	const num1, num2 int = 10, 5
	fmt.Println(num1 / num2)

	// Different way of declaring variables
	const n, m int = 4, 5
	const n1, m1 = 6, 7
	const (
		min1 = -500
		min2 //getting value from previous values
		min3
	)

	fmt.Println(min1, min2, min3)

}
