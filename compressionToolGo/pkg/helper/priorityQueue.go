package helper

import (
	"container/heap"
	"fmt"
)
type Item struct{
	value HuffmanNode
	priority int
	index int
}

func (item *Item)Value() HuffmanNode{
	return item.value	
}
func InitItem(v HuffmanNode, p int, i int) *Item{
	return &Item{value: v, priority: p, index: i}
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq)}

func (pq PriorityQueue) Less(i,j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq *PriorityQueue) Push(x any){
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any{
	old := *pq
	n:=len(old)
	item:=old[n-1]
	old[n-1]= nil
	item.index = -1
	*pq = old[0:n-1]
	return item

}

func (pq PriorityQueue) Swap(i, j int){
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index=j
}

func (pq *PriorityQueue) Update(item *Item, value HuffmanNode, priority int){
	item.value = value
	item.priority = priority
	heap.Fix(pq,item.index)
}

func PrintPriorityQueueItem(item Item){
	fmt.Printf("%s:%d %d\n",string(item.value.(*HuffmanLeafNode).Element()), item.value.Weight(), item.priority)
}