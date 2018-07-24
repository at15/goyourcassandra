package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/chzyer/readline"
	"github.com/gocql/gocql"
)

var connectOnStart bool

var _ readline.AutoCompleter = (*cqlshCompleter)(nil)

type cqlshCompleter struct {
}

// Readline will pass the whole line and current offset to it
// Completer need to pass all the candidates, and how long they shared the same characters in line
// Example:
//   [go, git, git-shell, grep]
//   Do("g", 1) => ["o", "it", "it-shell", "rep"], 1
//   Do("gi", 2) => ["t", "t-shell"], 2
//   Do("git", 3) => ["", "-shell"], 3
func (c *cqlshCompleter) Do(line []rune, pos int) (newLine [][]rune, length int) {
	fmt.Println("line is", string(line))
	return [][]rune{
		{'o'},
		{'i', 't'},
	}, len(line)
}

func main() {
	fmt.Printf("gocqlsh version %s\n", version)
	if connectOnStart {
		cluster := gocql.NewCluster("localhost")
		cluster.Keyspace = "system"
		session, err := cluster.CreateSession()
		if err != nil {
			log.Fatal(err)
			return
		}
		defer session.Close()
	}

	completer := cqlshCompleter{}
	l, err := readline.NewEx(&readline.Config{
		Prompt:       ">",
		HistoryFile:  "/tmp/readline.tmp",
		AutoComplete: &completer,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		switch line {
		case ":help":
			fmt.Println("there is no help")
		default:
			fmt.Println("unknown", line)
		}
	}
}
