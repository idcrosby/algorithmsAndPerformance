package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	)

func main() {
	test(generateRandomArray(1000000))
}

func test(input []int) {
	// fmt.Println("Input: " + arrayToString(input))
	start := time.Now()
	// output := BubbleSort(input)
	output := input
	elapsed := time.Since(start)
	if isSorted(output) {
		fmt.Printf("Bubble sort passed. Took: %s\n", elapsed)
	} else {
		// fmt.Println("Bubble sort failed: " + arrayToString(output))
	}
	start = time.Now()
	output = InsertionSort(input)
	elapsed = time.Since(start)
	if isSorted(output) {
		fmt.Printf("Insertion sort passed. Took: %s\n", elapsed)
	} else {
		fmt.Println("Insertion sort failed: " + arrayToString(output))
	}
	start = time.Now()
	output = SelectionSort(input)
	elapsed = time.Since(start)
	if isSorted(output) {
		fmt.Printf("Selection sort passed. Took: %s\n", elapsed)
	} else {
		fmt.Println("Selection sort failed: " + arrayToString(output))
	}
	start = time.Now()
	output = MergeSort(input)
	elapsed = time.Since(start)
	if isSorted(output) {
		fmt.Printf("Merge sort passed. Took: %s\n", elapsed)
	} else {
		fmt.Println("Merge sort failed: " + arrayToString(output))
	}
	start = time.Now()
	output = ComboMergeInsertionSort(input)
	elapsed = time.Since(start)
	if isSorted(output) {
		fmt.Printf("Combo sort passed. Took: %s\n", elapsed)
	} else {
		fmt.Println("Combo sort failed: " + arrayToString(output))
	}
	// fmt.Println("Output: " + arrayToString(output))
}

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
			// fmt.Printf("Swapping index %d with minimum %d at index %d \n", pointer, newMin, min)
			// fmt.Println("" + arrayToString(input))

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

// Create an array of integers of specified size
func generateRandomArray(size int) []int {

	var result = make([]int, size)
	for i, _ := range result {
		result[i] = rand.Int()
	}
	return result
}