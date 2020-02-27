package sort

func MergeSort(a []int, lo, mid, hi int) {
	i, j := lo, mid+1

	tmp := make([]int, len(a))
	for k := lo; k <= hi; k++ {
		tmp[k] = a[k]
	}

	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = tmp[j]
			j += 1
		} else if j > hi {
			a[k] = tmp[i]
			i += 1
		} else if tmp[i] > tmp[j] {
			a[k] = tmp[j]
			j += 1
		} else {
			a[k] = tmp[i]
			i += 1
		}
	}
}

func MergeSort2(a []int, lo, mid, hi int) []int {
	i, j := lo, mid+1
	tmp := []int{}

	for k := 0; k <= hi; k++ {
		if i > mid {
			tmp = append(tmp, a[j])
			j++
		} else if j > hi {
			tmp = append(tmp, a[i])
			i++
		} else if a[i] > a[j] {
			tmp = append(tmp, a[j])
			j++
		} else {
			tmp = append(tmp, a[i])
			i++
		}
	}
	return tmp
}

func MergeSort3(a []int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	MergeSort3(a, lo, mid)
	MergeSort3(a, mid+1, hi)
	MergeSort(a, lo, mid, hi)
}

func MergeSort4(a []int) {
	length := len(a)
	for i := 1; i < length; i++ {
		for j := 0; j < i; j += (i + i) {
			if (i + j + i - 1) > (length - 1) {
				MergeSort(a, j, j+i-1, length-1)
			} else {
				MergeSort(a, j, j+i-1, i+j-i-1)
			}
		}
	}
}
