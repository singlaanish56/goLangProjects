package wcgo

import(
	"flag"
)

var calBytes string
var calLines string
var calWords string
var calCharacters string

func InitFlags(){
	const (
		defaultFile = ""
		shorthand = "(shorthand)"
		usageBytes = "print the byte counts"
		usageLines = "print the newline counts"
		usageWords = "print the word counts"
		usageCharacters = "print the character counts"
	)


	flag.StringVar(&calBytes, "bytes", defaultFile,usageBytes)
	flag.StringVar(&calBytes, "c", defaultFile,usageBytes + shorthand)

	flag.StringVar(&calLines, "lines", defaultFile,usageLines)
	flag.StringVar(&calLines, "l", defaultFile,usageLines + shorthand)

	flag.StringVar(&calWords, "words", defaultFile,usageWords)
	flag.StringVar(&calWords, "w", defaultFile,usageWords + shorthand)

	flag.StringVar(&calCharacters, "chars", defaultFile,usageCharacters)
	flag.StringVar(&calCharacters, "m", defaultFile,usageCharacters + shorthand)
}

func ParseTheFlags(){

	flag.Parse()
	
	if flag.NFlag()==0 && len(flag.Args()) == 1{
		calBytes=flag.Arg(0)
		calLines=flag.Arg(0)
		calWords=flag.Arg(0)
	}

	if len(calBytes) > 0 {
		CalNoOfBytes(calBytes)
	}

	if len(calLines) > 0 {
		CalNoOfLines(calLines)
	}

	if len(calWords) > 0 {
		CallNoOfWords(calWords)
	}

	if len(calCharacters) > 0 {
		CalNoOfCharacter(calCharacters)
	}
	
}