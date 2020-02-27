package sort_test

import (
	"fmt"
	"github.com/asdfsx/algorithms/sort"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("sort test", func() {
	Context("bubble sort test", func(){
		It("simple test", func() {
                        array := []int{9, 8, 2, 1, 7, 5, 6, 3, 4}
                        sort.BubbleSort(array)
                        fmt.Println(array)
                        Expect(array).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
                })
	})
	Context("selection sort test", func() {
		It("simple test", func() {
			array := []int{9, 8, 2, 1, 7, 5, 6, 3, 4}
			sort.SelectionSort(array)
			fmt.Println(array)
			Expect(array).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
		})
	})
	Context("insert sort test", func() {
		It("simple test", func() {
			array := []int{9, 8, 2, 1, 7, 5, 6, 3, 4}
			sort.InsertSort(array)
			fmt.Println(array)
			Expect(array).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
		})
	})
	Context("shell sort test", func() {
		It("simple test", func() {
			array := []int{9, 8, 2, 1, 7, 5, 6, 3, 4, 11, 13, 21, 14, 15, 19, 18, 17, 16, 10}
			sort.ShellSort(array)
			fmt.Println(array)
			Expect(array).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15, 16, 17, 18, 19, 21}))
		})
	})
	Context("merge sort test", func() {
		It("simple test", func() {
			array := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
			sort.MergeSort(array, 0, (9+0)/2, 9)
			fmt.Println(array)
			Expect(array).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
		})
		It("simple test2", func() {
			array := []int{1, 3, 2}
			sort.MergeSort(array, 0, 1, 2)
			fmt.Println(array)
			Expect(array).To(Equal([]int{1, 2, 3}))
		})
		It("simple test3", func() {
			array := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
			result := sort.MergeSort2(array, 0, (9+0)/2, 9)
			fmt.Println(result)
			Expect(result).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
		})
		It("simple test4", func() {
			array := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
			sort.MergeSort3(array, 0, 9)
			fmt.Println(array)
			Expect(array).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
		})
		It("simple test5", func() {
			array := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
			sort.MergeSort4(array)
			fmt.Println(array)
			Expect(array).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
		})
	})
})
