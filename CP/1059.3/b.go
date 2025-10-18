package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)

	var output strings.Builder

	for ; t > 0; t-- {
		var n int
		var s string
		fmt.Fscan(in, &n)
		fmt.Fscan(in, &s)

		// Case: empty subsequence works
		if isPalindrome(s) {
			output.WriteString("0\n")
			continue
		}

		found := false

		// Try all possible counts of 0s (from start) and 1s (from end)
		for i := 0; i <= n && !found; i++ {
			for j := 0; j <= n && !found; j++ {
				pIndices := []int{}

				// take first i zeros
				count0 := 0
				for idx := 0; idx < n; idx++ {
					if s[idx] == '0' && count0 < i {
						pIndices = append(pIndices, idx+1)
						count0++
					}
				}

				// take last j ones
				count1 := 0
				for idx := n - 1; idx >= 0; idx-- {
					if s[idx] == '1' && count1 < j {
						pIndices = append(pIndices, idx+1)
						count1++
					}
				}

				// remove selected to form x
				mark := make([]bool, n)
				for _, v := range pIndices {
					mark[v-1] = true
				}

				x := make([]byte, 0, n)
				for k := 0; k < n; k++ {
					if !mark[k] {
						x = append(x, s[k])
					}
				}

				if isPalindrome(string(x)) {
					output.WriteString(fmt.Sprintf("%d\n", len(pIndices)))
					if len(pIndices) > 0 {
						for k := 0; k < len(pIndices); k++ {
							if k > 0 {
								output.WriteString(" ")
							}
							output.WriteString(fmt.Sprint(pIndices[k]))
						}
						output.WriteString("\n")
					}
					found = true
				}
			}
		}

		if !found {
			output.WriteString("-1\n")
		}
	}

	fmt.Print(output.String())
}
