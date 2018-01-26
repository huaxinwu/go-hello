package main
// 结构体
import "fmt"

/*  定义结构体   */
type Books struct {
	title string
	author string
	subject string
	book_id int
}


func main()  {
	/* 声明 Book1 为 Books 类型 */
	var Book1 Books

	/* Book1 描述 */
	Book1.title = "Go语言"
	Book1.author = "www.runbo.com"
	Book1.subject = "Go语言教程"
	Book1.book_id = 645512

	/* 打印 Book1 信息 */
	fmt.Printf( "Book1 title : %s\n", Book1.title)
	fmt.Printf( "Book1 author : %s\n", Book1.author)
	fmt.Printf( "Book1 subject : %s\n", Book1.subject)
	fmt.Printf( "Book1 book_id : %d\n", Book1.book_id)

	fmt.Println("---------------------------------")
	/* 打印 Book1 信息 */
	printBook(Book1)
	fmt.Println("*****************************************")
	/*  打印 Book1 信息,指针变量  */
	printBookPointer(&Book1)
}
/*
	打印books信息,对象变量
 */
func printBook(book Books){
	fmt.Printf("Book1 title : %s\n",book.title)
	fmt.Printf( "Book1 author : %s\n", book.author)
	fmt.Printf( "Book1 subject : %s\n", book.subject)
	fmt.Printf( "Book1 book_id : %d\n", book.book_id)
}
/*
	打印books信息,指针变量
 */
func printBookPointer(book *Books){
	fmt.Printf("Book1 title : %s\n",book.title)
	fmt.Printf( "Book1 author : %s\n", book.author)
	fmt.Printf( "Book1 subject : %s\n", book.subject)
	fmt.Printf( "Book1 book_id : %d\n", book.book_id)
}