package main

import "fmt"

func main() {
	var l, b, res int
	fmt.Scan(&l, &b)
	for counter := 1; l <= b; counter++ {
		l *= 3
		b *= 2
		res = counter
	}

	fmt.Println(res)
}
