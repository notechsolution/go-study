package logger

import (
	"fmt"
)

func Log(message string) {
	print(message)
}

func print(message string) {
	fmt.Println("use internal print " + message)
}


