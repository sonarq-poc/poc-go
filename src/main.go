package main

import "fmt"

func main() {
    fmt.Println("Go Proof of concept")
}

func addition1() {
	var a int = 1
	var b int = 10
	var c int
	c = a +b
	c += 2
 	c =+ 2

	fmt.Printf("Value of c is %d\n", c)
}

func addition2() {
	var a int = 1
	var b int = 10
	var c int
	c = a +b
	fmt.Printf("Value of c is %d\n", c)
}


