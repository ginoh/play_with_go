package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// // for i in {0..10}; do echo -n "$(($RANDOM % 100)), "; done

	fmt.Println("=== Slice のソート (パッケージ) ===")

	fmt.Println("=== 独自のバブルソート ===")
	list1 := []int{27, 16, 128, 41, 167, 87, 147, 73, 98, 127, 11}
	fmt.Printf("before sort: %v\n", list1)
	myBubbleSort2(list1)
	fmt.Printf("after  sort: %v\n", list1)

	fmt.Println("=== 独自のセレクションソート ===")
	list2 := []int{73, 30, 29, 64, 19, 27, 14, 31, 85, 18, 76}
	fmt.Printf("before sort: %v\n", list2)
	mySelectionSort(list2)
	fmt.Printf("after  sort: %v\n", list2)

	fmt.Println("=== 独自の挿入ソート ===")
	list3 := []int{98, 45, 6, 50, 92, 4, 37, 41, 89, 99, 84}
	fmt.Printf("before sort: %v\n", list3)
	myInsersionSort(list3)
	fmt.Printf("after  sort: %v\n", list3)

	fmt.Println("=== 独自のクイックソート ===")
	list4 := []int{4, 0, 59, 37, 36, 96, 51, 6, 16, 54, 34}
	fmt.Printf("(sort1) before sort: %v\n", list4)
	quickSortedList := myQuickSort(list4)
	fmt.Printf("(sort1) after  sort: %v\n", quickSortedList)

	list5 := []int{4, 0, 59, 37, 36, 96, 51, 6, 16, 54, 34}
	//list5 := []int{4, 0, 59, 36, 36, 96, 2, 6, 34}
	//list5 := []int{4, 81, 59, 48, 34, 34, 65, 6, 34}
	//list5 := []int{4, 81, 59}
	fmt.Printf("(sort2) before sort: %v\n", list5)
	myQuickSort2(list5, 0, len(list5)-1)
	fmt.Printf("(sort2) after  sort: %v\n", list5)

	list6 := []int{4, 0, 59, 37, 36, 96, 51, 6, 16, 54, 34}
	//list6 := []int{4, 0, 59, 36, 36, 96, 2, 6, 34}
	//list6 := []int{4, 81, 59, 48, 34, 34, 65, 6, 34}
	//list6 := []int{4, 81, 59}
	fmt.Printf("(sort3) before sort: %v\n", list6)
	myQuickSort3(list6)
	fmt.Printf("(sort3) after  sort: %v\n", list6)
}

