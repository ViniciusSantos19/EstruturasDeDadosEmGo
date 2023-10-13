package main

import "fmt"

type Node struct {
	data     int
	text     string
	next     *Node
	previous *Node
}

type DoubLinkedList struct {
	head *Node
	tail *Node
}

func (list *DoubLinkedList) Display() {
	current := list.head

	for current != nil {
		fmt.Printf("data -> %d \n", current.data)
		fmt.Printf("text -> %s \n", current.text)
		current = current.next
	}

}

func (list *DoubLinkedList) DisplayBackward() {
	current := list.tail

	for current != nil {
		fmt.Printf("data -> %d \n", current.data)
		fmt.Printf("text -> %s \n", current.text)
		current = current.previous
	}

}

func (list *DoubLinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list *DoubLinkedList) Append(data int, text string) {
	newNode := &Node{data, text, nil, nil}

	if list.IsEmpty() {
		list.head = newNode
		list.tail = newNode
		return
	}

	newNode.previous = list.tail
	list.tail.next = newNode
	list.tail = newNode
}

func (list *DoubLinkedList) AppendInTheHead(data int, text string) {
	newNode := &Node{data, text, nil, nil}

	if list.IsEmpty() {
		list.head = newNode
		list.tail = newNode
		return
	}

	newNode.next = list.head
	list.head.previous = newNode
	list.head = newNode
}

func (list *DoubLinkedList) AppendAfter(data int, text string, nodeData int) {
	newNode := &Node{data, text, nil, nil}

	if list.IsEmpty() {
		list.head = newNode
		list.tail = newNode
		return

	}

	current := list.head
	for current != nil {
		if current.data == nodeData {
			newNode.previous = current
			newNode.next = current.next

			if current.next != nil {
				current.next.previous = newNode
			} else {
				list.tail = newNode
			}

			current.next = newNode

			return

		}
		current = current.next
	}

}

func (list *DoubLinkedList) AppendBefore(data int, text string, nodeData int) {
	newNode := &Node{data, text, nil, nil}

	if list.IsEmpty() {
		list.head = newNode
		list.tail = newNode
		return

	}

	current := list.head

	if current.data == nodeData {
		newNode.next = current
		current.previous = newNode
		list.head = newNode
		return
	}

	for current != nil {
		if current.data == nodeData {
			newNode.next = current
			newNode.previous = current.previous
			current.previous = newNode

			return

		}
		current = current.next
	}

}

func (list *DoubLinkedList) AppendInOrder(data int, text string) {
	newNode := &Node{data, text, nil, nil}

	if list.IsEmpty() || data <= list.head.data {
		newNode.next = list.head
		list.head.previous = newNode
		list.head = newNode

		if list.tail == nil {
			list.tail = newNode
		}

		return
	}

	current := list.head
	for current != nil && current.next != nil && current.next.data <= data {
		current = current.next
	}

	newNode.next = current.next
	newNode.previous = current
	current.next = newNode

	if newNode.next != nil {
		newNode.next.previous = newNode
	} else {
		list.tail = newNode
	}
}

func main() {
	doubleLikedList := DoubLinkedList{}
	doubleLikedList.Append(1, "primeiro dado da lista")
	doubleLikedList.Append(2, "segundo dado da lista")
	doubleLikedList.AppendInTheHead(0, "inserindo no inicio da lista")
	doubleLikedList.AppendBefore(3, "teste append before no head", 0)
	doubleLikedList.AppendBefore(21, "teste previous de novo", 3)
	doubleLikedList.AppendAfter(17, "teste append after no head", 21)
	doubleLikedList.AppendBefore(24, "teste append before no tail", 2)
	doubleLikedList.AppendAfter(66, "teste de append after no tail", 2)
	doubleLikedList.AppendInOrder(64, "teste append em odem no in order")
	fmt.Println("imprimindo ao contrario")
	doubleLikedList.DisplayBackward()
	fmt.Println("imprimindo em ordem crescente")
	doubleLikedList.Display()
}
