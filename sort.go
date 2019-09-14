package main

import (
	"sync"
	"math/rand"
	"log"
	"time"
	"sort"
)

func replace(arr *[]float64, i, j int) {
	temp := (*arr)[j]
	(*arr)[j] = (*arr)[i]
	(*arr)[i] = temp
}

func partition(arr *[]float64, start, end, pivot int) int {
	replace(arr, pivot, end)
	endindex := end - 1
	curr := 0
	for curr <= endindex {
		if (*arr)[curr] > (*arr)[end] {
			replace(arr, curr, endindex)
			endindex -= 1
		} else {
			curr += 1
		}
	}
	replace(arr, endindex + 1, end)
	return endindex + 1
}

func parallelSort(arr *[]float64, start, end int, wg *sync.WaitGroup) {
	piv := partition(arr, start, end, rand.Intn(end - start) + start)

	if piv - 1 > start {
		wg.Add(1)
		go parallelSort(arr, start, piv - 1, wg)
	}
	if piv + 1 < end {
		wg.Add(1)
		go parallelSort(arr, piv + 1, end, wg)
	}
	wg.Done()
}

func ParallelSort(arr *[]float64) bool {
	var wg sync.WaitGroup
	go parallelSort(arr, 0, len(*arr)-1, &wg)
	wg.Add(1)
	wg.Wait()
	return true;
}

func sliceEq(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	for i := range(a) {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	arr := make([]float64, 0)
	for i := 0; i<1000; i++ {
		arr = append(arr, rand.Float64())
	}
	darr := make([]float64, len(arr))
	copy(darr, arr)
	prev := time.Now()
	ParallelSort(&arr)
	total := time.Since(prev)
	log.Printf("Parallel Sort: %v\n", total)
	prev = time.Now()
	sort.Float64s(darr)
	total = time.Since(prev)
	log.Printf("Normal Sort: %v\n", total)
	log.Println("Equal?", sliceEq(arr, darr))
	println("done")
	// Parallel-Sort should be slower, cool experinent though
}