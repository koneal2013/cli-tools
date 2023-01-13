package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/koneal2013/cli-tools/todo"
)

func main() {
	add := flag.Bool("add", false, "Task to be included in the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	del := flag.Int("del", 0, "Item to be deleted")
	verbose := flag.Bool("verbose", false, "Enable verbose output, showing information like date/time")
	removeCompleted := flag.Bool("remove_completed", false, "Remove completed tasks when the 'list' command is executed")
	flag.Parse()

	var todoFileName = ".todo.json"
	todoFileNameEnvVar := os.Getenv("TODO_FILENAME")
	if todoFileNameEnvVar != "" {
		todoFileName = todoFileNameEnvVar
	}
	l := &todo.List{}
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	withoutCompleted := &todo.List{}
	switch {
	case *list:
		if *removeCompleted {
			for _, t := range *l {
				if !t.Done {
					*withoutCompleted = append(*withoutCompleted, t)
				}
			}
			if !*verbose {
				fmt.Print(withoutCompleted)
				os.Exit(0)
			}
		}
		if *verbose {
			if *removeCompleted {
				fmt.Print(withoutCompleted.StringVerbose())
				os.Exit(0)
			}
			fmt.Print(l.StringVerbose())
			os.Exit(0)
		}
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
	case *add:
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *del > 0:
		if err := l.Delete(*del); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, errors.New("invalid input provided"))
		os.Exit(1)
	}
}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", errors.New("task cannot be blank")
	}
	return s.Text(), nil
}
