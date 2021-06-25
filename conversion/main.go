//Method of conversion

package main

import "fmt"

func main() {

	// This how we can declear a variable and it's value
	var i int = 66
	fmt.Printf("%v,%T", i, i)
	//This is how we can assign that value to another variable /// This is callaed conversion
	var j float64
	j = float64(i)

	fmt.Printf("%v,%T", j, j)
}
