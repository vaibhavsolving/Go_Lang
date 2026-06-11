package main


import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}


func Pic(dx, dy int) [][]uint8 {          //this is exrecise of slices
	result := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		result[y] = make([]uint8, dx)

		for x := 0; x < dx; x++ {
			result[y][x] = uint8(x * y)
		}
	}

	return result
}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//Range continued
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)              // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	fmt.Println(Pic(5, 5))
}
