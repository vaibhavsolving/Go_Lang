package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // Declare pointer P, point to i, p is storing memeory address of i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j, Now p is assigned the address of j, p no longer points to i
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
