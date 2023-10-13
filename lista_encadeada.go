package main

import "fmt"

type Node struct {
	data int
	name string
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Display() {
	current := list.head

	for current != nil {
		fmt.Printf("data -> %d\n", current.data)
		fmt.Printf("name -> %s \n", current.name)
		current = current.next
	}

}

func (list *LinkedList) IsOnList(data int) bool {
  if list.IsEmpty() {
    return false
  }

  current := list.head

  for current != nil  {
    if(current.data == data) {
      return true
    }
    current = current.next
  }

  return false

}


func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list *LinkedList) SizeOf() int {
	current := list.head
	size := 0
	for current != nil {
		size++
		current = current.next
	}

	return size
}

func (list *LinkedList) Append(data int, text string) {
	newNode := &Node{data, text, nil}

	if list.IsEmpty() {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode

}

func (list *LinkedList) AppendInTheHead(data int, text string) {
	newNode := &Node{data, text, nil}

	if list.IsEmpty() {
		list.head = newNode
		return
	}

	newNode.next = list.head
	list.head = newNode
}

func (list *LinkedList) AppendBefore(data int, text string, nodeData int) {
	newNode := &Node{data, text, nil}

	if list.IsEmpty() {
		list.head = newNode
		return
	}

	current := list.head
	for current != nil {
		if current.data == nodeData {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}

}

func (list *LinkedList) AppendAfter(data int, text string, nodeData int) {
	newNode := &Node{data, text, nil}

	if list.IsEmpty() {
		list.head = newNode
		return
	}

	current := list.head
	for current != nil {
		if current.data == nodeData {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}

}

func (list *LinkedList) AppendInOrder(data int, text string) {
	newNode := &Node{data, text, nil}

	if list.IsEmpty() || data < list.head.data {
		newNode.next = list.head
		list.head = newNode
		return
	}

	current := list.head

	for current.next != nil && current.next.data <= data {
		current = current.next
	}

	newNode.next = current.next
	current.next = newNode

}

func (list *LinkedList) Remove(data int) {
	if list.IsEmpty() == nil {
		return
	}

	if list.head.data == data {
		list.head = list.head.next
	}

	curret := list.head

	for curret.next != nil && curret.data != data {
		curret = curret.next
	}

	if curret.next != nil {
		curret.next = curret.next.next
	}

}

func (list *LinkedList) InsertionSort() {
	if list.IsEmpty() || list.head.next == nil {
		return
	}

	sorted := &LinkedList{}
	current := list.head

	for current != nil {
		next := current.next

		sorted.InsertedSort(current)

		current = next
	}

	list.head = sorted.head

}
func (list *LinkedList) InsertedSort(newNode *Node) {
	if list.IsEmpty() || list.head.data >= newNode.data {
		newNode.next = list.head
		list.head = newNode
		return
	}

	current := list.head

	for current.next != nil && current.next.data <= newNode.data {
		current = current.next
	}

	newNode.next = current.next
	current.next = newNode
}

func (list *LinkedList) AppendAfer(data int, text string, nodeData int) {
	newNode := &Node{data, text, nil}

	if list.IsEmpty() {
		list.head = newNode
		return
	}

	current := list.head
	for current != nil {
		if current.data == nodeData {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}

}

func main() {
	myLinkedList := LinkedList{}
	myLinkedList.Append(1, "olá mundo")
	myLinkedList.Append(2, "segundo insert")
	myLinkedList.AppendInOrder(0, "primeiro da lista")
	myLinkedList.AppendBefore(10, "texto", 1)
	myLinkedList.AppendAfter(12, "teste append after", 10)
	fmt.Printf("O tamanho da lista é: %d\n", myLinkedList.SizeOf())
	myLinkedList.Display()
	fmt.Println("Lista ondernada")
	myLinkedList.InsertionSort()
	myLinkedList.Display()
}
