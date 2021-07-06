package main
import "fmt"
//常量声名
const pi = 3.1415926

//常量批量声明
const (
	a = 100
	b
	c
)

//iota
//iota是go语言的常量计数器，只能在常量的表达式中使用。
//iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。
const (
	n1 = iota
	n2
	n3
	n4
)

func main()  {
	fmt.Println(n1, n2, n3, n4)
}