package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	lua "github.com/yuin/gopher-lua"
	"gopkg.in/yaml.v2"
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
	payload, err := Marshal(table)
	if err != nil {
		log.Panicf("marshaling yaml: %s", err)
	}
	fmt.Println(string(payload))
}

// Encode returns the JSON encoding of value.
func Marshal(value lua.LValue) ([]byte, error) {
	return yaml.Marshal(yamlValue{
		LValue:  value,
		visited: make(map[*lua.LTable]bool),
	})
}

var _ yaml.Marshaler = yamlValue{}

type yamlValue struct {
	lua.LValue
	visited map[*lua.LTable]bool
}

var (
	errNested      = errors.New("cannot encode recursively nested tables to JSON")
	errSparseArray = errors.New("cannot encode sparse array")
	errInvalidKeys = errors.New("cannot encode mixed or invalid key types")
)

func (j yamlValue) MarshalYAML() (data interface{}, err error) {
	switch converted := j.LValue.(type) {
	case lua.LBool:
		data = bool(converted)
	case lua.LNumber:
		data = float64(converted)
	case *lua.LNilType:
		data = nil
	case lua.LString:
		data = string(converted)
	case *lua.LTable:
		if j.visited[converted] {
			return nil, fmt.Errorf("circular reference")
		}
		j.visited[converted] = true

		key, value := converted.Next(lua.LNil)

		switch key.Type() {
		case lua.LTNil: // empty table
			data = []int{}
		case lua.LTNumber:
			arr := make([]yamlValue, 0, converted.Len())
			expectedKey := lua.LNumber(1)
			for key != lua.LNil {
				if key.Type() != lua.LTNumber {
					err = errInvalidKeys
					return
				}
				if expectedKey != key {
					err = errSparseArray
					return
				}
				arr = append(arr, yamlValue{value, j.visited})
				expectedKey++
				key, value = converted.Next(key)
			}
			data = arr
		case lua.LTString:
			obj := make(map[string]yamlValue)
			for key != lua.LNil {
				if key.Type() != lua.LTString {
					err = errInvalidKeys
					return
				}
				obj[key.String()] = yamlValue{value, j.visited}
				key, value = converted.Next(key)
			}
			data = obj
		default:
			err = errInvalidKeys
		}
	default:
		err = fmt.Errorf(`cannot encode ` + lua.LValueType(j.LValue.Type()).String() + ` to JSON`)
	}
	return
}
