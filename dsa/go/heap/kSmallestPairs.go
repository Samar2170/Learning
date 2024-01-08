package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	indexes []int
	value   int
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

func (h *PairHeap) Empty() bool {
	if len(*h) == 0 {
		return true
	}
	return false
}

func kSmallestPairs2(nums1 []int, nums2 []int, k int) [][]int {
	pq := &PairHeap{}
	heap.Init(pq)
	i, j := 0, 0

	for i < len(nums1) {
		j = 0
		for j < len(nums2) {
			cSum := nums2[j] + nums1[i]
			p := Pair{[]int{nums1[i], nums2[j]}, cSum * -1}
			heap.Push(pq, p)
			j++
		}
		i++
	}
	ans := [][]int{}
	n := 0
	for n < k {
		ce := heap.Pop(pq).(Pair)
		ans = append(ans, ce.indexes)
		n++
	}
	return ans
}
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	res := [][]int{}
	pq := &PairHeap{}
	heap.Init(pq)

	for i, v1 := range nums1 {
		heap.Push(pq, Pair{[]int{i, 0}, -1 * (v1 + nums2[0])})
	}

	for !pq.Empty() && k > 0 {
		currentMin := heap.Pop(pq).(Pair)
		i, j := currentMin.indexes[0], currentMin.indexes[1]
		// fmt.Println(i, j)
		res = append(res, []int{nums1[i], nums2[j]})

		// fmt.Println(res)
		nextElement := j + 1
		if nextElement < len(nums2) {
			heap.Push(pq, Pair{[]int{i, nextElement}, -1 * (nums1[i] + nums2[nextElement])})
		}
		k--
	}
	return res
}

// func kSmallestPairs3(nums1 []int, nums2 []int, k int) [][]int {
// 	i, j := 0, 0
// 	pq := &PairHeap{}
// 	heap.Init(pq)

// 	res := [][]int{}
// 	for i < len(nums1) && j < len(nums2) && len(res) < k*2 {
// 		pair := Pair{[]int{nums1[i], nums2[j]}, -1 * (nums1[i] + nums2[j])}
// 		heap.Push(pq, pair)
// 		if j == i {
// 			j++
// 		} else if j > i {
// 			i++
// 		}
// 	}
// 	fmt.Println(pq)
// 	return res
// }

func main() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	fmt.Println(kSmallestPairs(nums1, nums2, k))

	nums1 = []int{1, 1, 2}
	nums2 = []int{1, 2, 3}
	k = 2
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}
