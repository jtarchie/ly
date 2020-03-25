package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/jtarchie/ly"
)

type options struct {
	ConfigFile string `long:"config" short:"c" description:"the main lua file to execute" required:"true"`
	Format     string `long:"format" short:"f" description:"the output format of the payload (yaml - default, json)"`
}

func main() {
	l := ly.NewState()
	defer l.Close()

	if len(os.Args) < 2 || os.Args[1] == "" {
		log.Print("no script was defined")
		os.Exit(1)
	}

	var opts options
	_, err := flags.Parse(&opts)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	startTop := l.GetTop()

	filename := opts.ConfigFile
	if err := l.DoFile(filename); err != nil {
		log.Printf("could not run script: %s", err)
		os.Exit(1)
	}

	numReturnedValues := l.GetTop() - startTop

	for index := -numReturnedValues; index <= -1; index++ {
		if table := l.ToTable(index); table == nil {
			log.Printf("the last return value of the script must be a table")
			os.Exit(1)
		}

		table := l.ToTable(index)
		var payload []byte
		switch opts.Format {
		case "json":
			payload, err = ly.JSONMarshal(table)
			if err != nil {
				log.Panicf("marshaling into format %s: %s", opts.Format, err)
			}
			fmt.Println(string(payload))
		default:
			payload, err = ly.YAMLMarshal(table)
			if err != nil {
				log.Panicf("marshaling into format %s: %s", opts.Format, err)
			}
			fmt.Printf("---\n%s", string(payload))
		}
	}
}
