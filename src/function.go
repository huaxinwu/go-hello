package main
// 函数

import (
	"fmt"
	"math"
)

func main(){
	/*  定义局部变量  */
	var a int = 100
	var b int = 200
	var ret int

	/*  调用函数 */
	ret = max(a,b)
	fmt.Println("两个数中的最大值：%d\n",ret)
	/*  调用函数 */
	c,d := swap("Mahesh","Kumar")
	fmt.Println(c,d)

	fmt.Println("函数中的值传递")
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	/* 调用函数，函数作为值传递 */
	fmt.Println(getSquareRoot(9))

}

/*  定义函数,找出两个数中的最大值  */
func max(num1 int,num2 int) int{
	/*  定义局部变量  */
	var result int
	if(num1 > num2){
		result = num1
	}else{
		result = num2
	}
	return  result
}

/*   定义函数，交换两个字符串的位置           */
func swap(x string,y string)  (string,string) {
	return y,x
}

/*   函数作为一个闭包来使用    */
func getSquare() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
