package main

import "fmt"

func Print(str string, ch chan int) {
	fmt.Println(str)
	ch <- 1
}

func main() {
	fmt.Println("hello Urvashi")
	fmt.Println("hello Ashish")
	ch := make(chan int)
	go Print("Ashish", ch)
	<-ch
}
