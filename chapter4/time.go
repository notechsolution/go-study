package main

import (
	"fmt"
	"time"
)
func main() {
	t := time.Now()
	// need to use the standard datetime to format `Mon Jan 2 15:04:05 -0700 MST 2006`
	fmt.Printf("Now is %s\n", t.Format("02 Jan 2006 15:04"));
	fmt.Printf("Now is %s\n", t.Format("2006-01-02 15:04"));
}
