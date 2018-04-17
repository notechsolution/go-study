package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i:=0; i< 100; i++ {
		fmt.Printf("randon is %2.2f \n", rand.Float32());
	}
}
