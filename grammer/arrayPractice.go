package grammer

import "fmt"

func ArrayPoint() {
	//var x [6]int
	//x = [6]int{1,2,3,4,5,6}
	x := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(x[3])
}

func SliceLen() {
	x := make([]int, 3, 9)
	fmt.Println(len(x))
}

func ArrayPractice3() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5])
}

func FindMinNum() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	i := x[0]
	for _, num := range x {

		if i > num {
			i = num
		}

	}
	fmt.Println(i)

}
