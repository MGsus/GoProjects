package main

import (
	"fmt"
	"sort"
)

func solve(n, counter int) {
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	sort.Ints(arr)
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] {
			var up int
			up += arr[i-1] - arr[i] + 1
			arr[i] += up
			counter += up
		}
	}

	fmt.Println(counter)
}

func main() {
	var n, counter int
	fmt.Scanf("%d", &n)
	solve(n, counter)
}
