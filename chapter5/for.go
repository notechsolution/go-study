package main

import (
	"fmt"
	"strings"
)

func main() {
	for i:=0; i<25;i++ {
		fmt.Println(strings.Repeat("G", i+1));
	}

	for i:=0; i<25;i++ {
		for j:=0; j<=i;j++ {
			fmt.Print("G");
		}
		fmt.Print("\n");
	}
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}

}
