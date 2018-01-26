package main
// 指针
import "fmt"

func main()  {
	var a int = 10
	/*  表示一个指针    */
	var ip *int
	/* 指针变量的存储地址 */
	ip = &a
	/*  十六进制输出 */
	fmt.Printf("a变量的存储地址：%x\n",&a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )

	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr  )

}