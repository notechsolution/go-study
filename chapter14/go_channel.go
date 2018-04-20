package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("in main()");
	var ownerChannel chan string = make(chan string);
	go sendMessage(ownerChannel);
	go receiveMessage(ownerChannel);
	time.Sleep(3*1e9);
	fmt.Println("end of main()")
}
func receiveMessage(ownerChannel chan string) {
	fmt.Println("ready to receive message");
	firstData := <- ownerChannel;
	fmt.Printf("1:%s\n", firstData);

}
func sendMessage(ownerChannel chan string) {

	fmt.Println("ready to send 1st data");
	ownerChannel <- "I am the first one";
	fmt.Println("ready to send 2nd data");
	ownerChannel <- "I am the second one";
}

