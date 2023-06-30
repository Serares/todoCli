package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type ListAPI interface {
	Add(string)
	Complete(int) error
	Delete(int) error
}

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	newItem := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	// using the dereference notation
	// to point to the actual value that the pointer holds
	*l = append(*l, newItem)
}

func (l *List) Complete(i int) error {
	if i <= 0 || i > len(*l) {
		fmt.Printf("The item %d doesn't exist", i)
		return errors.New("item inexistent")
	}

	(*l)[i-1].Done = true
	(*l)[i-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(i int) error {
	if i <= 0 || i > len(*l) {
		return fmt.Errorf("item %d doesn't exist", i)
	}
	// the bellow notations excludes the last
	// index. So in the bellow case the :i-1 index will be excluded
	*l = append((*l)[:i-1], (*l)[i:]...)

	return nil
}

// provide a file name to save the existing list
func (l *List) Save(filename string) error {
	js, err := json.Marshal(*l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}
