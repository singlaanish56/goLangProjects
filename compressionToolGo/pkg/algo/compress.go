package algo

import (
	"bufio"
	"os"

	"github.com/singlaanish56/compressionToolGo/pkg/errors"
)

var store []byte
func Compress(filename string, output string) {

	f, err := os.Open(filename)
	errors.HandleError(err)

	defer f.Close()

	reader := bufio.NewReader(f)

	for{
		line, _, err := reader.ReadLine()
		if errors.HandleFileError(err){
			break;
		}

		store = append(store, line...)
	}

	HuffmanCompress(store, output)
}