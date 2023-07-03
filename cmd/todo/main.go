package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

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
	add := flag.Bool("add", false, "Add a task to the list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Give an id of the item to be completed")
	fileName := flag.String("filename", "", "Define the filename where the tasks will be persisted")
	delete := flag.Int("delete", 0, "Provide the task number to remove")
	verbose := flag.Bool("v", false, "This flag will display the date of each task")
	filterDoneTasks := flag.Bool("filter", false, "Use this flag to show only the tasks that are in progress")
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
		if *filterDoneTasks {
			l.FilterCompleted()
			// do the get at the end in case the filter was triggered
			defer l.Get(todoFileName)
		}

		if *verbose {
			fmt.Print(l.StringVerbose())
		}

		if !*verbose {
			fmt.Print(l)
		}
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		taskToAdd, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		l.Add(taskToAdd)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *delete > 0:
		if err := l.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	default:
		// invalid argument provided
		fmt.Fprintln(os.Stderr, "invalid option provided")
		os.Exit(1)
	}

}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	// use this to read the STDIN input
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}

	return s.Text(), nil
}
