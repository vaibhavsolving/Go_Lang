//A defer statement defers the execution of a function 
// until the surrounding function returns.

package main

import "fmt"

func main() {
	defer fmt.Println("world")     //will run at the end, wait for surrounding

	fmt.Println("hello")


	//stacking defers
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)       //will run in reverse order, last in first out
	}

	fmt.Println("done")
}
