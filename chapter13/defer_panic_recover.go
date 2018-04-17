package main

import "fmt"

func main() {
	fmt.Println("main: main func start");
	test("I AM ANOTHER ERROR");
	test("I AM AN ERROR");
	fmt.Println("main: test complete");
}

func test(message string){
	fmt.Println("before start test func");
	defer func(){
		if error := recover(); error !=nil {
			fmt.Printf("Oh my God, error here %s\n", error);
			if error == "I AM AN ERROR"{
				// throw runtime error
				panic(error);
			} else{
				fmt.Println("But still continue the rest");
			}

		}
	}();
	panic(message);
	fmt.Println("after test func");
}
