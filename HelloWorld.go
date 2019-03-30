package main

import (
	"fmt"
)
	
func main(){
	fmt.Println("Hello World")
	fmt.Println(sum(3, 4))
}

func sum(x int, y int) int{
	return x + y
}
