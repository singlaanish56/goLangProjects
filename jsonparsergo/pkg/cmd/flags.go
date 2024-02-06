package cmd

import (
	"flag"
	"github.com/singlaanish56/jsonparsergo/pkg/analysis"
)
var fileName string
var jsonString string

func InitFlags() {
	const (
		defaultFile = ""
		shorthand   = "(shorthand)"
		usageFile   = "Parses the Json if its valid or not"
	)

	flag.StringVar(&fileName,"file",defaultFile,usageFile)
	flag.StringVar(&fileName,"f",defaultFile,usageFile + shorthand)
	flag.StringVar(&jsonString,"s",defaultFile,usageFile)
}

func ParseTheFlags() {

	flag.Parse()

	if flag.NFlag()==0 && len(flag.Args()) == 1{
		fileName = flag.Arg(0)
		jsonString = flag.Arg(0)
	}

	if len(fileName) > 0 {
		analysis.ParseTheFile(fileName)
	}

	if len(jsonString) > 0{
		analysis.ParseTheString(jsonString)
	}
}