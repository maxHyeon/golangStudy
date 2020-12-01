package grammer

import (
	"fmt"
)

func StatementParctice() {

	//goLang hasn't While Statement

	// repeat statement with i
	sum0 := 0
	for i := 0; i < 10; i++ {
		sum0 += i
	}
	fmt.Println(sum0)

	// repeat statement with only condition
	sum, i := 0, 0
	for i < 10 {
		if i == 10 {
			break
		}

		if i%2 == 0 {
			fmt.Println("even Nuber")
		} else {
			fmt.Println("odd Number")
		}

		// switch with value
		switch i {
		case 0:
			fmt.Println("First Loop")
			fallthrough
		case 3:
			fmt.Println("fallthrough test")
		case 1, 2:
			fmt.Println("use two var")
		default:
			fmt.Println("thits is default")
		}

		// switch with condition statement
		switch { // didin't write var
		case i == 0:
			fmt.Println("First Loop")
			fallthrough
		case i < 1:
			fmt.Println("fallthrough test")
		case i < 5, i > 2:
			fmt.Println("use two var")
		default:
			fmt.Println("thits is default")
		}
		sum += i
		i++
		// no ++i
		/*
			multi line comment
		*/
	}

	fmt.Println(sum)

	//label
	x := 7
	table := [][]int{{1, 5, 9}, {2, 6, 5, 8}, {4, 5, 9, 10}}
LOOP: // label
	for row := 0; row < len(table); row++ {
		for col := 0; col < len(table[row]); col++ {
			if table[row][col] == x {
				fmt.Printf("found %d (row:%d, col:%d))\n", x, row, col)
				break LOOP
				// you can use continue
				// if use continue
				// this statement perform next loop declare label
				//continue LOOP
			}
		}
	}
	//label can use with goto, switch, select

	//sum,i := 0,0
	//for i < 10 {
	//	sum += i++ // ++ plus minus calc hasn't return so this statement accrue compile error
	//  ++i        // golang isn't allowed pre calc plusminus
	//}
	//fmt.Println(sum)
}
