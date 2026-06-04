package main
import "fmt"

func Add(a int, b int) int {
	fmt.Print("this is sum value of a and b : ")
	return a + b;
	// fmt.Println("this is sum value of a and b")   //Unreachable
}

func Sub(a int, b int) int {
	fmt.Print("this is sub value of a and b : ")      //reachable b/c above to return
	return a - b;
}

//what if multiple parameters have same type?
func Mul(a,b int) int {
	fmt.Print("this is multiple value of a and b : ")
	return a*b;
}


// What  if Multiple return values?
func swapp(a, b string) (string, string) {
	fmt.Print("this is swapped value of a and b : ")
	return b, a;
}

// what if Named return values?
func split(sum int) (a, b int) {                // whatever variable u will use in function define first here
	fmt.Print("this is split value of X : ")
	a = sum * 4/9;
	b = sum - a;
	// c = a + b;            //undefined error get b/c c is not defined in funtion
	return
}


func main() {
	fmt.Println(Add(100,99))
	fmt.Println(Sub(100,99))
	fmt.Println(Mul(100,99))
	a, b := swapp("Hello(1)", "World(2)")
	fmt.Println(a, b);
	fmt.Println(split(100));
}