package main

import (
	"fmt"
	"strings"
	"time"
	Logger "../logging"
)

var debug = Logger.Where();

func main() {
	//fmt.Println("is GOLAND has index of O is", strings.IndexFunc("GOLAND", mapping));
	start := time.Now();
	debug();
	var value string = strings.Map(mapping, "明珠北路64-3-2-804");
	fmt.Println("converted to :", value);
	fmt.Println("replaced to :", strings.Replace(value, "?", "", -1));
	time.Sleep(2000000000);
	end :=time.Now();
	fmt.Printf("main() took amount time: %d", end.Sub(start)/1000);
}

func mapping(value rune) rune {
	if value > 127 {
		debug();
		return '?';
	}
	return value;
}