func myBubbleSort(arr []int) {
	// len(arr) - 1 回繰り返す、最後の１個以外の順番が確定しているなら最後の1個も確定する
	for i := 0; i < len(arr)-1; i++ {
		// 確定していない末尾要素を確定するように比較の繰り返し
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// bubble sort って先頭要素を確定して行くのが元々だっけ？
func myBubbleSort2(arr []int) {
	// len(arr) - 1 回繰り返す、最後の１個以外の順番が確定しているなら最後の1個も確定する
	for i := 0; i < len(arr)-1; i++ {
		// 確定していない先頭要素を確定するように比較の繰り返し
		for j := len(arr) - 1; i < j; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

func mySelectionSort(arr []int) {
	// len(arr) - 1 回繰り返す、最後の１個以外の順番が確定しているなら最後の1個も確定する
	for i := 0; i < len(arr)-1; i++ {
		// 最小値のindex
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func myInsersionSort(arr []int) {
	// len(arr) - 1 回繰り返す, 先頭はソート済み扱い
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}

// myHeapsort
// 実装するのは大変そうなので、golang の heap を使ったソートを試す

// myMergeSort

// 愚直に pivot と比較して小さいグループ、大きいグループのスライスを作成したもの
// 新規スライスが作成されるのでスライスを戻り値としている
func myQuickSort(arr []int) []int {
	var rightSide, leftSide []int

	if len(arr) < 2 {
		return arr
	}
	// pivot とする値をランダムに選んでみる
	rand.Seed(time.Now().UnixNano())
	pivotIdx := rand.Intn(len(arr))

	// pivot を端の要素 (e.g. 先頭)とすれば
	// 条件分岐で skip しなくてもループ処理の開始や終了を調整することでもよさそう
	for i := 0; i < len(arr); i++ {
		if i == pivotIdx {
			continue
		}
		if arr[pivotIdx] < arr[i] {
			rightSide = append(rightSide, arr[i])
		} else {
			// 今回は等しい場合も leftSide にいれる
			leftSide = append(leftSide, arr[i])
		}
	}
	return append(append(myQuickSort(leftSide), arr[pivotIdx]), myQuickSort(rightSide)...)
}

// 一般的によくあるクイックソートの実装
func myQuickSort2(arr []int, left, right int) {
	if left >= right {
		return
	}

	// // やり方1
	i := left
	j := right
	// 真ん中を基準に設定
	pivot := arr[(i+j)/2]

	for true {
		// 先頭から pivot 以上の値を探す
		for arr[i] < pivot {
			i++
		}
		// 末尾から pivot 以下の値を探す
		for pivot < arr[j] {
			j--
		}
		// index が同一または逆転した時点で終了
		// 同一のindexを指すのは、
		// left, right の両方の交換対象がないまま、探査が pivot と同一の値 (同じ値が存在する場合もあるのでこの書き方)に到達した時
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]

		// 交換した分、index を進める
		// pivot と同値の物がある場合、この処理をしないとずっとswap しつづけることになりえる
		i++
		j--
	}

	// // やり方2
	// // pivot を末尾の値とする方法
	// i := left
	// j := right - 1
	// pivot := arr[right]

	// for true {
	// 	// 先頭から pivot 以上の値を探す
	// 	// 存在しない場合は pivot で止まる
	// 	for arr[i] < pivot {
	// 		i++
	// 	}
	// 	// 末尾から pivot より小さな値を探す
	// 	// "=" の条件がないと同値のもので swap しつづけることになりえる
	// 	// また、left側の index と交差しない範囲で探す
	// 	for pivot <= arr[j] && j > i {
	// 		j--
	// 	}

	// 	if i >= j {
	// 		// 探索終了時点の index の値が pivotより大きければ交換
	// 		if arr[i] > pivot {
	// 			arr[i], arr[right] = arr[right], arr[i]
	// 		}
	// 		break
	// 	}
	// 	fmt.Printf("swap arr[i]=%d, arr[j]=%d\n", arr[i], arr[j])
	// 	arr[i], arr[j] = arr[j], arr[i]
	// }

	// fmt.Printf("pivot=%d\n", pivot)
	// fmt.Printf("arr=%#v\n", arr)
	// fmt.Printf("myQuickSort2(%d, %d)\n", left, i-1)
	// fmt.Printf("myQuickSort2(%d, %d)\n", j+1, right)
	myQuickSort2(arr, left, i-1)
	myQuickSort2(arr, j+1, right)
}

// 再帰処理を利用しないクイックソート
// 参考
// https://programming-place.net/ppp/contents/algorithm/sort/006.html#quick_sort
// https://cattech-lab.com/science-tools/simulation-lecture-mini-4/
//
// 再帰で問題になるのはスタックオーバーフロー。なのでスタックを自前で管理するというのがこの実装
// 再帰処理で行っていたことをループで実装する、次の処理対象をスタックで管理しておくという方針
func myQuickSort3(arr []int) {
	// 分割範囲の下限・上限の２値の組み合わせをスタックとして管理する
	var stack [][2]int

	// TBD: 必要なスタックサイズのみ確保する

	// 処理対象の初期値を stack に積む
	stack = append(stack, [2]int{0, len(arr) - 1})

	// 処理対象が stack に残っている限り処理を続行
	for len(stack) > 0 {
		//fmt.Printf("stack=%#v\n", stack)
		// pop 処理を関数化したほうが見通しはよさそう
		left, right := stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]

		// 処理対象の範囲が逆転していたら何もしない
		if left >= right {
			continue
		}

		i := left
		j := right
		pivot := arr[(i+j)/2]

		// 分割処理のここは関数化してもよさそう
		for true {
			// 先頭から pivot 以上の値を探す
			for arr[i] < pivot {
				i++
			}
			// 末尾から pivot 以下の値を探す
			for pivot < arr[j] {
				j--
			}
			// index が同一または逆転した時点で終了
			if i >= j {
				break
			}
			arr[i], arr[j] = arr[j], arr[i]

			// 交換した分、index を進める
			i++
			j--
		}

		// stack へ push
		// left, right で要素数の多い方を先に stack に積んでおく
		// これは少ない方を先に処理するため
		// i-left => 左側の要素数、right-j => 右側の要素数
		if i-left < right-j {
			stack = append(stack, [2]int{j + 1, right}, [2]int{left, i - 1})
		} else {
			stack = append(stack, [2]int{left, i - 1}, [2]int{j + 1, right})
		}
	}
}
