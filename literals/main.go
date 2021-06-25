//This section has taekn for lterals

package main

import "fmt"

func main() {
	s := "Microsoft"

	fmt.Printf("%v,%T\n", s[2], s[2])

	//to know the letter than the mathematical value
	//important---- we can't manipulate the value of the string
	fmt.Printf("%v,%T", string(s[2]), string(s[2]))
}
