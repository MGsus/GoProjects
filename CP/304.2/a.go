package main

import "fmt"

func main() {
	var x, cost, usd, totalCost int
	fmt.Scanf("%d %d %d", &cost, &usd, &x)
	for i := 1; i <= x; i++ {
		totalCost += cost * i
	}
	if totalCost-usd < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(totalCost - usd)
	}

}
