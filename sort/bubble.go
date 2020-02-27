package sort

func BubbleSort(a []int) {
        length := len(a)
        for i := 0; i < length; i++ {
                for j := i; j < length; j++ {
                        if a[i] > a[j] {
                                a[i], a[j] = a[j], a[i]
                        }
                }
        }
}
