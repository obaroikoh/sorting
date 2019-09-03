package sorting

import "fmt"

//BubbleSort implementation
func BubbleSort(array []int, n int) {
	var i, j, temp int
	var flag bool
	for i = 0; i < n; i++ {
		for j = 0; j < n-i-1; j++ {
			flag = false
			if array[j] > array[j+1] {
				temp = array[j]
				array[j] = array[j+1]
				array[j+1] = temp

				flag = true
			}
		}

		if !(flag) {
			break
		}
	}

	fmt.Println("Sorted Array: ")
	for i = 0; i < n; i++ {
		fmt.Printf("%d\n", array[i])
	}
}
