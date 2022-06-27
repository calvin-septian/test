package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Word")

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println("ganjil")
		} else {
			fmt.Println("genap")
		}
	}

	name := []string{"andi", "budi", "cacing"}
	for _, v := range name {
		fmt.Println(v)
	}

	biodata()
}
