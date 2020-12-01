package grammer

import "fmt"

func VarPractice() {
	//init var
	//if didn't assign var, golang init zero value,
	// int = 0, float = 0.0, string = ""
	var a int
	var b string
	fmt.Println("init zerovalue: ", a, b)
	a = 1
	b = "hi"
	fmt.Println("init user value: ", a, b)

	//multiple var init
	// var name1, id1, address1 string
	//var (
	//	name string
	//	age int
	//	weight float32
	//
	//)

	// if var assign and init in sametime, you can skip declare type
	// var c = true
	// var testInt = 1

	// if you want skip var, use :=
	//testInt2 := 2

	// you can't use := when aleady assigned var, only use new var
	// and only use in func, can't use global var
	// testInt2 := 3

	//const
	const limit = 64
	const max int64 = 1024
	const max2 = 1024 * 1024
	//can't use calc in compile time
	// const c = getNumber()
	//multiple const
	const (
		cNum1 = 1
		cNum2 = 2
		cStr1 = "Thank"
	)
	//enum
	// golang didn't divide enum and const but golang has iota reservation keyword
	// simple declare enum with const
	const (
		SUN = 0
		MON = 1
		TUE = 2
		WED = 3
	)

	//use iota reserved keyword
	const (
		SUN1 = iota //start ZERO
		MON1
		TUE1
		WED1
	)

	type Color int

	const (
		RED Color = iota
		BLUE
		GREEN
	)
	// iota can use float type and combine with calc
	type ByteSize int64
	const (
		_           = iota             // assign 0
		KB ByteSize = 1 << (10 * iota) //1 << (10 * 1) 1024 -> '<<' is left shift
		MB                             //1 << (10 * 2) 1024 * 1024 -> '<<' is left shift
		GB                             // ..
	)

	const (
		DEFAULT_RATE = 5 + 0.3*iota // 5
		GREEN_RATE                  // 5 + 0.3 *1 = 5.3
		SILV_RATE                   // 5 + 0.3 *2 = 5.6
	)

}
