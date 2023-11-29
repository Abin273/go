package main

import (
	"errors"
	"fmt"
)

// Numbers
var var1 int
var var2 int = 200
var var3 uint // unsigned integer

var var4,var5 int // Multi declaration

var var6 float32
var var7 float64

// complex numbers a+bi
var var8 complex64
var var9 complex128

// String
var var10 string
var str string = "hello"

// Boolean
var var11 bool

// Constants
const RED = 0

const (
	BUS = iota  // 0
	CAR // 1
	BIKE // 2
)



func main(){
	var1 = 10
	fmt.Println(var1)
	fmt.Println(str)

	// Constants
	fmt.Println(BUS)
	fmt.Println(CAR)
	fmt.Println(BIKE)

	// short declaration
	c := 3000
	fmt.Println(c)

	// Type conversion
	var4 = 50
	var6 = float32(var4)

	// Multi variable
	a,b := 10,20
	fmt.Printf("%v + %v = %v\n", a, b, a*b)

	// Errors
	err := errors.New("Error occured 1234")
	fmt.Println(err)


}
