package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3]

	s3[1] = 600 // modify backing array

	fmt.Println(s1) //[10 600 30 40 50]
	fmt.Println(s4) // 	[600 30]

	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]

	arr1[1] = 2                       // change 1 change all
	fmt.Println(arr1, slice1, slice2) //[10 2 30 40] [10 2] [2 30]

	// USING APPEND: DO NOT SHARE SMAE BACKING ARRAY
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...) //newCars and cars doesnt share the same backing arrays

	cars[0] = "Nissan"         // doesnt change newCars as they do not share the same backing array
	fmt.Println(cars, newCars) //[Nissan Honda Audi Range Rover] [Ford Honda]

	// CAPACITY
	n1 := []int{10, 20, 30, 40, 50}
	newSlice := n1[:3]
	fmt.Println(len(newSlice), cap(newSlice)) // 3, 5 (cap = no of element in the backing array)

	newSlice = n1[2:5]                        // {30,40,50}
	fmt.Println(len(newSlice), cap(newSlice)) // 3, 3 (after first slice, cap left with 3)

	fmt.Printf("%p\n", &n1)
	fmt.Printf("%p %p\n", &n1, &newSlice)

	newSlice[0] = 1000
	fmt.Println("n1: ", n1)

	// MEMORY (Array > Slice)
	a := [5]int{1, 2, 3, 4, 5} // ARRAY
	s := []int{1, 2, 3, 4, 5}  // SLICE

	fmt.Printf("array's size in bytes: %d \n", unsafe.Sizeof(a)) //array's size in bytes: 40
	fmt.Printf("slice's size in bytes: %d \n", unsafe.Sizeof(s)) //slice's size in bytes: 24
}
