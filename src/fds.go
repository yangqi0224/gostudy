package main

import "fmt"

type annimal interface {
	eat()
	say()
}

type dog struct {
	annimal
}

type cat struct {
	annimal
}

func (d dog) eat() {
	fmt.Println("dog eat shit")
}
func (c cat) eat() {
	fmt.Println("cat eat fish")
}

func (d dog) say() {
	fmt.Println("wang wang")
}
func (c cat) say() {
	fmt.Println("meow")
}

const (
	l = iota
	p
	e
)

func main() {
	var b, c int
	b, c = test(2, 3)
	fmt.Println(l, p, e)
	fmt.Printf("b value is %d, c value is %d \n", b, c)

	dd, cc := dog{}, cat{}
	dd.say()
	dd.eat()
	cc.say()
	cc.eat()
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
