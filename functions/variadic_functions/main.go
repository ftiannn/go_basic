package main

import (
	"fmt"
	"strings"
)

// creating a variadic function
func f1(a ...int) {
	fmt.Printf("%T\n", a) // => []int, slice of int
	fmt.Printf("%#v\n", a)
}

// variadicfunction that modifies one of the arguments passed.
func f2(a ...int) {
	a[0] = 50
}

// creating a variadic function that calculates and returns the sum and product of its arguments
func sumAndProduct(a ...float64) (float64, float64) {
	sum := 0.     // float
	product := 1. // float

	for _, v := range a {
		sum += v
		product *= v
	}

	return sum, product
}

// mixing variadic and non-variadic parameters is allowed
// non-variadic parameters are always before the variadic parameter
func personInformation(age int, names ...string) string {
	fullName := strings.Join(names, " ")

	returnString := fmt.Sprintf("Age: %d, Full Name: %s", age, fullName)

	return returnString
}

func main() {
	// calling variadic functions
	// a variadic function can be invoked with zero or more arguments
	f1(1, 2, 3, 4)
	f1() // a is: []int(nil)

	// an example of a well-known variadic function is append() built-in function.
	// it appends items to an existing slice and returns back the same slice.
	nums := []int{1, 2}
	nums = append(nums, 3, 4, 5)

	f1(nums...)

	f2(nums...)                 // amending a[0] to 50
	fmt.Println("Nums: ", nums) //Nums:  [50 2 3 4 5]

	s, p := sumAndProduct(2., 5., 10.)
	fmt.Println(s, p) // -> 17 100

	info := personInformation(30, "Wolfgang", "Amadeus", "Mozart")
	fmt.Println(info) // => Age: 30, Full Name:Wolfgang Amadeus Mozart

}
