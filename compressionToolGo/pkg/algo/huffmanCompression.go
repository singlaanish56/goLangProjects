package algo

import "fmt"

func HuffmanCompress(lines []byte) {

	freqMap := CalculateTheFrequency(lines)
	
	for k, v := range freqMap {
		fmt.Printf("%s:%v\n",string(k), v)
	}

}