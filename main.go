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

	// go は他の言語と違いでifの条件文の()を省略できる
	// go では3項演算子を使えない 毎回完全なif文を書く必要がある
	// x := 27
	// if x%2 == 0 {
	// 	println(x, "is even")
	// }
	// // if文内でのみ使用できる変数を宣言できる。また条件式としても使用することが可能
	// if num := givemenumber(); num < 0 {
	// 	println(num, "is negative")
	// } else if num < 10 {
	// 	println(num, "has only one digit")
	// } else {
	// 	println(num, "has multiple digits")
	// }

	// switch文の基本形
	// sec := time.Now().Unix()
	// rand.Seed(sec)
	// i := rand.Int31n(10)

	// switch i {
	// case 0:
	// 	fmt.Print("zero...")
	// case 1:
	// 	fmt.Print("one...")
	// case 2:
	// 	fmt.Print("tow...")
	// default:
	// 	fmt.Print("no match...")
	// }
	// println("ok")
	// region, continent := location("Irvine")

	// println("I work in ", region, continent)
	// go でのfor文の基本形
	// sum := 0
	// for i := 1; i <= 100; i++ {
	// 	sum += i
	// }
	// fmt.Println("sum of 1...100 is", sum)

	// var num int64
	// rand.Seed(time.Now().Unix())
	// // go には他の言語でいうwhile文と同等のものがない.forを使って同じことができる
	// for num != 5 {
	// 	num = rand.Int63n(15)
	// 	fmt.Println(num)
	// }

	// for i := 1; i <= 4; i++ {
	// 	// go ではdeferを使用することで実行を遅延できる.使い所がいまいちわからない
	// 	// 最後に実行させたい場合に使うのかな
	// 	defer fmt.Println("deferred", -i)
	// 	fmt.Println("regular", i)
	// }
	// defer func() {
	// 	// recoverはdeferされた関数で呼び出すことができるもの
	// 	// プログラムがpanicになっていない場合はnilを返すがpanicになっている場合はnilを返さない
	// 	// panicになった際に処理をしたい場合に使う?
	// 	if r := recover(); r != nil {
	// 		println("Recoverd in main", r)
	// 	}
	// }()
	// panicloop()
	// fizzbuzz()
	// // guessroot(25)
	// sqrt(25)
	// panicnegative()
}

// func panicloop() {

// 	panicCount := 0
// 	for panicCount < 10 {
// 		if panicCount > 4 {
// 			println("panicking!")
// 			// go ではpanicを使用すると処理を止めれる
// 			panic("Panic in for loop")
// 		}
// 		// panicで処理が止まったとしても処理が中断される前にdifferで遅延させた処理が実行される
// 		defer println("defer in for loop", panicCount)
// 		println("printing in for loop", panicCount)
// 		panicCount = panicCount + 1
// 	}
// }

// func givemenumber() int {
// 	return -1
// }

// func fizzbuzz() {
// 	count := 0
// 	for count <= 100 {
// 		switch {
// 		case count%3 == 0 && count%5 == 0:
// 			println(count, "FizzBuzz")
// 		case count%3 == 0:
// 			println(count, "Fizz")
// 		case count%5 == 0:
// 			println(count, "Buzz")
// 		}
// 		count++
// 	}
// }

// func guessroot (x int) {
// sroot := 1
// res := sroot - (sroot -x) / (2 * sroot)
// println("first res", res)
// sroot = res
//  for res == sroot {
// 	println("A guess for square root is ", sroot);
// 	res = sroot - (sroot -x) / (2 * sroot)
// 	println("for res", res)
// 	sroot = res
// }
// println(res == sroot)
// println("Square root is:", sroot);
// }

// func sqrt(num float64) float64 {
// 	currguess := 1.0
// 	prevguess := 1.0
// 	for count := 1; count <= 10; count++ {
// 		prevguess = currguess
// 		currguess = prevguess - (prevguess*prevguess-num)/(2*prevguess)
// 		if currguess == prevguess {
// 			break
// 		}
// 		fmt.Println("A guess for square root is", currguess)
// 	}
// 	return currguess
// }

// func panicnegative() {
// 	val := 0
// 	for {
// 		fmt.Println("Enter number: ")
// 		fmt.Scanf("%d", &val)
// 		switch {

// 		case val < 0:
// 			panic("You entered a negative number!")
// 		case val == 0:
// 			fmt.Println("0 is neither negative nor positive")
// 		default:
// 			fmt.Println("You entered:", val)

// 		}
// 	}
// }

// func location(city string) (string, string) {
// 	var region string
// 	var continent string
// 	switch city {
// 	// いずれかに一致するのような書き方も可能
// 	case "Delhi", "Hyderbad", "Mumbai":
// 		region, continent = "India", "Asia"
// 	case "Irvine", "Los Angeles", "San Diego":
// 		region, continent = "Calfornia", "USA"
// 	default:
// 		region, continent = "Unknown", "Unknown"
// 	}

// 	return region, continent

// }
