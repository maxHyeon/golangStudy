package grammer

import "fmt"

func MapPrac() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Print(x["key"])
}

func DictPractice() {
	elem := make(map[string]string)
	elem["A"] = "ABC"
	elem["B"] = "BCD"
	elem["C"] = "CDE"

	fmt.Print(elem["C"])

	vlue, ok := elem["D"]
	fmt.Println(vlue, ok)
	if vlue, ok := elem["A"]; ok {
		fmt.Println(vlue, ok)
	}
}

func DoubleDictPractive() {
	elem2 := map[string]map[string]string{
		"A": map[string]string{
			"name":   "A",
			"number": "5",
		},
		"B": map[string]string{
			"name":   "B",
			"number": "5",
		},
	}
	if el2, ok := elem2["A"]; ok {
		fmt.Print(el2["name"], el2["number"])
	}
}
