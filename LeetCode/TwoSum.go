package main

import "fmt"

// program ini mengguna Brute Force
func main() {
	// twoSum memiliki dua parameter yaitu nums dengan bentuk array integer dan target dengan bentuk int
	result := twoSum([]int{2, 7, 11, 15}, 9)
	// Print hasil
	fmt.Println(result)
}

// function twoSum memiliki dua parameter dan return
func twoSum(nums []int, target int) []int {

	// inisialisasi perulangan dengan
	// nilai awal = 0
	// nilai 0 akan dikurangi banyaknya nilai pada array num dan lakukan penambahan value
	for i := 0; i < len(nums); i++ {
		// nilai j = nilai i + 1
		// nilai j akan dikurangi banyaknya nilai pada array num dan lakukan penambahan value
		for j := i + 1; j < len(nums); j++ {
			// jika nilai array i ditambah nilai array j cocok dengan target
			// maka return hasil tersebut
			// yang returnnya menghasilkan array index i dan j
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nums
}
