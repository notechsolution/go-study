package main

import (
	"fmt"
	"strconv"
	"os"
	"math"
)

func main() {
	var originalString string = "123";
	//var originalString string = "ABC";

	i, err := strconv.Atoi(originalString);
	if err != nil {
		fmt.Println("got error when convert `ABC` to integer: ", err);
		os.Exit(1);
	}
	fmt.Println("convert the string to integer:", i);
	value, ok, msg := mySqrt(-100);
	fmt.Println("mySqrt return ", value, ok, msg);

	// next line not work, as multiple-value return should be assign to single-context
	//value2 := mySqrt(100);
	//fmt.Printf("mySqrt with one return  %s", value2);

}

func mySqrt(input float64) (v float64, ok bool, msg string) {
	if input < 0 {
		return -1, false, "negetive can't sqrt"
	}
	return math.Sqrt(input), true, "success";
}
