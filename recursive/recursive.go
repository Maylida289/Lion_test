package main

import "fmt"

func rekursif(n int) int {

	if n <= 1 {
		return 1
	}
	return n * rekursif(n-1)
}

func main() {
	//contoh soal
	fmt.Println("rekursif dari 5 adalah", rekursif(5))

	//contoh lainnya
	fmt.Println("rekursif dari 10 adalah", rekursif(10))
	fmt.Println("rekursif dari 25 adalah", rekursif(25))
}
