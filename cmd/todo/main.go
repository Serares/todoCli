package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Serares/todo"
)

// hardcoded filename
const todoFileName = ".todo.json"

func main() {
	// extract an address of an empty instance of todo.List
	l := &todo.List{}

	// write to standard error if an error occurs and exit wih an exit code
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		// for this case (only one argument exists,the prgoram name) just print the existing tasks
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		// populate the list
		l.Add(item)

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
