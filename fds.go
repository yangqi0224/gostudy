package main

import "fmt"

func main() {
	var b, c int
	b, c = test(2, 3)
	fmt.Printf("b value is %d, c value is %d", b, c)
}

func test(a int, b int) (int, int) {
	//可同时声明并给多个变量复制
	var c = a + b
	//变量要么用var声明，要么用运算符 ':=' 直接赋值
	d := 20
	if c < 10 {
		c = d
	}
	return c, a
}
