package main

import "fmt"

func main() {
	var divisiveis []int

	for i:=1; i<=100; i++ {
		if ((i%3) == 0) {
			divisiveis = append(divisiveis, i);
		}
	}

	fmt.Println(divisiveis)
}