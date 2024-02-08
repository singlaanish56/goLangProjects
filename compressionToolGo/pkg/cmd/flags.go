package cmd

import (
	"flag"
	"github.com/singlaanish56/compressionToolGo/pkg/algo"
)

var inputFileName string
var ouputFileName string
func InitFlags(){
	const (
		defaultInputFile=""
		shorthand = "(shorthand)"
		inputFileusage = "Enter the file name to be compressed"
		outputFileusage = "Enter the file name for the output"
		defaultOutputFile = "result.txt"
	)	

	flag.StringVar(&inputFileName,"input",defaultInputFile,inputFileusage)
	flag.StringVar(&inputFileName,"f",defaultInputFile,shorthand+inputFileusage)
	flag.StringVar(&ouputFileName,"output",defaultOutputFile,outputFileusage)
	flag.StringVar(&ouputFileName,"o",defaultOutputFile,shorthand+outputFileusage)
}

func ParseTheFlags(){
	flag.Parse()

	if len(inputFileName) > 0 {
		algo.Compress(inputFileName, ouputFileName)
	}
}