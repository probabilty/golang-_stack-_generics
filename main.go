package main

import (
	"fmt"
	"stack/gstack"
)

func main() {
	//creating a new int stack
	s := gstack.New[int](0)
	for i := 1; i < 5; i++ {
		//insterting items to stack
		s.Push(i)
	}
	for i := 0; i < 5; i++ {
		//getting items back from the stack
		fmt.Println(s.Pop())
	}
}
