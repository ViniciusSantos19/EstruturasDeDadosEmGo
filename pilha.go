package main

import "fmt"

type Stacknode struct {
	data int
	text string
	next *Stacknode
}

type Stack struct {
	top *Stacknode
}

func NewStack() *Stack {
	return &Stack{top: nil}
}

func (stack *Stack) IsEmpty() bool {
	return stack.top == nil
}

func (stack *Stack) Push(data int, text string) {
	newNode := &Stacknode{data, text, nil}

	newNode.next = stack.top
	stack.top = newNode
}

func (stack *Stack) Pop() (*Stacknode, bool) {
	if stack.IsEmpty() {
		return nil, false
	}

	current := stack.top
	stack.top = stack.top.next
	return current, true

}

func (s *Stack) Peek() (*Stacknode, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	return s.top, true
}

func main() {
	stack := NewStack()
	stack.Push(1, "primeiro item a ilha")
	stack.Push(2, "segundo item da pilha")
	stack.Push(3, "terceiro item da pilha")
	newNode, ok := stack.Pop()
	fmt.Println("elemento desempilhado:", newNode.data, newNode.text)
	fmt.Println("ok", ok)
	newNode, ok = stack.Peek()
	fmt.Println("elemento no top da pilha:", newNode.data, newNode.text)
	fmt.Println("ok", ok)

}
