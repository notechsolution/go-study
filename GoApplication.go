package main

import (
	"fmt"
	"runtime"
	"os"
	"./logging"
)

var name = "G";

func main() {
	fmt.Println("Go Hello World!")
	logger.Log("Hello World From Logger module");
	var goos string = runtime.GOOS;
	fmt.Println("go os is:" + goos);
	path := os.Getenv("path");
	fmt.Println("environment path: " + path);
	a := 100;
	b := 200;
	a, b = b, a;
	fmt.Printf("a: %d, b: %d", a, b)
	fmt.Print(name);
	name := "O";
	fmt.Print(name);

	k :=6;
	switch k {
		case 4: fmt.Println("was <= 4"); fallthrough;
		case 5: fmt.Println("was <= 5"); fallthrough;
		case 6: fmt.Println("was <= 6"); fallthrough;
		case 7: fmt.Println("was <= 7"); fallthrough;
		case 8: fmt.Println("was <= 8"); fallthrough;
	default:
		fmt.Println("this is default case");
	}
}
