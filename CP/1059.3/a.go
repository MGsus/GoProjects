package main

import "fmt"

func main() {
	var t, n int
	fmt.Scanf("%d", &t)
	var arr []int
	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		var maxOut int
		for j := 0; j < n; j++ {
			var a int
			fmt.Scan(&a)
			if maxOut < a {
				maxOut = a
			}
		}
		arr = append(arr, maxOut)
	}

	for _, v := range arr {
		fmt.Println(v)
	}
}
