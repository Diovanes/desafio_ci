package main

import (
	"fmt"
)

func main() {
	// value1, err := strconv.ParseInt(os.Args[1], 10, 8)
	// value2, err := strconv.ParseInt(os.Args[2], 10, 8)

	// if err != nil {
	// 	fmt.Printf("fail")
	// } else {
	// 	fmt.Printf("%d\n", Soma(int(value1), int(value2)))
	// }
	fmt.Printf("%d\n", Soma(5, 5))
}

func Soma(x, y int) int {
	return x + y
}
