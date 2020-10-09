package main

import "fmt"

func main() {

	arr := []int{2, 3, 6, 1, 9, 8, 14, 12}

	fmt.Println(InsertionSort(arr))
}

// 插入排序算法实现
func InsertionSort(arr []int) []int {

	// 这个只是最外层循环
	for i := 0; i < len(arr)-1; i++ {

		for arr[i+1] < arr[i] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			// 进行递减数组下标,往前比较两个数大小
			if i > 0 {
				i--
			}
		}
	}
	return arr
}
