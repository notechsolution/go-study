package main

import (
	"fmt"
)

func main() {
	var name string = "Golang";
	fmt.Printf("Golang substr %s\n", name[0:2])

	nameA := []byte(name);
	nameA[0] = 'D';
	fmt.Printf("Golang change to %s\n", nameA)
}
