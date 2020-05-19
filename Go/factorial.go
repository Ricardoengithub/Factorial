package main

import "fmt"

func factorial(n int) int {
	if n < 2 {
		return 1
	}else{
		return n * factorial(n-1)
	}
}

func main() {
	fmt.Print("Enter a number: ")
	var n int
	fmt.Scanf("%d", &n)

	if n >= 0 {
		fmt.Println("Factorial de", n, "es", factorial(n))
	}else{
		fmt.Println("Choose another number")
	}
}