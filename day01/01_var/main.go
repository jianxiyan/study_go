package main
import "fmt"
//变量声明
// var name string
// var age int
// var isok bool

//批量声明
var (
	name string //""
	age int // 0
	isok bool // false
)

func main()  {
	name = "小米"
	age = 16
	isok = true
	//go语言中变量声明后必须得使用
	fmt.Println(name, age, isok)
}