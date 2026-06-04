/*Copyright (C) 2022 Mandiant, Inc. All Rights Reserved.*/
package main

import "fmt"

type structurea struct {
	test string
}

func add(a, b int) int      { return a + b }
func multiply(a, b int) int { return a * b }

//go:noinline
func neverInlined(x int) int { return x * x }

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	var structa structurea
	structa.test = "hi"

	fmt.Println(structa)

	c := make(chan int)
	s := []int{7, 2, 8, 9}
	go sum(s, c)

	messages := make(chan string)
	go func() { messages <- "ping" }()

	fmt.Println("Hello, this is a test")

	msg := <-messages
	fmt.Println(msg)

	x := <-c
	fmt.Println(x)

	fmt.Println(add(1, 2))
	fmt.Println(multiply(3, 4))
	fmt.Println(neverInlined(5))
}
