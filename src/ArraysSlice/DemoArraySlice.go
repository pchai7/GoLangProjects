package main

import "fmt"

func main() {

	// three dots indicates we can make array as long as we want it.
	myArray := [5]int{}
	myArray[0] = 90
	myArray[1] = 2
	myArray[2] = 42
	myArray[3] = 24
	myArray[4] = 99

	mySlice := myArray[:]

	mySlice = append(mySlice, 100)

	fmt.Println(myArray)
	fmt.Println(len(myArray))
	fmt.Println(mySlice)
}