package main

import "fmt"

func twoSumMap(nums []int, target int) []int {
	sz := len(nums)
	swapmap := make(map[int]int)
	for i := 0; i < sz; i++ {
		resto := target - nums[i]
		_, exists := swapmap[resto]
		if exists {
			return []int{swapmap[resto], i}
		}
		swapmap[nums[i]] = i
	}
	return nil
}

func main() {
	var a, b, c, target int

	fmt.Scanf("%d %d %d %d", &a, &b, &c, &target)

	nums := []int{a, b, c}

	res := twoSumMap(nums, target)
	fmt.Println(res)
}
