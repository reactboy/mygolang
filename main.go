// mainは作成しているアプリケーションが実行可能であることをGoで定義する方法
// pakcage main は特別
// packageとは一般的なソースコードファイルのセット。
package main

// importステートメントを利用することで別のパッケージ内に配置されている他のコードからプログラムにアクセスできるよになる。
// この場合はmainパッケージからfmtパッケージにアクセスできるようになっている

// funcステートメントは関数の宣言に使用する予約語
// main関数はプログラムの開始点
// package main 全体で1つのmain関数のみ記述できる
func main() {
	// fmt.Println("Hello World!")
	// //変数の宣言は var を使用する
	// // string型の変数firstNameを宣言している
	// var firstName string
	// // 変数の後ろに, で別の変数を定義できる
	// var middleName, lastName string
	// var age int
	// // var キーワードの後ろに()を利用して以下のように記述できる
	// var (
	// 	school, company string
	// 	income int
	// )
	// // 変数の初期化する場合は型を指定しなくてもGoが設定された値から推定する
	// var (
	// 	// 以下のように一行で書くこともできる。
	// 	address, zip = "tokyo", 1500051
	// )
	// // 以下のように := を使用するとvar で宣言せずに変数の初期化ができる。これが最も一般的
	// award := "top sales of the year"
	// price := 10000
	// fmt.Println(firstName, middleName, lastName, age)
	// println(school, company, income)
	// fmt.Println(address,zip)
	// fmt.Println( award, price)
	// // 宣言した変数を一度も使用しなかった場合にGoはエラーになる
	// // const を使用することで定数を宣言できる
	// // 使用しない変数を宣言することはできないが使用しない定数を宣言することはできる
	// const StatusOK = 0
	// // Goを学習するとruneという用語を目にするかもしれないが、runeは単にint32の別名
	// // Unicode文字を表すために使用される
	// rune := 'G'
	// println(rune)
	// // Go では変数の宣言時初期化しなかった場合の規定値が全てのデータ型に存在する
	// // int 0 
	// // float +0.000000e+000
	// // bool false
	// // string 空の値("")

	// // Goでは型のキャストを明示的に行う必要がある
	// var integer16 int16 = 127
	// var integer32 int32 = 32767
	// println(int32(integer16) + integer32)

	// // Go でコマンドライン引数を受け取るには osパッケージとos.Argsを使用する
	// // number1, _ := strconv.Atoi(os.Args[1])
	// // println(number1)
}