package main

import "fmt"

// defer ステートメントは、 defer へ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させるものです。
// defer へ渡した関数の引数は、すぐに評価されますが、その関数自体は呼び出し元の関数がreturnするまで実行されません。
func Defer() {
	defer fmt.Println(test_func())
	fmt.Println("hello")
}

func test_func() string {
	fmt.Println("hello f()")
	return "end"
}


// defer へ渡した関数が複数ある場合、その呼び出しはスタック( stack )されます。
// 呼び出し元の関数がreturnするとき、 defer へ渡した関数は LIFO(last-in-first-out) の順番で実行されます。
func stacking_defers() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}