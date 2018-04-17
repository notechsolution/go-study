package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var firstName, lastName string;
	fmt.Println("Please input your name:");
	fmt.Scanln(&firstName, &lastName);
	fmt.Printf("firstName: %s, lastName: %s \n", firstName, lastName);

	newReader := bufio.NewReader(os.Stdin);
	input,_:= newReader.ReadString('\n');
	fmt.Printf("your input is %s\n", input);
}
