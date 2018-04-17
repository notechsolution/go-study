package main

import (
	"fmt"
)

func main() {
	pf := func(message string) {fmt.Println(message)};
	pf("Hello World");
	fmt.Printf("pf is an %T with value %v \n", pf, pf);
	fmt.Println("f() returns ", f());


	f := adder();
	fmt.Printf("first time: %v-", f(1));
	fmt.Printf("second time: %v-", f(20));
	fmt.Printf("first time: %v\n", f(120));

	for i:=0; i<=10; i++ {
		fmt.Printf("fibonacci(%d) is %d\n", i, fibonacci(i));
	}
	for i:=0; i<=10; i++ {
		fmt.Printf("fibonacci1(%d) is %d\n", i, fibonacci1(i));
	}
	fmt.Printf("fibonacci1(%d) is %d\n", 5, fibonacci1(5));
}

func f() (ret int){
	defer func () {ret++}();
	return 1;
}


func adder() (func (b int) int) {
	var x int;
	return func (b int) int {
		x +=b;
		return x;
	}
}

//f(3) = f(2)+f(1);
// f(n) = f(n-1) + f(n-2)

func fibonacci(n int) (res int) {
	var i_1, i_2 int;
	for i:=0; i<=n; i++ {
		if i<= 1{
			i_1 = 1;
			i_2 = 1;
			res = 1;
		} else {
			res = i_1+i_2;
		}
		i_2, i_1 = i_1, res;
	}
	return;
}

func fibonacci1(n int) (res int){
	f:= fibonacciInClosure();
	for i:=0; i<=n; i++ {
		res = f();
	}
	return;
}

func fibonacciInClosure() (func () int) {
	var i_1, i_2 = 1, 1;
	return func() (res int){
		res = i_2;
		i_2, i_1 = i_1, (i_1+i_2)
		return;
	}
};


