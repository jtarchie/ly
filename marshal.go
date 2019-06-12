package ly

import (
	"errors"
	"fmt"

	lua "github.com/yuin/gopher-lua"
	"gopkg.in/yaml.v2"
)

// Encode returns the YAML encoding of value.
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
	errNested      = errors.New("cannot encode recursively nested tables to YAML")
	errSparseArray = errors.New("cannot encode sparse array")
	errInvalidKeys = errors.New("cannot encode mixed or invalid key types")
)

func (j yamlValue) MarshalYAML() (interface{}, error) {
	var data interface{}

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
			return nil, errNested
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
					return nil, errInvalidKeys
				}
				if expectedKey != key {
					return nil, errSparseArray
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
					return nil, errInvalidKeys
				}
				obj[key.String()] = yamlValue{value, j.visited}
				key, value = converted.Next(key)
			}
			data = obj
		default:
			return nil, errInvalidKeys
		}
	default:
		return nil, fmt.Errorf(`cannot encode ` + lua.LValueType(j.LValue.Type()).String() + ` to YAML`)
	}
	return data, nil
}
