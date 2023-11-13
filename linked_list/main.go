package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct{
	head *node
	length int
}

// Adiciona um nó na lista
func (l *LinkedList) prepend(n *node){
	nextHead := l.head
	l.head = n
	l.head.next = nextHead
	l.length ++
}

// Itera sobre a lista e imprime no console o valor dos nós
func (l LinkedList) PrintList(){

	toPrint := l.head

	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length --
	}

	fmt.Println("")

}

// Deleta um valor da lista
func (l * LinkedList) DeleteWithValue(value int){

	if l.length == 0{
		return 
	}

	if l.head.data == value{
		l.head = l.head.next
		l.length --
		return
	}

	previousToDelete := l.head

	for previousToDelete.next.data != value{
		if previousToDelete.next.next == nil {
			return 
		}
		previousToDelete = previousToDelete.next
	}

	previousToDelete.next = previousToDelete.next.next
	l.length -- 
}

func main(){
	linkList := LinkedList{}
	node1 := &node{data:12}
	node2 := &node{data:13}
	node3 := &node{data:14}
	node4 := &node{data:15}

	linkList.prepend(node1)
	linkList.prepend(node2)
	linkList.prepend(node3)
	linkList.prepend(node4)

	linkList.PrintList()
	linkList.DeleteWithValue(1300)
	linkList.PrintList()
	linkList.DeleteWithValue(13)
	linkList.PrintList()
	linkList.DeleteWithValue(15)
	linkList.PrintList()

	emptyList := LinkedList{}
	emptyList.DeleteWithValue(123)
	emptyList.PrintList()



}