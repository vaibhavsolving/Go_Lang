package main

import "fmt"

var c, python, java bool

func main() {
	// var i int  // declared but not assigned, so it will have the zero value of int which is 0
	var i = 42 // declared and assigned, so it will have the value 42
	// var i = 10 or var i int = 10 both are valid.
	var j, k int = 1, 2; // multiple variable declaration and assignment

	fmt.Println(i, c, python, java)  //o/p: 42, false, false,false b/c declared but not assigned
	fmt.Println(j, k)
}
