package main

import "fmt"

func main() {
	fmt.Println("different data types in goloang")
	var a [3]int
	a[1] = 10
	println(a[0])
	println(a[1])
	println(a[2])
	println(a[len(a) -1])
}