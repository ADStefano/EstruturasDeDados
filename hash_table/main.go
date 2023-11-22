package main

import "fmt"

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func Init() *HashTable {
	result := &HashTable{}

	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

// Insere uma chave no HashTable
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Procura uma chave no HashTable
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete uma chave do HashTable
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// Insere uma chave no bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Printf("%s já existe\n", k)
	}
}

// Procua uma chave no bucket
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// Deleta uma chave do bucket
func (b *bucket) delete(k string) {

	if b.head.key == k{
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil{
		if previousNode.next.key == k{
			previousNode.next = previousNode.next.next
		}
		previousNode.next = previousNode.next.next
	}
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func main() {

	hashTable := Init()
	nameList := []string{
		"RANDY",
		"ERIC",
		"KENNY",
		"KYLE",
		"BUTTERS",
		"TOKEN",
		"STAN",
	}

	for _, v := range nameList{
		fmt.Printf("Inserindo: %s\n", v)
		hashTable.Insert(v)
	}

	nameList = append(nameList, "GORLOMI")

	for _, v := range nameList{

		fmt.Println(v,hashTable.Search(v))
	}

	hashTable.Delete("ERIC")

	for _, v := range nameList{

		fmt.Println(v,hashTable.Search(v))
	}
}
