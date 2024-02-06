package errors

import(
	"log"
	"io"
)

func HandleError(err error){
	if err != nil {
		log.Fatal(err)
	}
} 


func HandleFileError(err error) bool{
	if err!=nil{
		if err == io.EOF {
			return true
		}
		log.Fatal(err)
	}

	return false
}
