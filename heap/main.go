package main

import "fmt"


type MaxHeap struct {
	slice []int
}

// Adiciona um elemente ao heap(monte)
func (h *MaxHeap) Insert(key int){
	h.slice = append(h.slice, key)
	h.maxHeapifyUp(len(h.slice) -1)
}

// Retorna a maior chave e a remove do heap
func (h *MaxHeap) Extract() int{
	extracted := h.slice[0]

	l := len(h.slice) - 1

	if len(h.slice) == 0 {
		fmt.Println("Impossível extrair porque o tamanho do slice é 0")
		return - 1
	}

	// Coloca o ultimo index na raiz
	h.slice[0] = h.slice[l]
	h.slice = h.slice[:l]
	h.maxHeapifyDown(0)

	return extracted
}

// Reordena a arvore de baixo para cima
func (h *MaxHeap) maxHeapifyUp(index int){
	for h.slice[parent(index)] < h.slice[index]{
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// Reordena a arvore de cima para baixo
func (h *MaxHeap) maxHeapifyDown(index int){
	lastIndex := len(h.slice) - 1
	l,r := left(index), right(index)
	childToCompare := 0

	// Loop para quando o index tem pelo menos um nó filho
	for l <= lastIndex{
		if l == lastIndex{ // Quando o nó filho esquerdo é o único filho
			childToCompare = l
		} else if h.slice[l] > h.slice[r]{ // Quando o nó filho esquerdo é maior
			childToCompare = l
		} else {  // Quando o nó filho direito é maior
			childToCompare = r
		}

		// Compara o valor do slice no index com o nó filho maior e troca se for menor
		if h.slice[index] < h.slice[childToCompare]{
			h.swap(index, childToCompare)
			index = childToCompare
			l,r = left(index), right(index)
		} else {
			return
		}
	}
}

// Retorna o pai do index do nó passado
func parent(i int) int {
	return (i-1) / 2
}

// Retorna o index do nó filho esquerdo(menor)
func left(i int) int{
	return 2 * i + 1
}

// Retorna o index do nó filho direto(maior)
func right(i int) int{
	return 2 * i + 2
}

// Muda as chaves do slice
func (h *MaxHeap) swap(i1,i2 int){
	h.slice[i1], h.slice[i2] = h.slice[i2], h.slice[i1]
}

func main(){

	h := &MaxHeap{}
	buildHeap := []int{10,20,30,5,7,9,11,13,15,17}

	for _, v := range buildHeap{
		h.Insert(v)
		fmt.Println(h)
	}

	fmt.Println("---------------------------")

	for i:=0;i<5;i++{
		h.Extract()
		fmt.Println(h)
	}
}