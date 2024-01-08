package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	index, value int
}

type PairHeap []Pair

func (h PairHeap) Len() int { return len(h) }
func (h PairHeap) Less(i, j int) bool {
	return h[i].value > h[j].value || (h[i].value == h[j].value && h[i].index < h[j].index)
}
func (h PairHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)
	i, j, index := 0, 0, 0

	pq := &PairHeap{}
	heap.Init(pq)

	for j < len(nums) {
		fmt.Println(i, j)

		pair := Pair{j, nums[j]}
		heap.Push(pq, pair)
		fmt.Println(pq)

		if j-i+1 < k {
			j++
		} else if j-i+1 == k {
			for (*pq)[0].index < i {
				heap.Pop(pq)
			}
			first := (*pq)[0]
			res[index] = first.value
			if first.index < i+1 {
				heap.Pop(pq)
			}
			i++
			j++
			index++
		}
	}

	return res
}

func main() {
	s1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(s1, 3))

	s2 := []int{1, -1}
	fmt.Println(maxSlidingWindow(s2, 2))

	s3 := []int{9}
	fmt.Println(maxSlidingWindow(s3, 1))

}
