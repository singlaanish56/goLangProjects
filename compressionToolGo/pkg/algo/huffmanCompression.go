package algo

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"sync"
	"unicode"

	"github.com/singlaanish56/compressionToolGo/pkg/errors"
	"github.com/singlaanish56/compressionToolGo/pkg/helper"
)

type QueueItem struct {
	Node   helper.HuffmanNode
	Prefix string
}

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
	fmt.Println("freq map")
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

	fmt.Println("tree made")
	//create a prefix map to store the prefix for each character
	tree := heap.Pop(&pq).(*helper.Item)

	prefixMap := make(map[rune]string)
	huffmanQueue := make([]QueueItem, 0)

	huffmanQueue = append(huffmanQueue, QueueItem{tree.Value(), ""})

	fmt.Println("created the huffman queue")
	for len(huffmanQueue) > 0 {
		tempItem := huffmanQueue[0]
		huffmanQueue = huffmanQueue[1:]
		if tempItem.Node.Leaf() {
			prefixMap[tempItem.Node.(*helper.HuffmanLeafNode).Element()] = tempItem.Prefix
		} else {

			huffmanQueue = append(huffmanQueue, QueueItem{tempItem.Node.(*helper.HuffmanInternalNode).Left(), tempItem.Prefix + "0"})
			huffmanQueue = append(huffmanQueue, QueueItem{tempItem.Node.(*helper.HuffmanInternalNode).Right(), tempItem.Prefix + "1"})
		}
	}

	fmt.Println("created the prefix map")
	// output filename
	// add the header which includes the character frequency map
	f, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	errors.HandleError(err)

	defer f.Close()

	fileMutex.Lock()
	defer fileMutex.Unlock()
	_, e1 := f.WriteString("------------------header start---------------------\n")
	errors.HandleError(e1)

	linesAdded := 3

	for k, v := range freqMap {
		s := string(k) + " : " + strconv.Itoa(v) + "\n"
		_, e := f.WriteString(s)
		errors.HandleError(e)
		linesAdded++
	}

	_, e2 := f.WriteString("------------------header end---------------------\n\n")
	errors.HandleError(e2)

	fmt.Println(" loop")	
	//write the prefix for every character in the original file
	fileContent := ""
	for _, line := range lines {
		runes := []rune(string(line))
		for i := 0; i < len(runes); i++ {
			r := runes[i]
			if unicode.IsSpace(r) {
				fileContent += string(r)
			} else {
				fileContent += prefixMap[r]
			}
		}
	}

	fmt.Println("file content")
	_, e3 := f.Write([]byte(fileContent))
	errors.HandleError(e3)
}
