package analysis

import (

	"github.com/singlaanish56/jsonparsergo/pkg/tokens"
)

type Lexer struct{
	input []rune
	currentChar rune
	startPosition int
	nextCharPosition int
}

func CreateLexer(input string) *Lexer{
	lexer := &Lexer{input:[]rune(input),nextCharPosition:0}
	lexer.nextChar()
	return lexer
}

func (lexer *Lexer) nextChar() {
	if lexer.nextCharPosition >= len(lexer.input){
		lexer.currentChar = 0
	}else{
		lexer.currentChar = lexer.input[lexer.nextCharPosition]
	}

	//fmt.Printf("%s\n",string(lexer.currentChar))
	lexer.startPosition = lexer.nextCharPosition
	lexer.nextCharPosition++
}

func (lexer *Lexer) GetToken() tokens.Token {

	var c = lexer.currentChar

	//if whitespace or other wise ignore
	if(c == ' ' || c=='\t' || c=='\n' || c=='\r'){
		lexer.nextChar()
	}

	//isNumber
	if(lexer.currentChar >= '0' && lexer.currentChar <= '9'){
		start := lexer.startPosition	

		for (lexer.currentChar >= '0' && lexer.currentChar <= '9'){

			lexer.nextChar()
		}

		return tokens.Token{string(lexer.input[start:lexer.startPosition]),tokens.Number,start,lexer.startPosition}
	}

	//isString
	if(lexer.currentChar=='"'){
		start := lexer.startPosition+1

		for{
			lexer.nextChar()

			if(lexer.currentChar == '"' || lexer.currentChar== 0){
				break;
			}
		}

		s := string(lexer.input[start:lexer.startPosition])
		startP := lexer.startPosition
		lexer.nextChar()
		return tokens.Token{s,tokens.String,start,startP}
	}

	//is true, false, null
	if(lexer.currentChar == 't' || lexer.currentChar == 'f' || lexer.currentChar == 'n'){
		start := lexer.startPosition

		for lexer.currentChar>='a' && lexer.currentChar<='z'{
			lexer.nextChar()
		}

		s := string(lexer.input[start:lexer.startPosition])

		if s=="true"{
			return tokens.Token{s,tokens.True,start,lexer.startPosition}
		}else if s=="false"{
			return tokens.Token{s,tokens.False,start,lexer.startPosition}
		}else if s=="null"{
			return tokens.Token{s,tokens.Null,start,lexer.startPosition}
		}else{
			lexer.currentChar =0
			t1 := tokens.Token{string(lexer.currentChar),tokens.Invalid, lexer.startPosition, lexer.nextCharPosition}
			lexer.nextChar()
			return t1
		}

	}

	c = lexer.currentChar
	var t tokens.Token
	//isSingleCharacter
	if(c=='{'){
		t = tokens.Token{string(c), tokens.LeftCurlyBR, lexer.startPosition, lexer.startPosition+1}
	}else if (c=='}'){
		t = tokens.Token{string(c), tokens.RightCurlyBR, lexer.startPosition, lexer.startPosition+1}
	}else if(c=='['){
		t = tokens.Token{string(c), tokens.LeftBoxBR, lexer.startPosition, lexer.startPosition+1}
	}else if(c==']'){
		t = tokens.Token{string(c), tokens.RightBoxBR, lexer.startPosition, lexer.startPosition+1}
	}else if(c==':'){
		t = tokens.Token{string(c), tokens.Colon, lexer.startPosition, lexer.startPosition+1}
	}else if(c==','){
		t =  tokens.Token{string(c), tokens.Comma, lexer.startPosition, lexer.startPosition+1}
	}else{
		lexer.currentChar =0
		t=tokens.Token{string(lexer.currentChar),tokens.Invalid, lexer.startPosition, lexer.nextCharPosition}
	}
	
	lexer.nextChar()
	return t
}
