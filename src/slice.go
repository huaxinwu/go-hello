package main
// 切片
import "fmt"

func main(){
	/*    定义切片并初始化，指定类型、长度、容量       */
	var numbers = make([]int,3,5)

	/*  打印切片   */
	printSlice(numbers)

	fmt.Println("---------------切片是空判断------------------------")
	var numbers1 []int
	if(numbers1 == nil){
		fmt.Println("切片是空的")
	}
	fmt.Println("---------------切片追加和拷贝------------------------")
	/* 允许追加空切片 */
	numbers1 = append(numbers1, 0)
	printSlice(numbers1)

	/* 向切片添加一个元素 */
	numbers1 = append(numbers1, 1)
	printSlice(numbers1)
	/* 同时添加多个元素 */
	numbers1 = append(numbers1, 2,3,4)
	printSlice(numbers1)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers3 := make([]int, len(numbers1), (cap(numbers1))*2)
	/* 拷贝 numbers1 的内容到 numbers3 */
	copy(numbers3,numbers1)
	printSlice(numbers3)

	fmt.Println("--------------切片截取-------------------------------")

	numbers2 := []int{0,1,2,3,4,5,6,7,8}
	/* 打印原始切片 */
	fmt.Println("numbers2 ==", numbers2)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers2[1:4] ==", numbers2[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers2[:3] ==", numbers2[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers2[4:] ==", numbers2[4:])

}

/*  打印切片   */
func printSlice(x []int)  {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}