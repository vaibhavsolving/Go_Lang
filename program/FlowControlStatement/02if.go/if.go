package main



//syntax of if statement is....
/*
if initialization; condition {             if initialization; condtion {   }
    // code
}

*/
import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {       //math.Pow(x,n) means x^n;
		return v;
	} else {
		fmt.Printf("%g >= %g\n", v, lim)   //%g is used to print float value in compact form.
	}
	return lim;
}



// calculating square root by newton's method
func Sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("Iteration", i+1, ":", z)
	}

	return z
}



func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"              //recursive call
	}
	return fmt.Sprint(math.Sqrt(x))       //Sprint gives only o/p without printing
}

func main() {
	fmt.Println(sqrt(2), "and", sqrt(-16))
	fmt.Println(pow(4,1,10), pow(4,2,20), pow(4,3,30));
	fmt.Println("Final Answer:", Sqrt(2))
}
