package main

type Stack struct {
	Items   []int
	MaxSize int
}

func (s *Stack) Push(val int) {
	if len(s.Items) < s.MaxSize {
		s.Items = append(s.Items, val)
	}
}

func (s *Stack) Pop() int {
	if len(s.Items) == 0 {
		return -1
	}
	popped := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return popped
}

func (s *Stack) Peek() int {
	if len(s.Items) == 0 {
		return -1
	}
	return s.Items[len(s.Items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stack) IsFull() bool {
	return len(s.Items) == s.MaxSize
}

func (s *Stack) Size() int {
	return len(s.Items)
}

func (s *Stack) Clear() {
	s.Items = []int{}
}

func (s *Stack) Search(val int) int {
	for i := len(s.Items) - 1; i >= 0; i-- {
		if s.Items[i] == val {
			return len(s.Items) - i
		}
	}
	return -1
}

func main() {
	s := Stack{MaxSize: 5}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
}
