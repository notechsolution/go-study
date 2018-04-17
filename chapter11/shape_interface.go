package main

import "fmt"
import "../shaper"

func main() {
	square := shaper.Square{6};
	retangle := &shaper.Retangle{6,8};

	shapers := []shaper.Shaper{square,retangle};

	for n, _ := range shapers {
		fmt.Println("shaper detail:", shapers[n]);
		fmt.Println("shaper area:", shapers[n].Area())
	}

	Classfier(111, shaper.Square{19}, true, shaper.Retangle{3,6}, "hello world")
}

func Classfier(items ...interface{})  {
	for _, item := range items {
		switch item.(type) {
		case bool:
			fmt.Printf("%t is boolean\n", item );
		case shaper.Retangle:
			fmt.Printf("%v is Retangle\n", item );
		case shaper.Square:
			fmt.Printf("%v is Square\n", item );
		default:
			fmt.Printf("don't know the type of %s\n", item );
		}
	}
}
