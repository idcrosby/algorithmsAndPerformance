package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sort"
	"time"
	)

func main() {

	var result = make([]int, 10)
	for i, _ := range result {
		result[i] = rand.Intn(102)
	}
	input := result //generateRandomArray(10)
	fmt.Println("Input: " + arrayToString(input))
	fmt.Println("closestPair: " + arrayToString(ClosestPair(input)))
	test(generateRandomArray(100000))
}

func test(input []int) {
	// fmt.Println("Input: " + arrayToString(input))
	inputCopy := make([]int, len(input))
	copy(inputCopy, input)
	start := time.Now()
	output := BubbleSort(inputCopy)
	elapsed := time.Since(start)
	if isSorted(output) && arraysEqual(output, input) {
		fmt.Printf("Bubble sort passed. Took: %s\n", elapsed)
	} else {
		// fmt.Println("Bubble sort failed: " + arrayToString(output))
	}

	inputCopy = make([]int, len(input))
	copy(inputCopy, input)
	start = time.Now()
	output = InsertionSort(inputCopy)
	elapsed = time.Since(start)
	if isSorted(output) && arraysEqual(output, input) {
		fmt.Printf("Insertion sort passed. Took: %s\n", elapsed)
	} else {
		fmt.Println("Insertion sort failed: " + arrayToString(output))
	}

	inputCopy = make([]int, len(input))
	copy(inputCopy, input)
	start = time.Now()
	output = SelectionSort(inputCopy)
	elapsed = time.Since(start)
	if isSorted(output) && arraysEqual(output, input) {
		fmt.Printf("Selection sort passed. Took: %s\n", elapsed)
	} else {
		fmt.Println("Selection sort failed: " + arrayToString(output))
	}

	inputCopy = make([]int, len(input))
	copy(inputCopy, input)
	start = time.Now()
	output = MergeSort(inputCopy)
	elapsed = time.Since(start)
	if isSorted(output) && arraysEqual(output, input) {
		fmt.Printf("Merge sort passed. Took: %s\n", elapsed)
	} else {
		fmt.Println("Merge sort failed: " + arrayToString(output))
	}

	inputCopy = make([]int, len(input))
	copy(inputCopy, input)
	start = time.Now()
	output = ComboMergeInsertionSort(inputCopy)
	elapsed = time.Since(start)
	if isSorted(output) && arraysEqual(output, input) {
		fmt.Printf("Combo sort passed. Took: %s\n", elapsed)
	} else {
		fmt.Println("Combo sort failed: " + arrayToString(output))
	}
	// fmt.Println("Output: " + arrayToString(output))

	inputCopy = make([]int, len(input))
	copy(inputCopy, input)
	start = time.Now()
	sort.Ints(inputCopy)
	elapsed = time.Since(start)
	if isSorted(inputCopy) && arraysEqual(inputCopy, input) {
		fmt.Printf("Built-in sort passed. Took: %s\n", elapsed)
	} else {
		fmt.Println("Built-in sort failed: " + arrayToString(inputCopy))
	}

	// TestForTrend(MergeSort, 100000, 1000)
}

func TestForTrend(fn sortFunc, maxSize, step int) {

	// var output int[]
	for i := step; i < maxSize; i = i+step {
		var avgTime time.Duration
		input := generateRandomArray(i)
		var totalTime time.Duration;
		for j := 0; j < 10; j++ {
			start := time.Now()
			output := fn(input)
			elapsed := time.Since(start)
			totalTime += elapsed
			if !isSorted(output) {
				fmt.Println("Error. Array sorting failed")
				return
			}
		}
		avgTime = totalTime / 10;
		// avgTime = int(totalTime.Nanoseconds() / int64(10)) * time.Nanosecond
		fmt.Printf("Average Time for input size %d is %s\n", i, avgTime)
	}
}

// Type to represent generic sort function
type sortFunc func([]int) []int

func BubbleSort(input []int) []int {

	var done bool

	for !done {
		done = true
		for i, val := range input {
			if i < (len(input) - 1) && val > input[i+1] {
				//swap
				input[i] = input[i+1]
				input[i+1] = val
				done = false
			}
		}
	}
	return input
}

