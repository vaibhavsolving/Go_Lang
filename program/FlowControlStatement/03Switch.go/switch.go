

/*

switch os {
case "linux":
	fmt.Println("Linux")
case "darwin":
	fmt.Println("macOS")
default:
	fmt.Println("Other")
}

---> they are equivalent 

if os == "linux" {
	fmt.Println("Linux")
} else if os == "darwin" {
	fmt.Println("macOS")
} else {
	fmt.Println("Other")
}
*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	day := 5

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend")
	}

//2nd example of switch statement

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

//3rd example of switch statement
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	//switch without condition is same as switch true
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}