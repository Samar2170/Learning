package main

import (
	"container/heap"
	"fmt"
)

type CSlice []int

func (s CSlice) Indices(i int) []int {
	indices := []int{}
	for indx, el := range s {
		if el == i {
			indices = append(indices, indx)
		}
	}
	return indices
}

type Project struct {
	profit int
}

type ProjectHeap []Project

func (h ProjectHeap) Len() int { return len(h) }
func (h ProjectHeap) Less(i, j int) bool {
	return h[i].profit < h[j].profit
}
func (h ProjectHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *ProjectHeap) Pop() interface{} {
	lastEl := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return lastEl
}
func (h *ProjectHeap) Push(x interface{}) {
	*h = append(*h, x.(Project))
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	steps := 0
	p2 := CSlice(capital)
	for steps < k {
		indices := p2.Indices(w)
		fmt.Println(w, capital)
		fmt.Println(indices)
		projects := ProjectHeap{}
		heap.Init(&projects)
		for _, indx := range indices {
			heap.Push(&projects, Project{profits[indx]})
		}
		fmt.Println(projects)
		finalProject := heap.Pop(&projects).(Project)
		w = finalProject.profit + w
		steps++
	}
	fmt.Println(w)
	return w
}

func main() {
	// profits := []int{1, 2, 3}
	// capital := []int{0, 1, 1}
	// findMaximizedCapital(2, 0, profits, capital)

	profits := []int{1, 2, 3}
	capital := []int{0, 1, 2}
	findMaximizedCapital(3, 0, profits, capital)

}
