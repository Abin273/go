package main

import "fmt"

func main(){
	fmt.Println("Array sample")

	// We cant change array size after declaration

	// Eg 1

	var arr1 [3]string // Initial vlaue will be empty array. ie, []
	arr1 = [3]string{"abi", "amal", "anu"}
	fmt.Println("arr1: ",arr1)

	// Eg 2
	arr2 := [3]string{ "bus", "car", "bike" }
	fmt.Println("arr2: ", arr2)

	// Eg 3
	arr3 := arr2 // This is a copy operation , in go here we are not assigning reference or pointer of arr2 to arr3 
	fmt.Println(arr3)

	arr2[1] = "aeroplane"
	fmt.Println("arr2: ",arr2)
	fmt.Println("arr3: ",arr3)

	// Integer array demo

	var intArray [3]int
	fmt.Println("intArray", intArray)
	fmt.Println("intArray", intArray[0])
	
	intArray[1] = 403
	fmt.Println("intArray", intArray)

	// find length of array
	fmt.Println("Length of inArray is : ",len(intArray))

}