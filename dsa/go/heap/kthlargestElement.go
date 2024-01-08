package main

import (
	"container/heap"
)

type Pair struct {
	value int
}

type PairHeap []Pair

func (h PairHeap) Len() int { return len(h) }
func (h PairHeap) Less(i, j int) bool {
	return h[i].value > h[j].value
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

func findKthLargest(nums []int, k int) int {
	pq := &PairHeap{}
	heap.Init(pq)

	for _, n := range nums {
		p := Pair{n}
		heap.Push(pq, p)
	}
	i := 1
	for i < k {
		heap.Pop(pq)
		i++
	}
	ans := heap.Pop(pq).(Pair)
	return ans.value
}

func main() {
	s1 := []int{3, 2, 1, 5, 6, 4}
	println(findKthLargest(s1, 2))

	s2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	println(findKthLargest(s2, 4))

	s3 := []int{1}
	println(findKthLargest(s3, 1))
}
