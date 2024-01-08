package main

import "fmt"

type MinStack struct {
	Stack []int
	Min   []int
}

func Constructor() MinStack {
	return MinStack{Stack: []int{}, Min: []int{}}
}

func (this *MinStack) Push(val int) {
	if len(this.Stack) == 0 {
		this.Stack = append(this.Stack, val)
		this.Min = append(this.Min, val)
	} else {
		if val < this.Min[len(this.Min)-1] {
			this.Min = append(this.Min, val)
		} else {
			this.Min = append(this.Min, this.Min[len(this.Min)-1])
		}
		this.Stack = append(this.Stack, val)
	}
}

func (this *MinStack) Pop() {
	this.Min = this.Min[:len(this.Min)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}

func main() {
	obj := Constructor()
	obj.Push(2)
	obj.Push(0)
	obj.Push(3)
	obj.Push(0)
	obj.Pop()
	obj.Pop()

	obj.Pop()
	obj.Pop()
	fmt.Println(obj)
}
