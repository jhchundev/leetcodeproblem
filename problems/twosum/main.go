package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for j := i + 1; j < len(nums); j++ {
			num2 := nums[j]
			if num+num2 == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	fmt.Println("Hello world!")
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
