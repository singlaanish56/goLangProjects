package algo

import (
	"container/heap"
	//"fmt"

	"github.com/singlaanish56/compressionToolGo/pkg/helper"
)

func HuffmanCompress(lines []byte) {

	freqMap := CalculateTheFrequency(lines)

	pq := make(helper.PriorityQueue, len(freqMap))
	i := 0

	//create the weighted huffman tree.
	//each alphabet and frequency is a node, put the nodes in a priority queue, ordered by freq(less first)
	for k, v := range freqMap {

		node := helper.InitLeafNode(k, v)
		pq[i] = helper.InitItem(node, node.Weight(), i)
		i++
	}

	heap.Init(&pq)
	// pickup the first two and join them to form a left and right node with total weight as the internal node
	//push it back in the priority queue to repeat the process
	for pq.Len() > 1 {
		item1 := heap.Pop(&pq).(*helper.Item)
		item2 := heap.Pop(&pq).(*helper.Item)

		combinedWeight := item1.Value().Weight() + item1.Value().Weight()
		internalNode := helper.InitInternalNode(combinedWeight, item1.Value(), item2.Value())
		heap.Push(&pq, helper.InitItem(internalNode, combinedWeight, 0))
	}

	//assign the codes to the tree, left node gets a 0 and right node gets a 1
}
