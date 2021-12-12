package common

import (
	"fmt"
)

func PrettyPrintMatrix(ints [][]int) {
	for i := range ints {
		for j := range ints[i] {
			fmt.Printf("%d", ints[i][j])
		}
		fmt.Printf("\n")
	}
}

func PrettyPrintBools(matr [][]bool) {
	for i := range matr {
		for j := range matr[i] {
			if matr[i][j] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
