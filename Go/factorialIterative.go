package main

import "fmt"

func factorial(n int) int {
	sum := 1

	for i := 2; i <= n; i++{
		sum*=i
	}
	return sum
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