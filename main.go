package main

import "fmt"

func main() {
	data := []int{5, 3, 8, 6, 2, 7, 1, 4}
	fmt.Printf("sorted: %v\n", merge(data, 0, len(data)-1))
}

func merge(data []int, src, dst int) []int {
	length := dst - src + 1
	result := make([]int, length)

	if length <= 0 {
		return []int{}
	}

	if length == 1 {
		return []int{data[src]}
	}

	if length == 2 {
		if data[src] <= data[dst] {
			result[0], result[1] = data[src], data[dst]
		} else {
			result[1], result[0] = data[src], data[dst]
		}
		return result
	}

	middle := (src + dst + 1) / 2

	var left, right []int
	left = merge(data, src, middle-1)
	right = merge(data, middle, dst)

	for current, i, j := 0, 0, 0; i < len(left) || j < len(right); current++ {
		// left all filled
		if i == len(left) {
			result[current] = right[j]
			j++
			continue
		}

		// right all filled
		if j == len(right) {
			result[current] = left[i]
			i++
			continue
		}

		// in middle, choose smaller
		if left[i] <= right[j] {
			result[current] = left[i]
			i++
		} else {
			result[current] = right[j]
			j++
		}
	}

	return result
}
