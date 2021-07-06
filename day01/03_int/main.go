package main
import "fmt"

type move1 interface {
	move()
}

type car struct {
	name string
}

func (c *car) move() {
	fmt.Println("move")
} 

func main()  {
	// var a = new(int)
	// fmt.Println(a)
	// b := make(map[string]int, 0)
	// fmt.Printf("cap: %v\n", len(b))
	// b["11"] = 1
	// b["22"] = 2
	// fmt.Printf("len: %d\n", len(b))
	var a1 move1
	a1 = &car{name: "im is car"}
	a1.move()
	fmt.Println(a1)

}