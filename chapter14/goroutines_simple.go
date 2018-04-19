package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("in Main()");
	go longWait();
	go shortWait();
	fmt.Println("about to sleep in main()");
	time.Sleep(10*1e9);
	fmt.Println("at the end of main()");
}
func shortWait() {
	fmt.Println("beging short wait()");
	time.Sleep(2 * 1e9);
	fmt.Println("end of shortWait()")
}
func longWait() {
	fmt.Println("beging long wait()")
	time.Sleep(6*1e9);
	fmt.Println("end of longWait()");
}
