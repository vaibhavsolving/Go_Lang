package main

import "fmt"

type Vertex12 struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex12{1, 2})

	v := Vertex12{
		X: 3,
		Y: 4,
	}
	fmt.Println(v)
	v.X = 4
	v.Y = 777
	fmt.Println(v.X, "and" , v.Y);

	//pointer to struct
	P := &v
	v.X = 1e9;
	fmt.Println(*P)
	fmt.Println(v.X, "and" , v.Y);
	fmt.Println(P);

	fmt.Println(v1, p, v2, v3)    //struct literals

}

// Struct literals (changing struct values)
var (
	v1 = Vertex12{1, 2}  // has type Vertex
	v2 = Vertex12{X: 1}  // Y:0 is implicit
	v3 = Vertex12{}      // X:0 and Y:0
	p  = &Vertex12{1, 2} // has type *Vertex
)
