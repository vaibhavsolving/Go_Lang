package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)



	//can write like this too
	for ; sum < 1000; {    //for loop and only condition initialization above and increment part below.
    sum += sum
	} 
	fmt.Println(sum);             //see yaha Println likha h

	//or 
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum);


	for i := 1; i <= 10; i++ {
		fmt.Printf("2 X %v = %v\n", i, 2*i)   //see yaha Printf likha h coz of %v,
	}
	for i := 1; i <=10; i++{
		fmt.Printf("3 X %v = %v\n", i, i*3)      
	}

	j := 1
	for {
		fmt.Println(j)

		if j == 4 {
			break;
		}
		j++;
	}
}
