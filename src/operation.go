package main

// 运算符

import "fmt"

func main(){
	var a int = 12;
	var b int = 10
	var c int

	fmt.Println("算术运算")
	c = a + b
	fmt.Println("a+b的值：%d\n",c)
	c = a - b
	fmt.Println("a-b的值：%d\n",c)
	c = a * b
	fmt.Println("a*b的值：%d\n",c)
	c = a / b
	fmt.Println("a/b的值：%d\n",c)
	c = a % b
	fmt.Println("a%b的值：%d\n",c)
	a++
	fmt.Println("a的值：%d\n",a)

	fmt.Println("关系运算")
	if(a > b){
		fmt.Println("a大于b")
	}else {
		fmt.Println("a小于b")
	}

	fmt.Println("逻辑运算")
	var d bool = true
	var f bool = false
	if(d && f){
		fmt.Println("结果为true\n")
	}else{
		fmt.Println("结果为false\n")
	}

	fmt.Println("指针运算")
	var ptr *int
	ptr = &a
	fmt.Println("ptr的值：%d\n",*ptr)
}
