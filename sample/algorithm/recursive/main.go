package main

import "fmt"

func main() {
	fmt.Printf("5!=%d\n", factorial(5))
	fmt.Printf("5!=%d\n", factorial2(5))
	fmt.Printf("5!=%d\n", factorial3(5))
	fmt.Printf("5!=%d\n", factorial4(5))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorial2(n int) int {
	stack := []int{n}
	result := 1
	current := 0
	for len(stack) > 0 {
		current = stack[0]
		stack = stack[0:0]
		if current <= 0 {
			current = 1
		} else {
			stack = append(stack, current-1)
		}
		result *= current
	}
	return result
}

func factorial3(n int) int {
	var stack []int

	for i := n; i > 0; i-- {
		stack = append(stack, i)
	}

	result := 1
	for i := 0; i < n; i++ {
		v := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		result *= v
	}
	return result
}

func factorial4(n int) int {
	result := 1
	for i := n; i > 0; i-- {
		result *= i
	}
	return result
}
