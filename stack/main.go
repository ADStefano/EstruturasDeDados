package main

import "fmt"

type Stack struct{
	items []int
}

// Adiciona um item na pilha
func (s *Stack) Push(i int){
	s.items = append(s.items, i)	
}

// Remove o ultimo item da pilha e retorna o valor removido
func (s *Stack) Pop() int{
	lastItem := len(s.items) - 1
	s.items = s.items[:lastItem]
	return lastItem
}

func main(){
	s := Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	s.Push(50)
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)

}