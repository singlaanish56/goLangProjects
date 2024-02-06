package stack

import "sync"

type Stack struct {
	items []string

	rwLock sync.RWMutex
}

func New() *Stack{
	stack := &Stack{items:[]string{}}

	return stack
}

func (stack *Stack) Push(item string){

	stack.rwLock.Lock()

	stack.items = append(stack.items, item)

	stack.rwLock.Unlock()
}

func (stack *Stack) Pop() *string{

	if len(stack.items) ==0 {
		return nil
	}

	stack.rwLock.Lock()

	n:=len(stack.items)
	item := stack.items[n-1]
	stack.items = stack.items[0:n-1]

	stack.rwLock.Unlock()

	return &item
}

func (stack *Stack) Top() *string{
	if len(stack.items) ==0 {
		return nil
	}

	stack.rwLock.Lock()

	n:=len(stack.items)
	item := stack.items[n-1]

	stack.rwLock.Unlock()

	return &item
}

func (stack *Stack) IsEmpty() bool{

	return len(stack.items) == 0
}