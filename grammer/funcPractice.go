package grammer

import "fmt"

func multiply(w, h int) (int, int) {
	return w * 2, h * 2
}

func MultiplyTest() {
	fmt.Println(multiply(2, 6))
}
