package main

import "fmt"

func main() {
	fmt.Println("is number  = ", fib(9))
	fmt.Println("Hi Pop os", "Hello")
}

func fib(n int) int {
	//Fibonacci
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
