package main

import "fmt"

func main() {
	name := []int{1, 2, 3, 4, 5}
	sp := name[1:]
	var pp = make([]int, 10)

	//切片拷贝
	copy(pp, sp)
	//切片长度
	fmt.Println(len(pp))

	//切片容量
	fmt.Println(cap(sp))

	//切片元素
	fmt.Printf("the slice address is %p \n", pp)
	changeSplice(sp)
	fmt.Println(sp)
	fmt.Println(pp)
}

func changeSplice(a []int) {

	fmt.Printf("the slice address is %p \n", a)
	b := append(a, 1, 2, 2, 3, 4)

	point := &a
	*point = b
	fmt.Println(a)
}
