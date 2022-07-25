package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func main() {

	fmt.Println("Hello, World!")

	fmt.Println("*** Variables and constants ***")

	var num = 5
	_ = num

	var firstName, lastName, age = "John", "Muir", 55
	fmt.Printf("My name is: %s %s.\n", firstName, lastName)
	fmt.Printf("My age is: %d.\n", age)

	street := "123 Main Street"

	var (
		city  = "Busytown"
		zip   = 12345
		state = "WA"
	)

	// Constant
	const country = "USA"

	fmt.Printf("The address is: %s, %s, %s %d, %s.\n ", street, city, state, zip, country)
	fmt.Println(zip)

	// Use the %T printing verb in the Printf() function to find the type
	// Converting variable types
	var year = "2022"
	fmt.Printf("%T\n", year)
	y, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%T\n", y)
		//fmt.Println("The year is:", y)
		fmt.Printf("The year is %d\n", y)
	}

	fmt.Println("*** if-else ***")

	// Using if-else
	if num%2 == 1 {
		fmt.Println("Number is odd")
	} else {
		fmt.Println("Number is even")
	}

	fmt.Println("*** Functionas and error handling ***")

	// Calling a function
	num1, num2 := 0, 1

	// Calling swap function - pass by value - numbers won't change
	swap(num1, num2)
	fmt.Printf("Num1: %d \n", num1)
	fmt.Printf("Num2: %d \n", num2)

	// Calling swap function - pass by reference - numbers change
	swapByRef(&num1, &num2)
	fmt.Printf("Num1: %d \n", num1)
	fmt.Printf("Num2: %d \n", num2)

	// Calling a function that could return an error
	if v, err := doSomething(); !err {
		fmt.Println(v)
	}

	fmt.Println("*** loops ***")

	// Loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for loop without init and post
	// same as while in other languages
	max := 10
	a, b := 0, 1
	for b <= max {
		fmt.Println(b)
		a, b = b, a+b
	}

	// infinite loop
	// for {
	//}

	fmt.Println("*** Dates ***")

	// Dates
	fmt.Println(time.Now().Format("Mon 2006-01-02 15:04:05"))

	fmt.Println("*** arrays ***")

	// Iterating over a range of values
	// OS is an array of type string with length 3
	OS := [3]string{"iOS", "Android", "Windows"}

	fmt.Println(len(OS))
	for i, v := range OS {
		fmt.Println(i, v)
	}

	fmt.Println("*** Slices ***")

	// Creating an empty slice
	s := make([]int, 5)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	s = append(s, 6, 7, 8)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	// Extracting part of a slice
	fmt.Println(s[6:8])

	// Inserting an item into a slice
	s, error := insert(s, 2, 9)
	if error == nil {
		fmt.Println(s)
	} else {
		fmt.Println(error)
	}

	// Using for-range to iterate through a slice
	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Println("*** Structs ***")
	// See type definition for point
	var pt1 point
	pt1.x = 3.1
	pt1.y = 5.7
	pt1.z = 4.2
	fmt.Println(pt1.x)

	fmt.Println("*** Maps ***")

	fmt.Println("*** Filter Function ***")
	// f is a slice of int
	f := []int{1, 2, 3, 4, 5}

	// Using the filter function to get all even numbers
	evens := filter(f,
		func(val int) bool {
			return val%2 == 0
		})
	fmt.Println(evens)

	fmt.Println("*** Fibonacci using Closure ***")

	// Fibonacci using closure
	gen := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(gen()) // returns 1
	}
}

func doSomething() (int, bool) {
	return 5, false
}

// func swap() with pass by value
func swap(a, b int) {
	a, b = b, a
}

// func swap() with pointer
func swapByRef(a, b *int) {
	*a, *b = *b, *a
}

// The filter() function takes in a slice of int and a function with a single int parameter and return value of type bool
func filter(arr []int, cond func(int) bool) []int {
	result := []int{}
	for _, v := range arr {
		if cond(v) {
			result = append(result, v)
		}
	}
	return result
}

// The fib() function returns a function that returns an int.
// In this case, the fib() function returns a closure (which is an anonymous function)
func fib() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f1, f2 = f2, (f1 + f2)
		return f1
	}
}

func insert(orig []int, index int, value int) ([]int, error) {
	if index < 0 {
		return nil, errors.New("Index canno be less than 0")
	}
	if index >= len(orig) {
		return append(orig, value), nil
	}

	orig = append(orig[:index+1], orig[index:]...)
	orig[index] = value
	return orig, nil
}

type point struct {
	x float32
	y float32
	z float32
}
