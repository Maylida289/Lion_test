package main

import "fmt"

func FindMinAndMax(arr []int) string {
	// your code here
	max := arr[1]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}

	}
	return fmt.Sprintf("nilai max nya : %d", max)
}

func main() {
	//contoh soal
	fmt.Println("[3, 5, 1, 9, 2]\n", FindMinAndMax([]int{3, 5, 1, 9, 2}))
	fmt.Println("")

	//contoh lainnya
	fmt.Println("[2, -5, -4, 22, 7, 7]\n", FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	fmt.Println("")
	fmt.Println("[200, 0, 11, 120, 35, 110]\n", FindMinAndMax([]int{200, 0, 11, 120, 35, 110}))
	fmt.Println("")

}
