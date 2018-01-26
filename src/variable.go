package main


// 没有指定类型，go编译器会根据值自己指定类型，duck typing
var a = "菜鸟教程"
// 指定变量的类型
var b string = "runnbo.com"
// 没有给变量赋值，go编译器会自动采用该类型的默认值
var c bool


var x,y int
// 这种因式分解关键字的写法一般用于声明全局变量
var(
	d int
	e bool
)

var f,g int = 1,2

// 这种不带声明格式的只能在函数体中出现
// h,i=1,2

// 全局变量允许声明后不使用
var l,m int = 4,5

func main(){
	println("变量声明")
	println(a,b,c)

	println("多变量声明")
	// 局部变量声明方式并赋值
	h,i:=1,2
	println(x,y,d,e,f,g,h,i)
	// 局部变量声明并初始化，但是没有使用编译不会通过
	//var k string = "abc"
}

