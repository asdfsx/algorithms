package sort

func ShellSort(a []int) {
	length := len(a)
	// 计算gap的值
	for gap := length / 2; gap > 0; gap /= 2 {
		// 将数组变为gap有序
		for i := gap; i < length; i += 1 {
			// 将a[i] 插入到a[i-gap],a[i-2*gap],a[i-3*gap]....中
			for j := i; j >= gap; j -= gap {
				if a[j-gap] > a[j] {
					a[j-gap], a[j] = a[j], a[j-gap]
				}
			}
		}
	}
}
