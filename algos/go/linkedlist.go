package main

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Insert(value int) {
	newNode := Node{Value: value, Next: nil}

	if ll.Head == nil {
		ll.Head = &newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &newNode
}

func createLinkedList(nums []int) LinkedList {
	i, l := 0, len(nums)
	ll := LinkedList{}
	for i < l {
		ll.Insert(nums[i])
		i++
	}
	return ll
}

func (ll *LinkedList) Display() {
	current := ll.Head
	for current.Next != nil {
		println(current.Value)
		current = current.Next
	}
	println(current.Value)
}
func (ll *LinkedList) Search(val int) bool {
	current := ll.Head
	for current.Next != nil {
		if current.Value == val {
			return true
		}
		current = current.Next
	}
	if current.Value == val {
		return true
	}
	return false
}

func (ll *LinkedList) InsertAtStart(val int) {
	currentHead := ll.Head
	newNode := Node{
		Value: val,
		Next:  currentHead,
	}
	ll.Head = &newNode
}
func (ll *LinkedList) InsertAt(val, position int) {
	if position == 0 {
		ll.InsertAtStart(val)
		return
	}
	i := 0
	current := ll.Head
	newNode := Node{Value: val, Next: nil}

	for current.Next != nil && i < position-1 {
		current = current.Next
		i++
	}
	if current == nil {
		return
	}
	newNode.Next = current.Next
	current.Next = &newNode
}

func (ll *LinkedList) DeleteHead() {
	ll.Head = ll.Head.Next
}
func (ll *LinkedList) Pop() {
	current := ll.Head
	for current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.Head == nil
}
func ReverseLinkedList(list1 *LinkedList) *LinkedList {
	llNew := LinkedList{}
	c := list1.Head

	for c != nil {
		llNew.InsertAtStart(c.Value)
		c = c.Next
	}
	return &llNew
}

func MergeLL(list1, list2 *LinkedList) *LinkedList {
	mergedList := LinkedList{}
	c1, c2 := list1.Head, list2.Head

	for c1 != nil && c2 != nil {
		if c1.Value < c2.Value {
			mergedList.Insert(c1.Value)
			c1 = c1.Next
		} else {
			mergedList.Insert(c2.Value)
			c2 = c2.Next
		}
	}
	for c1 != nil {
		mergedList.Insert(c1.Value)
		c1 = c1.Next
	}
	for c2 != nil {
		mergedList.Insert(c2.Value)
		c2 = c2.Next
	}
	return &mergedList
}

func (ll *LinkedList) Length() int {
	current := ll.Head
	count := 0
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func (ll *LinkedList) HasCycle() bool {
	if ll.Head == nil {
		return false
	}
	slow, fast := ll.Head, ll.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	ll := createLinkedList(nums)
	println("Displaying linked list")
	ll.Display()

	println(ll.Search(5))
	println("Searching for 5")

	println("Displaying linked list after inserting 0 at start")
	ll.InsertAtStart(0)
	ll.Display()

	println("Displaying linked list after inserting 9 at position 4")
	ll.InsertAt(9, 4)
	ll.Display()

	println("Displaying linked list after deleting head")
	ll.DeleteHead()
	ll.Display()

	println("Displaying linked list after popping")
	ll.Pop()
	ll.Display()

	println("Displaying linked list after merge")
	l1 := LinkedList{}
	l1.Insert(1)
	l1.Insert(3)
	l1.Insert(5)
	l2 := LinkedList{}
	l2.Insert(2)
	l2.Insert(4)
	l2.Insert(6)

	l3 := MergeLL(&l1, &l2)
	l3.Display()

	println("Displaying linked list after reversing")
	l4 := ReverseLinkedList(l3)
	l4.Display()

	println("Length of linked list")
	println(l4.Length())

	println("Checking if linked list has cycle")
	println(l4.HasCycle())

	// Creating a cycle
	l4.Head.Next.Next.Next.Next.Next = l4.Head
	println(l4.HasCycle())

}
