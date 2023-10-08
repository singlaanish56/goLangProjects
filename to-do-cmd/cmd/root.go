/*
Copyright © 2023 Anish Singa <singlaanish56@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)


var file string
var ShowDone = false
var ShowUndone = false

var RootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A simple command line based to do app",
	Long: `  A simple command line based to do app,
  which just tries to track your day to day task,
	
  Because dont we all love our little adhd, anxiety riddled minds
  that run in multiple directions, unless we write and scratch something
  of tiny little lists.`,
}

func Execute(filename string) {
	file = filename
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var createTaskCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a task to the list",
	Long: "Add a task to the list",
	Args:cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		w, err := os.OpenFile(file,os.O_APPEND|os.O_CREATE|os.O_RDWR,0666)
		if err !=nil{
			panic(err)
		}
		defer w.Close()

		task:=strings.Join(args," ")
		_, err=fmt.Fprintln(w,task)
		if err !=nil{
			panic(err)
		}
	},
}

var showTaskCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the tasks in the list",
	Long: "Show the tasks in the list",
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open(file)
		if err!=nil{
			log.Fatal(err)
		}
		defer f.Close()
		br := bufio.NewReader(f)
		n :=1
		for {
			b,_,err:=br.ReadLine()
			if err !=nil{
				if err != io.EOF{
					fmt.Println("done")
				}
				break;
			}
			line :=string(b)
			if strings.HasPrefix(line,"-"){
				if ShowDone{
					fmt.Printf("--done-- %d : %s\n",n,strings.TrimSpace(line[1:]))
				}
			}else if ShowUndone{
				fmt.Printf("%d : %s\n",n,strings.TrimSpace(line))
			}
			n++
		}
	},
}

var deleteTaskCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete the tasks in the list",
	Long: "Delete the tasks in the list",
	Args:cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var ids[] int
		for _, arg := range args{
			id, err := strconv.Atoi(arg)
			if err!= nil{
				log.Fatal(err)
			}
			ids = append(ids,id)
		}

		w, err := os.Create(file+"_")
		if err != nil{
			log.Fatal(err)
		}

		f, err := os.Open(file)
		if err!=nil{
			log.Fatal(err)
		}

		br:=bufio.NewReader(f)
		for n:=1;;n++{
			b, _, err:= br.ReadLine()

			if err != nil {

				if err != io.EOF{
					fmt.Printf("done")
				}

				break
			}

			for _, i := range ids{
				if i != n{
					_, err:= fmt.Fprintf(w,"%s\n",string(b))
					if err!=nil{
						log.Fatal(err)
					}
					
				}
			}
			f.Close()
			w.Close()
			err = os.Remove(file)
			if err != nil{
				log.Fatal(err)
			}

			os.Rename(file+"_",file)
		}

	},
}

var doneTaskCmd = &cobra.Command{
	Use:   "tick [id]",
	Short: "Complete the tasks in the list",
	Long: "Complete the tasks in the list",
	Args:cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var ids[] int
		for _, arg := range args{
			id, err := strconv.Atoi(arg)
			if err!= nil{
				log.Fatal(err)
			}
			ids = append(ids,id)
		}

		w, err := os.Create(file+"_")
		if err != nil{
			log.Fatal(err)
		}

		f, err := os.Open(file)
		if err!=nil{
			log.Fatal(err)
		}

		br:=bufio.NewReader(f)
		for n:=1;;n++{
			b, _, err:= br.ReadLine()

			if err != nil {

				if err != io.EOF{
					fmt.Printf("done")
				}

				break
			}

			for _, i := range ids{
				line := strings.TrimSpace(string(b))
				present:=false
				if i==n{
					present=true
				}

				if present &&!strings.HasPrefix(line,"-"){
					_, err:= fmt.Fprintf(w,"-%s\n",line)
					if err != nil{
						log.Fatal(err)
					}
				}else{
					_, err:= fmt.Fprintf(w,"%s\n",line)
					if err!=nil{
						log.Fatal(err)
					}
				}
			}
		}
			f.Close()
			w.Close()
			err = os.Remove(file)
			if err != nil{
				log.Fatal(err)
			}

			os.Rename(file+"_",file)

	},
}

var undoneTaskCmd = &cobra.Command{
	Use:   "untick [id]",
	Short: "Untick the tasks in the list",
	Long: "Untick the tasks in the list",
	Args:cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var ids[] int
		for _, arg := range args{
			id, err := strconv.Atoi(arg)
			if err!= nil{
				log.Fatal(err)
			}
			ids = append(ids,id)
		}

		w, err := os.Create(file+"_")
		if err != nil{
			log.Fatal(err)
		}

		f, err := os.Open(file)
		if err!=nil{
			log.Fatal(err)
		}

		br:=bufio.NewReader(f)
		for n:=1;;n++{
			b, _, err:= br.ReadLine()

			if err != nil {

				if err != io.EOF{
					fmt.Printf("done")
				}

				break
			}

			for _, i := range ids{
				line := strings.TrimSpace(string(b))
				present:=false
				if i==n{
					present=true
				}

				if present && strings.HasPrefix(line,"-"){
					_, err:= fmt.Fprintf(w,"%s\n",strings.TrimPrefix(line,"-"))
					if err != nil{
						log.Fatal(err)
					}
				}else{
					_, err:= fmt.Fprintf(w,"%s\n",line)
					if err!=nil{
						log.Fatal(err)
					}
				}
			}
		}
			f.Close()
			w.Close()
			err = os.Remove(file)
			if err != nil{
				log.Fatal(err)
			}

			os.Rename(file+"_",file)

	},
}

var clearTaskCmd = &cobra.Command{
	Use:   "clear",
	Short: "Start with a clean state",
	Long: "Start with a clean state",
	Run: func(cmd *cobra.Command, args []string) {

		_ = os.Remove(file)

	},
}

func init() {
	// for the create task
	RootCmd.AddCommand(createTaskCmd)

	//for the show task
	RootCmd.AddCommand(showTaskCmd)
	showTaskCmd.Flags().BoolVarP(&ShowUndone,"undone","u",true,"Show only the undone tasks")
	showTaskCmd.Flags().BoolVarP(&ShowDone,"done","d",true,"Show only the done tasks")

	//for the delete task
	RootCmd.AddCommand(deleteTaskCmd)

	
	//for the done task
	RootCmd.AddCommand(doneTaskCmd)
	
	
	//for the undone task
	RootCmd.AddCommand(undoneTaskCmd)

	//for the clear task
	RootCmd.AddCommand(clearTaskCmd)
}


