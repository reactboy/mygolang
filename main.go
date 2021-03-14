// mainは作成しているアプリケーションが実行可能であることをGoで定義する方法
// pakcage main は特別
// packageとは一般的なソースコードファイルのセット。
package main

// importステートメントを利用することで別のパッケージ内に配置されている他のコードからプログラムにアクセスできるよになる。
// この場合はmainパッケージからfmtパッケージにアクセスできるようになっている
import "fmt"

// funcステートメントは関数の宣言に使用する予約語
// main関数はプログラムの開始点
// package main 全体で1つのmain関数のみ記述できる
func main() {
	fmt.Println("Hello World!")
}