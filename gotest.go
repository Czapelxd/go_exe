package main

// the func paramiter must be specify and the result too

func add(x, y float64) float64 {
	return x + y
}

const x int = 5

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {

	var a int = 62
	var b float64 = float64(a) //this alowse to convert var from to another

	x := a // x will be type int

	// num1, num2 := 5.6, 9.5

	// above shortcut works only 64 byte precision
	// in case of using 32 the var hat to look like:
	// var num1, num2 float32 = 5.6, 9.5

	// fmt.Println(add(num1, num2))

	// w1, w2 := "Hey", "there"

	// fmt.Println(multiple(w1, w2))

}
