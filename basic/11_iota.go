package main

import "fmt"

func main() {
	// iota：常量生成器，每隔一行自动累加1， 常用于给常量复制
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a = %d, b = %d, c= %d\n", a, b, c)
	
	// iota 遇到const，重置为0
	const d = iota
	fmt.Printf("d = %d", d)
	
	// 可以只写一个iota
	const (
		a1 = iota
		b1  
		c1 
	)
	fmt.Printf("a1 = %d, b1 = %d, c1 = %d\n", a1, b1, c1)
	
	//如果同一行，值一样
	const (
		i		 = iota
		j1,j2,j3 = iota, iota, iota
		k		 = iota
	)
	fmt.Printf("i=%d, j1=%d,j2=%d,j3=%d,k=%d\n", i, j1, j2, j3, k)	
}