package gocss

import (
	"fmt"
	"strconv"
)

var setProperty = documentElementStyle.Get("setProperty")

func CreateStringVar(initialValue string) (string, func(string)) {
	varName := "--" + generateId()
	sv := setProperty.Call("bind", documentElementStyle, varName)
	setValue := func(v string) {
		sv.Invoke(v)
	}
	setValue(initialValue)
	return fmt.Sprintf("var(%s)", varName), setValue
}

func CreateIntVar(initialValue int, units string) (string, func(int)) {
	varName := "--" + generateId()
	sv := setProperty.Call("bind", documentElementStyle, varName)
	setValue := func(v int) {
		sv.Invoke(strconv.Itoa(v) + units)
	}
	setValue(initialValue)
	return fmt.Sprintf("var(%s)", varName), setValue
}