func SelectionSort(input []int) []int {
	var swap bool
	var newMin int

	for pointer := 0; pointer < len(input); pointer++ {
		swap = false
		min := pointer
		for i, el := range input {

			if i > pointer && el < input[min] {
				min = i
				swap = true
			}
		}
		if swap {
			// swap
			newMin = input[min]
			input[min] = input[pointer]
			input[pointer] = newMin
		}
	}
	return input
}

func BastardizedSelectionSort(input []int) []int {
	var done bool
	pointer := 0
	for !done {
		done = true
		for i, el := range input {
			if pointer < len(input) && el < input[pointer] {
				// swap
				input[i] = input[pointer]
				input[pointer] = el
				pointer++
				done = false
			}
		}
	}
	return input
}
 
func InsertionSort(input []int) []int {

	var result = make([]int, len(input))
	pointer := 0

	for _, in := range input {
		for i, out := range result {
			if in < out || i == pointer {
				// shift and insert
				copy(result[i+1:len(result)], result[i:len(result)-1])
				result[i] = in
				pointer++
				break;
			}
		}
	}

	return result
}

func MergeSort(input []int) []int {

	size := len(input)
	if size < 2 {
		// Already sorted return
		return input
	} else {
		var split int
		if size % 2 == 0 {
			split = size/2
		} else {
			split = (size+1)/2
		}
		return merge(MergeSort(input[0:split]), MergeSort(input[split:size]))
	}
}

func ComboMergeInsertionSort(input []int) []int {
	size := len(input)
	if size < 2 {
		return input
	} else if size < 90 {
		return InsertionSort(input)
	} else {
		var split int
		if size % 2 == 0 {
			split = size/2
		} else {
			split = (size+1)/2
		}
		return merge(MergeSort(input[0:split]), MergeSort(input[split:size]))
	}
}

func merge(inOne []int, inTwo []int) []int {

	i, j := 0, 0
	sizeOne := len(inOne)
	sizeTwo := len(inTwo)
	sizeTotal := sizeOne + sizeTwo
	var result = make([]int, sizeTotal)

	for k := 0; k < sizeTotal; k++ {

		// Check if either array is empty
		if i == sizeOne {
			copy(result[k:len(result)], inTwo[j:len(inTwo)])
			break
		} else if j == sizeTwo {
			copy(result[k:len(result)], inOne[i:len(inOne)])
			break;
		} else if inOne[i] < inTwo[j] {
			result[k] = inOne[i]
			i++
		} else {
			result[k] = inTwo[j]
			j++
		}
	}

	return result
}

func ClosestPair(input []int) []int {
	sorted := MergeSort(input)
	result := sorted[:2]
	dist := result[1] - result[0]
	for i := 2; i < len(sorted); i++ {
		if newDist := sorted[i] - sorted[i-1]; newDist < dist {
			result[0] = sorted[i-1]
			result[1] = sorted[i]
			dist = newDist
		}
	}
	return result
}

func ClosestPair2D(input [][]int) [][]int {
	return nil
}


// Util methods

func arrayToString(input []int) string {

	if len(input) < 1 {
		return "[]"
	}

	output := "["

	for _, i := range input {
		output = output + strconv.Itoa(i) + ", "
	}

	output = output[0:len(output)-2]
	output = output + "]"
	return output
}

func isSorted(input []int) bool {
	for i, el := range input {
		if i > 1 && input [i-1] > el {
			return false
		}
	}
	return true
}

func arraysEqual(one, two []int) bool {

	copyOne := make([]int, len(one))
	copy(copyOne, one)
	copyTwo := make([]int, len(two))
	copy(copyTwo, two)

	for _, el := range copyOne {
		for j, el2 := range copyTwo {
			if el == el2 {
				copyTwo = append(copyTwo[:j], copyTwo[j+1:]...)
				break
			}
		}
	}
	return len(copyTwo) == 0
}

// Create an array of integers of specified size
func generateRandomArray(size int) []int {

	var result = make([]int, size)
	for i, _ := range result {
		result[i] = rand.Int()
	}
	return result
}