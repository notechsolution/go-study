package main

import (
	"fmt"
	"../shaper"
	"reflect"
)

func main() {
	retangle := shaper.Retangle{6, 8};
	fmt.Printf("retangle %v, Area(): %d\n", retangle, retangle.Area());
	r := reflect.ValueOf(retangle);

	// 1. reflect type
	fmt.Printf("type:%s\n", r.Type());
	fmt.Printf("kind:%s\n", r.Kind());

	// 2. reflect fields
	retanglePointer := &shaper.Retangle{8, 8};
	pointerReflect := reflect.ValueOf(retanglePointer);
	fmt.Printf("numMethods:%d\n",  pointerReflect.NumMethod());

	square := shaper.Square{8};
	sr := reflect.ValueOf(square);
	fmt.Printf("numFields:%d, numMethods:%d\n", sr.NumField(), sr.NumMethod());
}
