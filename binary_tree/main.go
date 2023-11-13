package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

var count int

// Procura um valor na árvore e retorna true se encontrar
func (n *Node) Search(k int) bool {

	count ++

	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}

	return true
}

// Adiciona um nó caso o valor não exista
func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func main() {
	tree := &Node{Key: 10}
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(5)
	tree.Insert(9)
	fmt.Println(tree)
	fmt.Println(tree.Search(9))
	fmt.Println(count)
}
