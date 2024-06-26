package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func ceil(x float64) int {
	return int(math.Ceil(x))
}

func sortFunc(wg *sync.WaitGroup, seq []int, ch chan<- []int) {
	defer wg.Done()
	fmt.Println("before sorting subarray:", seq)
	sort.Ints(seq)
	ch <- seq
}

func mergeArrays(arrays [][]int) []int {
	merged := arrays[0]
	for i := 1; i < len(arrays); i++ {
		merged = mergeTwoArrays(merged, arrays[i])
	}
	return merged
}

func mergeTwoArrays(a, b []int) []int {
	result := make([]int, len(a)+len(b))
	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result[k] = a[i]
			i++
		} else {
			result[k] = b[j]
			j++
		}
		k++
	}
	for i < len(a) {
		result[k] = a[i]
		i++
		k++
	}
	for j < len(b) {
		result[k] = b[j]
		j++
		k++
	}
	return result
}

func main() {
	fmt.Println("Enter series of integers:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	inputArr := strings.Fields(input)

	intArr := make([]int, len(inputArr))
	for i, indexVal := range inputArr {
		intArr[i], _ = strconv.Atoi(indexVal)
	}

	portionSize := ceil(float64(len(intArr)) / 4.0)
	fmt.Println("Portion size:", portionSize)

	var wg sync.WaitGroup
	ch := make(chan []int, 4)

	for i := 0; i < 4; i++ {
		start := i * portionSize
		end := start + portionSize
		if end > len(intArr) {
			end = len(intArr)
		}
		wg.Add(1)
		go sortFunc(&wg, intArr[start:end], ch)
	}

	wg.Wait()
	close(ch)

	var sortedSubarrays [][]int
	for eachPortion := range ch {
		sortedSubarrays = append(sortedSubarrays, eachPortion)
	}

	finalSlice := mergeArrays(sortedSubarrays)
	fmt.Println("Sorted array:", finalSlice)
}
