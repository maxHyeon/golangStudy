package array

import (
	"fmt"
)

func grammerParctice() {
	sum, i := 0, 0
	for i < 10 {
		if i%2 == 0 {
			fmt.Println("even Nuber")
		} else {
			fmt.Println("odd Number")
		}
		switch {
		case i == 0:
			fmt.Println("First Loop")
		}
		sum += i
		i++
		// no ++i
		/*
			multi line comment
		*/
	}
	fmt.Println(sum)
}
