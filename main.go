package main

import "fmt"

func main() {
	fmt.Println(fib(5)) // ใช้ 'Println' ที่ตัว 'L' เป็นตัวใหญ่
}

func fib(n int) int {
	// ฟังก์ชัน Fibonacci
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
