package main

import "fmt"

func main() {
	// ** METHOD 1 without struct **//
	// title1, author1, year1 := "The Devine Comedy", "Dante Aligheri", 1320
	// title2, author2, year2 := "Macbeth", "William Shakespeare", 1606

	// fmt.Println("Book1: ", title1, author1, year1)
	// fmt.Println("Book2: ", title2, author2, year2)

	//** METHOD 2: STRUCT **//
	// creating a struct type
	type book struct {
		title  string
		author string
		year   int
	}

	// combining different fields of the same type on the same line
	type book1 struct {
		title, author string
		year          int
	}

	// declaring, initializing and assigning a new book value, all in one step
	lastBook := book{"The Divine Comedy", "Dante Aligheri", 1320} //this is a struct literal and order matters
	fmt.Println(lastBook)

	// Declaring a new book value by specifying field: value (order doesn't matter)
	bestBook := book{title: "Animal Farm", year: 1945, author: "George Orwell"}
	_ = bestBook

	//if we create a new struct value by omitting some fields they will be zero-valued according to their type
	aBook := book{title: "Just a random book"}
	fmt.Printf("%#v\n", aBook) // => main.book{title:"Just a random book", author:"", year:0}

	// retrieving the value of a struct field
	fmt.Println(lastBook.title) // => The Divine Comedy
	// page := lastBook.pages // => error: unknown field

	// updating a field
	lastBook.author = "Leo Tolstoy"
	lastBook.year = 1878
	fmt.Printf("lastBook: %+v\n", lastBook) // => lastBook: {title:The Divine Comedy author:The Best year:2020}
	// + modifier with %v  printed out both the field names and their values

	// comparing struct values
	// two struct values are equal if their corresponding fields are equal.
	randomBook := book{title: "Random Title", author: "John Doe", year: 100}
	fmt.Println(randomBook == lastBook) // => false

	// = creates a copy of a struct
	myBook := randomBook
	myBook.year = 2020              // modifying only myBook
	fmt.Println(myBook, randomBook) // => {Random Title John Doe 2020} {Random Title John Doe 100}

}
