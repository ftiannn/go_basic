package main

import "fmt"

func main() {
	outer := 19
	_ = outer

	people := [5]string{"Helen", "Mark", "Brenda", "Andy", "Michael"}
	friends := [2]string{"Mark", "Mary"}

outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("=== after break ===")

	i := 0

	// CREATE A LOOP LIKE FOR-LOOP
loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	// ERROR it's not permitted to jump over the declaration of x
	// 	goto todo
	// 	x:=5
	// todo:
	// 	fmt.Println("bla bla")

	// SWITCH CASE
	language := "golang"

	switch language {
	case "Python":
		fmt.Println("Don't use any curly braces. Uses indentation")
	case "Go", "golang":
		fmt.Println("Uses curly braces {}")
	default:
		fmt.Println("What... english?")
	}

	n := 5
	// comparing the result of an expression which is bool to another bool value
	switch true {
	case n%2 == 0:
		fmt.Println("Even!")
	case n%2 != 0:
		fmt.Println("Odd!")
	default:
		fmt.Println("Never here!")
	}

	//** Switch simple statement **//

	// Syntax: statement (n:=10), semicolon and a switch condition
	//(true in this case, we are comparing boolean expressions that return true)
	// we can remove the word "true" because it's the default
	switch n := 10; true {
	case n > 0:
		fmt.Println("Positive")
	case n < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}

}
