/*  An integer value identifying a Unicode code point.
    A rune literal is expressed as one or more characters enclosed in single quotes, as in 'x' or '\n'.
*/

package main

import (
	"fmt"
)

func main() {
	const word = "世界"
	bytes := []byte(word)
	fmt.Println(bytes)
	var x []rune
	x = []rune{'世', '界'}
	xBytes := []byte(string(x))
	fmt.Println(xBytes)
}
