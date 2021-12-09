package common

import "fmt"

func PrettyPrintInts(ints [][]int) {
	for i := range ints {
		for j := range ints[i] {
			fmt.Printf("%d", ints[i][j])
		}
		fmt.Printf("\n")
	}
}
