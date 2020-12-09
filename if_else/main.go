package main

import "fmt"

func main() {
	price, inStock := 100, true
	if price > 80 {
		fmt.Println("Too expensive!")
	}

	if price <= 100 && inStock {
		fmt.Println("Buy It")
	}

	// if price {} --> Not valid in GO

	// only one if branch will be executed
	if price < 100 {
		fmt.Println("It's cheap!")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else { //executed only once if all the if branches are false (it's optional)
		fmt.Println("It's Expensive!")
	}

}
