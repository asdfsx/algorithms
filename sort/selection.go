package sort

func SelectionSort(a []int) {
	length := len(a)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}
