package wcgo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)


func CalNoOfBytes(calBytes string){
	f, err := os.Stat(calBytes)
	HandleError(err)
	fmt.Printf("%d\t", f.Size())
}

func CalNoOfLines(calLines string){

	// could also be done with ReadFile, but that loads the whole file at once into the splice
	// which could perform worse
	f, err := os.Open(calLines)
	HandleError(err)

	defer f.Close()
	reader := bufio.NewReader(f)
	lineCount := 0
	for {
		if _, _, err := reader.ReadLine(); HandleFileError(err){
			break
		}

		lineCount++;
	}

	fmt.Printf("%d\t", lineCount)
}

func CallNoOfWords(calWords string){
	f, err := os.Open(calWords)
	HandleError(err)

	defer f.Close()
	reader := bufio.NewReader(f)
	wordCount := 0
	for {
		line, _, err := reader.ReadLine()
		if HandleFileError(err){
			break
		}
		words := strings.Fields(string(line))
		wordCount += len(words)
	}

	fmt.Printf("%d\t", wordCount)
}

func CalNoOfCharacter(calCharacters string){
	f, err := os.Open(calCharacters)
	HandleError(err)

	defer f.Close()
	reader := bufio.NewReader(f)
	charCount := 0
	for {
		line, _, err := reader.ReadLine()
		if HandleFileError(err){
			break
		}
		charCount += utf8.RuneCountInString(string(line))
	}

	fmt.Printf("%d\t", charCount)
}
