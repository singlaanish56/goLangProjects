package main

import (
	"github.com/singlaanish56/compressionToolGo/pkg/cmd"
	"fmt"
)

func main() {
	fmt.Println("this is it trying")
	cmd.InitFlags()
	cmd.ParseTheFlags()
}