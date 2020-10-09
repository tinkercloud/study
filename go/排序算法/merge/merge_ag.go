package main

import "fmt"

func main() {
	ar := []int{1, 5, 8, 2, 4, 14, 11, 9}

	rs := MergeSort(ar)
	fmt.Println(rs)
}

// 拆分左右树
func MergeSort(arr []int) []int {

	// 如果数组只剩下一个元素,返回
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2

	left := MergeSort(arr[0:mid])
	fmt.Printf("Left = %d\n", left)

	right := MergeSort(arr[mid:])

	fmt.Printf("Right = %d\n", right)

	return MergeSortArray(left, right)

}

func MergeSortArray(left []int, right []int) []int {
	sortArray := make([]int, len(left)+len(right))

	l := 0
	r := 0

	for i := 0; i < len(sortArray); i++ {

		if l >= len(left) {
			sortArray[i] = right[r]
			r++
			continue
		} else if r >= len(right) {
			sortArray[i] = left[l]
			l++
			continue
		}

		if left[l] > right[r] {
			sortArray[i] = right[r]
			r++
		} else {
			sortArray[i] = left[l]
			l++
		}

	}
	return sortArray

}
