package main

import (
	"fmt"
	"math"
)

var c, python, java bool
var i1, j1 int = 1, 2

// fmt.Println(i1, j1)
func main() {
	// var i int  // declared but not assigned, so it will have the zero value of int which is 0
	var i = 42 // declared and assigned, so it will have the value 42
	// var i = 10 or var i int = 10 both are valid.

	var m, k int = 11, 22; // multiple variable declaration and assignment
	fmt.Println(i, c, python, java)  //o/p: 42, false, false,false b/c declared but not assigned
	fmt.Println(m, k)
	fmt.Println(i1,j1) 

	// Short variable declaration (only inside functions)
	a := 555 // short variable declaration, only valid inside functions
	fmt.Println(a)
	fmt.Println(add(3,5));  //e.g of short varibale in function.
	fmt.Println(add2(3,5));

	t, python, java := true, false, "no!"  //short variable declaration...
	fmt.Println(t, python, java)


	//Zero Vaules
	var w int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", w, f, b, s)     //o/p: "c" language ki tarah..

	// Type Conversions
	var x, y int = 3, 4
	var f1 float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f1)
	fmt.Println(x, y,f1, z)

	// Know data type of variable
	c := 42
	fmt.Printf("c is of type %T\n", c);
	ram := "jai shree ram"
	fmt.Printf("ram is of type %T\n", ram);
}

/*
func add(a, b int) (c int) {    // error b/c "C" is already defined as global variable, so we can not use it as return variable in function
	c := 92;
	c += (a + b);
	return c;
}
	*/

	func add(a, b int) int{
		c := 92;
		c += (a + b);
		return c;
	}

	func add2(a, b int) (c int) {
	c = 92;
	c += (a + b);
	return c;
	}
