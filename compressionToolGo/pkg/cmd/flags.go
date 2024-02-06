package cmd

import (
	"flag"
	"github.com/singlaanish56/compressionToolGo/pkg/algo"
)

var fileName string
func InitFlags(){
	const (
		defaultFile=""
		shorthand = "(shorthand)"
		usage = "Enter the file name to be compressed"
	)	

	flag.StringVar(&fileName,"file",defaultFile,usage)
	flag.StringVar(&fileName,"f",defaultFile,shorthand+usage)
}

func ParseTheFlags(){
	flag.Parse()

	if len(fileName) > 0 {
		algo.Compress(fileName)
	}
}