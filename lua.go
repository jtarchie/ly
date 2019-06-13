package ly

import (
	lua "github.com/yuin/gopher-lua"
)

type null struct {
	lua.LValue
}

func (null) String() string       { return "null" }
func (null) Type() lua.LValueType { return lua.LTNil }

var _ lua.LValue = null{}

func NewState() *lua.LState {
	l := lua.NewState()
	l.SetGlobal("null", null{})
	return l
}
