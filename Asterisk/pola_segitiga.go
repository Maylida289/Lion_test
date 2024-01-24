package main

import "fmt"

func triangle(n int) {
	for i := 0; i <= n; i++ {
		for j := 4; j >= i; j-- {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Print(" ")
		}
		for l := 0; l <= i; l++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	triangle(5)
}
