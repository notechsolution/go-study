package main

import (
	"fmt"
)

func main() {
	var array [5]string;
	// value assignment
	arrayCopy := array;
	arrayPointer := &array
	fmt.Printf("array address %p \n", &array );
	fmt.Printf("arrayCopy address %p \n", &arrayCopy );
	fmt.Printf("arrayPointer address %p \n", &arrayPointer );
	array[0] = "arrayIndex0";
	arrayCopy[0] = "arrayCopyIndex0";
	arrayPointer[1] = "arrayPoinerIndex1";
	fmt.Printf("array value %s \n", array );
	fmt.Printf("arrayCopy value %s \n", arrayCopy );
	fmt.Printf("arrayPointer value %s \n", arrayPointer );


	items := [...]int{10, 20, 30, 40, 50};
	for _, item := range items {
		item *= 2
	}

	fmt.Printf("item %d", items);
}
