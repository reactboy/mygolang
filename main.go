package main

import "fmt"

func main() {
	fmt.Println("different data types in goloang")
	// goの配列は宣言、初期化時にサイズを指定する必要がある。
	// また作成後にサイズを変更することはできない。
	var a [3]int
	a[1] = 10
	println(a[0])
	println(a[1])
	println(a[2])
	println(a[len(a) -1])

	//goでの配列の初期化
	// 要素数5のstringの配列を作成している
	// 初期化時に必ずしも要素数の数ぶん値を設定する必要はない
	cities := [5]string{"New York", "Paris", "Berlin", "Madrid",}
	println(cities)
}