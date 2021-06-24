//** importent **// extends form 02_variables
//Here a variable container has used to seed up varibale declaration process
//

package main

import "fmt"

var (
	i string  = "This is a way to write variable and string called container"
	y float64 = 66.8
	s string  = "This is an another method to type string and declare variable"
)

func main() {

	//when wnats to see the class of a assigned datatype and variable we can use this method
	fmt.Printf("%v, %T", y, y)
}
