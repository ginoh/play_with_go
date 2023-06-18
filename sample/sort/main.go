package main

import (
	"fmt"
	"sort"
)

// sort package
// https://pkg.go.dev/sort
func main() {
	// // for i in {0..10}; do echo -n "$(($RANDOM % 100)), "; done

	fmt.Println("=== Slice のソート (パッケージ) ===")
	list1 := []int{60, 75, 8, 86, 52, 85, 37, 26, 98, 45, 30}
	fmt.Printf("before sort: %v\n", list1)

	// Slice() の第２引数にi < j になる条件を判定する関数を指定する、関数の内容で昇順・降順を制御
	// 安定したソートではないので等しい内容の順序が逆になる可能性がある
	// 安定したソートの場合は SliceStable を使う
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	fmt.Printf("after  sort: %v\n", list1)
}
