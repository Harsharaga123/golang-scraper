package main

import (
	"fmt"
	"strings"
	"math"
	"time"
)

// Define a struct
type Person struct {
	Name string
	Age  int
}

func main() {
	// ----- Data Types and Built-in Methods -----
	// int data type
	a := 10
	b := 20
	sum := a + b
	product := a * b
	fmt.Println("Sum:", sum)               // Sum of two integers
	fmt.Println("Product:", product)       // Product of two integers

	// float64 data type
	x := 5.5
	y := 2.3
	avg := (x + y) / 2
	squareRoot := math.Sqrt(x)
	fmt.Println("Average:", avg)           // Average of two floats
	fmt.Println("Square Root of x:", squareRoot) // Square root using math package

	// string data type
	str := "Hello, Golang"
	replacedStr := strings.Replace(str, "Golang", "World", 1)
	length := len(str)
	fmt.Println("Replaced String:", replacedStr) // String replace
	fmt.Println("Length of string:", length)     // Length of the string

	// bool data type
	isTrue := true
	isFalse := !isTrue
	fmt.Println("isTrue:", isTrue)           // Boolean value
	fmt.Println("isFalse:", isFalse)         // Boolean negation

	// ----- Data Structures and Control Structures -----
	// Array data structure
	numbers := [5]int{1, 2, 3, 4, 5}
	arraySum := 0

	// For loop to calculate sum of array elements
	for _, num := range numbers {
		arraySum += num
	}
	fmt.Println("Sum of array:", arraySum)

	// Struct data structure
	person := Person{Name: "Siddharth", Age: 25}

	// If-else control structure
	if person.Age >= 18 {
		fmt.Println(person.Name, "is an adult.")
	} else {
		fmt.Println(person.Name, "is not an adult.")
	}

	// Slice and for loop
	names := []string{"Amit", "Manasa", "Siddharth"}
	for i, name := range names {
		fmt.Printf("Person %d: %s\n", i+1, name)
	}

	// ----- Concurrency Example -----
	// Using Goroutines for concurrency
	go printNumbers()
	go printLetters()

	// Wait for Goroutines to complete
	time.Sleep(6 * time.Second)
	fmt.Println("All tasks completed!")
}

// Function to print numbers (Concurrency)
func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

// Function to print letters (Concurrency)
func printLetters() {
	for _, letter := range "ABCDE" {
		time.Sleep(1 * time.Second)
		fmt.Printf("%c\n", letter)
	}
}
