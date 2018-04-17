package logger

import (
	"runtime"
	"fmt"
)

func Where() func(){
	return func (){
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d\n", file, line);
	}

}
