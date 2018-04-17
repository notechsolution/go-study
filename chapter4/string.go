package main

import (
	"fmt"
	"strings"
)
func main() {
	// Prefix & Suffix
	fmt.Printf("Does Golang has prefix %s ", "Go");
	fmt.Println( strings.HasPrefix("Golang", "Go"))

	// Contains
	fmt.Println("Does Golang contains lan ", strings.Contains("Golang", "lan"));

	// Replace
	fmt.Println("Google replace the o as ", strings.Replace("Google", "o", "0", -1));
}
