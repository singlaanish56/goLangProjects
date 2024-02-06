package algo

import (
	"container/heap"
	"fmt"

	"github.com/singlaanish56/compressionToolGo/pkg/helper"
)



func HuffmanCompress(lines []byte) {

	freqMap := CalculateTheFrequency(lines)

	pq := make(helper.PriorityQueue, len(freqMap))
	i:=0

	//create the weighted huffman tree.
	//each alphabet and frequency is a node, put the nodes in a priority queue, ordered by freq(less first)
	for k, v := range freqMap {
		fmt.Printf("%s:%v\n", string(k), v)

		node := helper.HuffmanLeafNode{k, v}
		pq[i] = &helper.Item{
			value: node,
			priority: node.Weight(),
			index: i,
		}
		i++
	}

	// pickup the first two and join them to form a left and right node with total weight as the internal node
	//push it back in the priority queue to repeat the process
	for pq.len()>0{
		item := heap.Pop(&pq).(*helper.Item)
		fmt.Printf("%s:%d, %d\n",string(item.value.Element()), item.value.Weight(), item.priority)
	}
}
