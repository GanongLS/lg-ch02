package main

import "fmt"

// remove either type conversion and the code will not compile!
func main() {
	var x int = 10
	var y float64 = 30.2
	// converts the int to a float64, so it can be added to a float64
	var z float64 = float64(x) + y
	// converts the float64 to an int, so it can be added to an int
	var d int = x + int(y)
	fmt.Println(z, d)
}
