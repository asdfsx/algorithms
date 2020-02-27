package sort

func InsertSort(a []int) {
	length := len(a)
	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
}

func InsertSort2(a []int) {
	length := len(a)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
