package main

import "fmt"

func main() {
	i := 0;
	// defer 语句，直接执行
	defer fmt.Println("defer print i:", i);
	// defer 匿名函数，也就是闭包，其变量继承调用函数的作用域
	defer func(){fmt.Println("defer in a func: ", i)}();
	i++;
	fmt.Println("sequence i:", i);
}


