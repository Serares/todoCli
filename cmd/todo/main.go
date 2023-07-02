package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Serares/todo"
)

// hardcoded filename
var todoFileName = ".todo.json"

const fileNameEnv = "TODO_FILENAME"

func main() {

	if os.Getenv(fileNameEnv) != "" {
		todoFileName = os.Getenv(fileNameEnv)
	}
	// extract an address of an empty instance of todo.List
	l := &todo.List{}

	// insert a custom message for the usage of the program
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s CLI tool\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "HINT: Tasks prefixed with an X are completed\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage information:\n")
		flag.PrintDefaults()
	}
	task := flag.String("task", "", "Add a task to the list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Give an id of the item to be completed")
	fileName := flag.String("filename", "", "Define the filename where the tasks will be persisted")
	flag.Parse()
	// write to standard error if an error occurs and exit wih an exit code
	if *fileName != "" {
		todoFileName = *fileName
	}
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		// we implemented the Stringer interface
		// so we can just print the list directly
		fmt.Print(l)
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		l.Add(*task)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// invalid argument provided
		fmt.Fprintln(os.Stderr, "invalid option provided")
		os.Exit(1)
	}
}
