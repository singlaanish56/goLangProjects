package algo

import (
	"container/heap"
	"os"
	"strconv"
	"sync"

	//"fmt"

	"github.com/singlaanish56/compressionToolGo/pkg/errors"
	"github.com/singlaanish56/compressionToolGo/pkg/helper"
)

var fileMutex sync.RWMutex

func HuffmanCompress(lines []byte, output string) {

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
		internalNode.SetLeftEdge(0)
		internalNode.SetRightEdge(1)
		heap.Push(&pq, helper.InitItem(internalNode, combinedWeight, 0))
	}

	//var prefixMap map[rune]int




	// output filename
	// add the header which includes the character frequency map
	f, err := os.OpenFile(output, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
	errors.HandleError(err)

	defer f.Close()

	fileMutex.Lock()
	defer fileMutex.Unlock()
	_, e1 := f.WriteString("------------------header start---------------------\n")
	errors.HandleError(e1)

	linesAdded := 3

	for k, v := range freqMap{
		s := string(k) + " : " + strconv.Itoa(v) + "\n"
		_, e := f.WriteString(s)
		errors.HandleError(e)
		linesAdded++
	}

	_, e2 := f.WriteString("------------------header end---------------------\n")
	errors.HandleError(e2)

}
