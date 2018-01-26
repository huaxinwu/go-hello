package main
// 语句：条件语句、循环语句
import "fmt"

func main(){
	fmt.Println("select语句")

	var c1,c2,c3 chan int
	var i1,i2 int

	select {
	case i1 = <-c1:
		fmt.Println("received",i1,"from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")

	}

	fmt.Println("循环语句")
	var b int = 15
	var a int
	numbers :=[6] int{1,2,3,5}
	for a := 0; a < 10;a++ {
		fmt.Printf("a的值：%d\n",a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}
	for i,x := range  numbers {
		fmt.Printf("第%d位x的值：%d\n",i,x)
	}

}
