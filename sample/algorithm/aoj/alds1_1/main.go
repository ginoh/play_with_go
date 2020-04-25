package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	values, err := input()
	if err != nil {
		log.Fatal(err)
	}

	// algo1
	// for j が 1 〜 n-1 まで
	// 	for i が 0 〜 j-1 まで
	// 		maxv = (maxv と R[j]-R[i]のうち大きい方)
	// maxv 初期値は 10^9 * (-1)以下。もしくは R1-R0
	// o(n^2)
	// n := len(values)
	// y, _ := strconv.Atoi(values[1])
	// x, _ := strconv.Atoi(values[0])

	// maxv := y - x
	// for j := 1; j < n-1; j++ {
	// 	for i := 0; i < j-1; i++ {
	// 		y, _ := strconv.Atoi(values[j])
	// 		x, _ := strconv.Atoi(values[i])
	// 		tmp := y - x

	// 		if maxv < tmp {
	// 			maxv = tmp
	// 		}
	// 	}
	// }
	// fmt.Printf("max value is %d\n", maxv)

	// algo2
	//minv = R[0]
	//for j が 1〜n-1 まで
	//	maxv = (maxvとR[j]-minvのうち大きい方)
	//	minv = (minvとR[j]のうち小さい方)
	// o(n)
	n := len(values)
	y, _ := strconv.Atoi(values[1])
	x, _ := strconv.Atoi(values[0])
	minv := x
	maxv := y - x
	for j := 1; j < n-1; j++ {
		z, _ := strconv.Atoi(values[j])
		if maxv < z-minv {
			maxv = z - minv
		}
		if minv > z {
			minv = z
		}
	}
	fmt.Printf("max value is %d\n", maxv)
}

func input() ([]string, error) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	input := string(bytes)
	lines := strings.Split(input, "\n")
	n, err := strconv.Atoi(lines[0])
	if err != nil {
		return nil, err
	}
	_ = n
	return lines[1 : 1+n], nil
}
