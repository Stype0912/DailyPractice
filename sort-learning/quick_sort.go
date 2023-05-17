package sort_learning

func quickSort(arr []int, begin, end int) {
	if begin < end {
		key := arr[begin]
		i := begin
		j := end
		for i < j {
			for i < j && arr[j] > key {
				j--
			}
			if i < j {
				arr[i] = arr[j]
				i++
			}
			for i < j && arr[i] < key {
				i++
			}
			if i < j {
				arr[j] = arr[i]
				j--
			}
		}
		arr[i] = key
		quickSort(arr, begin, i-1)
		quickSort(arr, i+1, end)
	}
}
