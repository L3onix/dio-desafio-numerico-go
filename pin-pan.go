package main

import (
"fmt"
"strconv"
)

func main() {
	var frase string

	for i:=1; i<=100; i++ {
		if i%3 == 0 {
			frase += "pin "
		}
		if i%5 == 0 {
			frase += "pan "
		}
		if (i%3) + (i%5) != 0  {
			frase += strconv.Itoa(i) + " "
		}
	}

	fmt.Println(frase)
}
