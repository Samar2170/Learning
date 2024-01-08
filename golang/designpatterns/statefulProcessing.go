package main

import "fmt"

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
	// fmt.Printf("Event handled. Total events: %d\n", e.count)
}

func (c *Counter) GetValue() int {
	return c.count
}

type ErrorSupervisor struct {
	counter Counter
}

func main() {

	counterInstance := &Counter{}
	errCounter := &Counter{}
	er := &ErrorSupervisor{counter: *errCounter}
	incrementFn := (*Counter).Increment
	valueFn := (*Counter).GetValue

	incrementFn(counterInstance)
	incrementFn(counterInstance)

	incrementFn(&er.counter)
	incrementFn(&er.counter)
	incrementFn(&er.counter)
	fmt.Println(valueFn(counterInstance))
	fmt.Println(valueFn(&er.counter))
}
