package main

import (
	"fmt"
	"unsafe"
)
func main() {
	const carriers int = 100;
	const pi float32 = 3.1415926
	const pi64 float64 = 3.1415926
	fmt.Printf("const carriers[int] size is %d bytes \n", unsafe.Sizeof(carriers) )
	fmt.Printf("const pi[float32] size is %d bytes \n", unsafe.Sizeof(pi) )
	fmt.Printf("const pi[float64] size is %d bytes \n", unsafe.Sizeof(pi64) )
}
