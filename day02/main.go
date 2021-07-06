package main

import (
	"fmt"
)

func a(b bool) {
	fmt.Println(b)
}

func main() {
	a("1" > "2")
}
