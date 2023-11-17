package main

import "fmt"

const Alphabet = 26

// Representa um nó do trie
type Node struct{
	children [Alphabet]*Node
	isEnd bool
}

// Representa um trie e tem um ponteiro apontado para raiz
type Trie struct{
	root *Node

}

// Cria um novo Trie
func InitTrie() *Trie{
	result := &Trie{root:&Node{}}
	return result
}

// Insere uma palavra no trie
func (t *Trie) Insert(w string){
	wordLen := len(w)
	currentNode := t.root
	for i:=0; i<wordLen; i++{
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil{
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Procura uma palavra na arvore e retorna true caso a encontre
func (t *Trie) Search(w string) bool{
	wordLen := len(w)
	currentNode := t.root
	for i:=0; i<wordLen; i++{
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil{
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.isEnd

}

func main(){
	newTrie := InitTrie()

	toAdd := []string{
		"aragorn",
		"aragog",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	for _, value := range(toAdd){
		newTrie.Insert(value)
	}

	fmt.Println(newTrie.Search("oregano"))
	fmt.Println(newTrie.Search("golang"))
}