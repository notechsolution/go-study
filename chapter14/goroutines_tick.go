package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("go routines with tick - main function");
	tickChannel := time.Tick(1e9);
	 i :=0;
	for {
		<- tickChannel;
		go maybeHttpCall(i);
		i++;
	}
}
func maybeHttpCall(index int) {
	fmt.Println("call ", index);
}
