package grammer

import (
	"fmt"
)

func Array() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
}

func CalAvg() {
	var x [6]float64
	x[0] = 100
	x[1] = 93
	x[2] = 60
	x[3] = 53
	x[4] = 72
	x[5] = 44
	// x := [3]float64 { 10,20,40} Multi line able

	var total float64 = 0
	//for i :=0; i<len(x); i++{
	//	total += x[i]
	//}
	for _, vlu := range x {
		total += vlu
	}
	fmt.Print(total / float64(len(x)))
}

func SlicePractice() {
	//var x []float64

	//x := make([]float64, 6)
	arr := []float64{1, 2, 3, 4}
	x := arr[0:]
	arr2 := make([]float64, 2)
	xC := arr2[0:]
	copy(xC, x)
	fmt.Print(x)
	fmt.Print(xC)
	x2 := append(x, 4, 5)
	fmt.Print(x2)

	//arr := []float64(1,2,3,4)
	//x := arr[0:5]
	//x := arr[:]
}

//https://www.codingnuri.com/golang-book/6.html
