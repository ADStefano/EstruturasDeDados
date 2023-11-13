package main

import "fmt"

type Queue struct {
	item []int
}

// Adiciona um item no fim da lista
func (q *Queue) Enqueue(i int) {
	q.item = append(q.item, i)
}

// Remove o primeiro item da lista
func (q *Queue) Dequeue() int {
	firstItem := q.item[0]
	q.item = q.item[1:]
	return firstItem

}

func main() {
	q := Queue{}
	q.Enqueue(100)
	q.Enqueue(200)
	q.Enqueue(300)
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
}
