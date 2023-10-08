/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

import "github.com/singlaanish56/goLangProjects/to-do-cmd/cmd"

func main() {

	filename:=""
	todoFilename := ".todo"
	currDir, err := os.Getwd()
	if err == nil{
		filename = filepath.Join(currDir,todoFilename)
		fmt.Println(filename)
	}

	cmd.Execute(filename)
}
