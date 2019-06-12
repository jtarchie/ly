package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jtarchie/ly"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	l := lua.NewState()
	defer l.Close()

	if len(os.Args) < 2 || os.Args[1] == "" {
		log.Print("no script was defined")
		os.Exit(1)
	}

	filename := os.Args[1]
	if err := l.DoFile(filename); err != nil {
		log.Printf("could not run script: %s", err)
		os.Exit(1)
	}

	index := -1
	if table := l.ToTable(index); table == nil {
		log.Printf("the last return value of the script must be a table")
		os.Exit(1)
	}

	table := l.ToTable(index)
	payload, err := ly.Marshal(table)
	if err != nil {
		log.Panicf("marshaling yaml: %s", err)
	}
	fmt.Println(string(payload))
}
