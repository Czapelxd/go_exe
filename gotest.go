package main

import (
	"fmt"
)

// the func paramiter must be specify and the result too

func add(x, y float64) float64 {
	return x + y
}

const x int = 5

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	// num1, num2 := 5.6, 9.5

	// above shortcut works only 64 byte precision
	// in case of using 32 the var hat to look like:
	// var num1, num2 float32 = 5.6, 9.5

	// fmt.Println(add(num1, num2))

	w1, w2 := "Hey", "there"

	fmt.Println(multiple(w1, w2))

}
