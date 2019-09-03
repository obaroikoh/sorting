package sorting

func swap(array []int, firstIndex, secondIndex int) {
	temp := array[firstIndex]
	array[firstIndex] = array[secondIndex]
	array[secondIndex] = temp
}

func indexOfMin(array []int, startIndex, n int) int {
	minValue := array[startIndex]
	minIndex := startIndex
	for i := minIndex + 1; i < n; i++ {
		if array[i] < minValue {
			minIndex = i
			minValue = array[i]
		}
	}
	return minIndex
}

//SelectionSort Implementation
func SelectionSort(array []int, n int) {
	for i := 0; i < n; i++ {
		index := indexOfMin(array, i, n)
		swap(array, i, index)
	}
}
