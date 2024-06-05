package main

import "fmt"

func numMatch(nums []int, target int) (matchIndex [][]int) {
	var isExist = map[int]bool{}
	for keyA, valA := range nums {
		for keyB, valB := range nums {
			if valA+valB == target && keyA != keyB {
				if !isExist[keyA] && !isExist[keyB] {
					matchIndex = append(matchIndex, []int{keyA, keyB})
					isExist[keyA] = true
					isExist[keyB] = true
				}
			}
		}
	}
	return
}

func main() {
	nums := []int{4, 12, 8, 9, 10, 6, 1}
	match13 := numMatch(nums, 13)
	match16 := numMatch(nums, 16)
	fmt.Println("Indeks 13: ", match13)
	fmt.Println("Indeks 16: ", match16)
}
