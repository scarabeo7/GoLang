package main

import (
	"fmt"
)

func main() {

	// Lab 1
	// var myString string
	// fmt.Println("Hello World!")
	// myString = "1234"
	// fmt.Println(myString)

	// fmt.Println("Welcome to GO!")
	// myString = "5678"
	// fmt.Println(myString)

	/*************************************
	*              OR                    *
	*************************************/

	// Lab 2
	// myString := "Hello World!"
	// fmt.Println(myString)
	// myString = "1234"
	// fmt.Println(myString)
	// myString = "Welcome to GO!"
	// fmt.Println(myString)
	// myString = "5678"
	// fmt.Println(myString)

	// Lab 3
	// names := []string{"Daniel", "Rupert", "Emma"}
	// colours := []string{"Red", "Blue", "Green", "Yellow"}

	// // first slice to print names
	// for _, name := range names {

	// 	fmt.Println(name)

	// 	// second slice to print colours
	// 	for _, colour := range colours {
	// 		fmt.Println(colour)
	// 	}

	// }

	// Lab 4
	// myAge := 2000

	// if myAge < 2000 {
	// 	fmt.Println("You were born in the 20th Century")
	// } else if myAge > 2000 {
	// 	fmt.Println("You were born in the 21st Century")
	// } else if myAge == 2000{
	// 	fmt.Println("You were born on the Millenium!")
	// }

	//Lab 5

	// We can create a couple of slices the same way we have seen before
	myFirstSlice := []string{"Daniel", "Rupert", "Emma"}
	mySecondSlice := []string{"Radcliffe", "Grint", "Watson"}

	// if we want to print them out we can create a function that loops through
	// and prints for us rather then duplicate code
	printStringSlice(myFirstSlice)
	printStringSlice(mySecondSlice)

	// lab part
	result := add(20, 30)
	fmt.Println(result) // 50
	result, remainder := divide(10, 3)
	fmt.Println(result)    // 3
	fmt.Println(remainder) // 1

	//Lab 6 func call
	print()
	leapYearChecker(2000)
	
	answer:= subtract(15, 10)
	fmt.Println(answer)
}

// Lab 6
func print() {
	// to use a package function we have to first call the package name,
	// in this case "fmt", then it is followed by the function name
	fmt.Println("Hello World!")

	// fmt gives us more power to write strings than the inbuilt function
	// fmt has a large amount of format chracters: %d in this case is for
	// printing integer arguments
	fmt.Printf("Now you have %d problems, but go ain't %d.\n", 99, 1)
}

// printStringSlice is a function that takes in an arguement of a string slice
// and prints each element on a new line
func printStringSlice(slice []string) {
	for _, elem := range slice {
		fmt.Println(elem)
	}
}

// add will add the two numbers in the arguements and return the result
func add(a, b int) int {
	// as we say we are returning an int, we MUST return one. For now while we
	// write the code we leave it as 0 to ensure we are returning
	// the correct types
	return a + b
}

// divide will divide a by b and return the result and remainder as 2 ints
func divide(a, b int) (result, remainder int) {
	// as we say we are returning multiple int, we MUST return them. For now
	// while we write the code we leave it as 0,0 to ensure we are returning
	// the correct types
	return a / b, a % b

}

// Leap year checker
func leapYearChecker(year int) {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Printf("%d is a leap year \n", year)
	} else {
		fmt.Printf("%d is not a leap year \n", year)
	}
}

func subtract(a, b int) int {
	return a - b 
}
