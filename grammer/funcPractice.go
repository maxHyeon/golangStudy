package grammer

import "fmt"

func multiply(w, h int) (int, int) {
	return w * 2, h * 2
}

func MultiplyTest() {
	fmt.Println(multiply(2, 6))
}

func ClosureTest() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 3))
}

func MulReturn() (int, int) {
	return 1, 1
}

func MulParamFuncTest() {
	fmt.Println(MulParamFunc(1, 2, 3, 4, 5, 6, 7, 8))
}
func MulParamFunc(args ...int) int {
	total := 0
	for _, j := range args {
		total += j
	}
	return total
}
