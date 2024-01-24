package main

import "fmt"

func palindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}
	return true
}
func main() {
	//contoh soal
	fmt.Println(palindrome("radar"))
	fmt.Println(palindrome("hello"))

	//contoh lainnya
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("kita"))

}
