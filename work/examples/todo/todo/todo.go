package todo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ToDo() {
	help()
	prompt([]string{})
}

func help() {
	fmt.Println(`Commands:
+ <String> - Add a TODO entry
- <Int>    - Delete the numbered entry
q          - Quit
h          - Print this help`)
}

func prompt(todos []string) IO {
	fmt.Println("\nCurrent TODO list:")
	_ = fmap(numberedToDo.put, numbered(todos))
	interpret(getLine(), todos)
	return nil
}

type numberedToDo struct {
	n int
	s string
}

func numbered(todos []string) []numberedToDo {
	var f func(n int, todos []string) []numberedToDo
	f = func(n int, todos []string) []numberedToDo {
		if len(todos) == 0 {
			return nil
		}
		return prepend(numberedToDo{n, todos[0]}, f(n+1, todos[1:]))
	}
	return f(0, todos)
}

type IO interface{}

func (n numberedToDo) put() IO {
	fmt.Printf("%v: %s\n", n.n, n.s)
	return nil
}

// getLine always returns a non-empty string (length > 0)
func getLine() string {
	fmt.Printf("\n=> ")
	l, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	if l := strings.ReplaceAll(l, "\n", ""); len(l) != 0 {
		return l
	}
	return getLine()
}

func interpret(command string, todos []string) IO {
	switch c, r := command[0], strings.TrimSpace(command[1:]); c {
	case '+':
		prompt(append(todos, r))
	case '-':
		n, _ := strconv.Atoi(r)
		if todos, ok := delete(n, todos); ok {
			return prompt(todos)
		}
		fmt.Println("No TODO entry matches the given number")
		prompt(todos)
	case 'q':
		return nil
	case 'h':
		help()
		prompt(todos)
	default:
		fmt.Printf("Invalid command: %q\n", command)
		prompt(todos)
	}
	return nil
}

func delete(n int, l []string) (newl []string, ok bool) {
	if n >= len(l) {
		return l, false
	}
	return append(l[:n], l[n+1:]...), true
}

func putTodos(todos []string) IO {
	var p func(int, []string)
	p = func(n int, todo []string) {
		if len(todo) <= n {
			return
		}
		fmt.Printf("%v: %s\n", n, todo[n])
		p(n+1, todo)
	}

	p(0, todos)
	return nil
}
