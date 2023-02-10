package main

import "fmt"

var x = 10

func main() {
    fmt.Println("Go Proof of concept")
	fmt.Println("Value of x:", x)
		foo() 
	fmt.Println("Value of x:", x)
	foo2()
	foo2()

	fmt.Println("Value of x:", x)
	fmt.Println("Value of x:", x)
}

	
function doNothing()
{
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

func foo2() {
		x =x-5
}

func foo() {
	x = x + 2
}

