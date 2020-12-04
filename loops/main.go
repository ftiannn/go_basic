package main

import "fmt"

func main() {
	// FOR LOOP
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		// or i++ here
	}

	// WHILE LOOP
	j := 10
	for j >= 10 {
		fmt.Println(j)
		j--
	}

	//INFINITE LOOP
	// sum := 0
	// for {
	// 	sum++
	// }
	// fmt.Println(sum) // This line is never reached (Throw error at compile time)

	// FOR LOOP AND CONTINUE
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	// FOR LOOP AND BREAK
	count := 0
	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13\n", i)
			count++
		}
		if count == 10 {
			break
		}
	}
	fmt.Println("=== end of loop ===")
}
