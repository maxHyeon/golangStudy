package grammer

import "fmt"

func CodeConventionInfo() {

	fmt.Println("GoLang CODE Convention")
	// this code's reference in the 'TheBook's Go lang web Programming introduce'

	// golang use camelCase
	// golang's Var orient simple word but restric var's scope really small
	// many var in golang use one alphabet
	// see below sample code

	//func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	//	p := newPrinter()
	//	p.doPrint(a, true, true )
	//	n, err = w.Write(p.buf)
	//	p.free()
	//	return
	//}

	// gatter, setter
	// gatter func didin't attache get, use naun
	//func Connection() *Conn { // didn't use GetConnection()
	//	//
	//	return conn
	//}
	// Setter use Set

}
