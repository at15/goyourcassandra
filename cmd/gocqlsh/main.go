package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocql/gocql"
	"github.com/peterh/liner"
)

func main() {
	// TODO: print version etc.
	fmt.Println("gocqlsh version 0.1")
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "system"
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()
	line := liner.NewLiner()
	defer line.Close()
	line.SetMultiLineMode(true)
	line.SetCompleter(func(line string) []string {
		if strings.HasSuffix(line, "a") {
			return []string{"apple", "add", "ageis"}
		}
		return nil
	})
	// TODO: history
	for {
		select {
		// TODO: handle signals, though I think liner is also handling ctrl d etc.
		default:
			// TODO: should have username in the prompt
			l, err := line.Prompt("cqlsh> ")
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Printf("you typed %s\n", l)
		}
	}
}
