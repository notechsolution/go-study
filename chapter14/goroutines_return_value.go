package main

import "fmt"

func main() {
	fmt.Println("in the main()");

	var c chan int = make(chan int);
	go func(a int, b int, c chan int){
		fmt.Println("start to sum");
		c <- a +b;
	}(70, 15, c);

	fmt.Println("after call sum routines");
	fmt.Printf("sum result %d\n", <- c);
}
