//This is string concatination means
//concatinating values or ovverriding value one upon other

package main

import "fmt"

func main() {
	s := "Microsoft this is a string"
	s2 := "The quick brown fox jumps over the lazy dog"
	fmt.Printf("%v,%T\n\n", s+s2, s+s2)

	//conberting string to bytes //asking for it's ASCii values

	byt := []byte(s2)
	fmt.Printf("%v,%T\n\n", byt, byt)

	//Rune
	//Rune are just type aliases for int32
	//Rune can only take one character for each time
	r := 'e'
	fmt.Printf("%v,%T", r, r)

}
