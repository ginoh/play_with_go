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
	fmt.Printf("before sort: %v\n", list4)
	quickSortedList := myQuickSort(list4)
	fmt.Printf("after  sort: %v\n", quickSortedList)
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

func myQuickSort(arr []int) []int {
	var leftSide, rightSide []int

	if len(arr) < 2 {
		return arr
	}
	rand.Seed(time.Now().UnixNano())
	pivotIdx := rand.Intn(len(arr))

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
