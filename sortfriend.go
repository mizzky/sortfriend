package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("insertsort: [%fsec]\n", sortShot(insersionSort, 100, 1000))
}

func sortShot(f func([]int) []int, shotNum, size int) float64 {
	n := 0
	start := time.Now()
	for n < shotNum {
		slice := generateSlice(size)
		_ = f(slice)
		n++
	}
	end := time.Now()
	return end.Sub(start).Seconds()
}

func generateSlice(size int) []int {
	s := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		s[i] = rand.Intn(999) - rand.Intn(999)
	}
	return s
}

func insersionSort(list []int) []int {
	var n = len(list)
	for i := 1; i < n; i++ {
		tmp := list[i]
		if list[i-1] > tmp {
			j := i
			for j > 0 && list[j-1] > tmp {
				list[j] = list[j-1]
				j--
			}
			list[j] = tmp
		}
	}
	return list
}
