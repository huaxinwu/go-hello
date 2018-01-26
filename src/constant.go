package main

import "fmt"
import "unsafe"


// 这种因式分解关键字的写法一般用于声明全局常量
const(
	e11 = "abc"
	f11 = len(e11)
	g11 = unsafe.Sizeof(f11)
)
// iota再次出现时候，值加1,iota默认值为0
const(
	a11 = iota
	b11 = iota
	c11 = iota

)
func main(){

	println("常量声明")
	// 常量声明并赋值
	const LENGTH = 10
	const WIDTH = 5
	var area int
	//  常量多重赋值
	const a,b,c = 1,false,"str"

	area = LENGTH * WIDTH
	fmt.Println("面积为：%d",area)
	println()
	fmt.Println(a,b,c)

	println("全局常量声明")
	println(e11,f11,g11)

	println("常量iota用法")
	println(a11,b11,c11)
	// 常量枚举声明时候，可以省略后面赋值iota
	const(
		e = 100
		f
		g
	)
	println(e,f,g)
}
