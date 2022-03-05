package main

import (
	"fmt"
	"time"
)

func main() {

	/// フィールドを3つ持つ構造体
	type point struct {
		ID   string
		x, y int
	}
	p := point{
		ID: "Bob",
		x:  1,
		y:  2,
	}
	fmt.Println(p)

	/// ゼロ値で初期化
	var array [5]int
	fmt.Println("array length:", len(array))

	/// 5つの要素を持つ配列として定義
	arrayLiteral := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arrayLiteral length:", len(arrayLiteral))

	/// 要素数から配列数を推論。この場合は5
	arrayInterference := [...]int{1, 2, 3, 4, 5}
	fmt.Println("arrayInterference length:", len(arrayInterference))

	/// 配列のインデックスと値を指定
	/// indexの指定がないものはゼロ値
	arrayIndex := [...]int{1: 2, 2: 1, 5: 5, 7: 13}
	fmt.Println("arryaIndex:", arrayIndex)
	fmt.Println("arrayIndex length:", len(arrayIndex))

	var slice []int
	fmt.Println("slice: ", slice)

	/// 5つの要素を持つsliceとして定義
	sliceLiteral := []int{1, 2, 3, 4, 5}
	fmt.Println("sliceLiteral: ", sliceLiteral)

	var m map[string]int
	fmt.Println(m)
	/// 2つの要素を持つマップを定義
	mapLiteral := map[string]int{
		"John":    42,
		"Richard": 33,
	}
	fmt.Println(mapLiteral)

	/// 新しい型MyDureationをtime.Durationを基底として定義
	type MyDuration time.Duration
	d := MyDuration(100)

	/// %Tを使うことで変数に代入されている値の型情報を出力する。
	fmt.Printf("%T\n", d)

	/// MyDuration型で規定型として定義しているtime.Durationへのキャスト
	td := time.Duration(d)
	md := 100 * d

	fmt.Println("md:", md)
	fmt.Printf("td: %T, md: %T\n", td, md)

	src := []int{1, 2, 3, 4}
	fmt.Println(src, len(src), cap(src))
	/// スライスに構造体を追加した際にcap(容量)は倍になる。そのため今回は4の容量に対して追加して8となった。
	src = append(src, 5)
	fmt.Println(src, len(src), cap(src))

}
