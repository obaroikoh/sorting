package sorting

func InsertionSort(array []int, n int) {
	var i, j, key int
	for i = 1; i < n; i++ {
		j = i
		for j > 0 && array[j-1] > array[j] {
			key = array[j]
			array[j] = array[j-1]
			array[j-1] = key
			j--
		}
	}
}
