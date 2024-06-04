package main

import (
	"fmt"
	"strings"
)

// Defining structure of A single memory node storing a variables type and its value
type node struct {
	Type int
	Data interface{}
}

// Defining structure of A stack frame
type funcMemory struct {
	Mem  map[string]node
	size int
}

// Structure of the call stack
type callStack struct {
	size  int
	Stack map[string]*funcMemory
}

// Should be type Type int but I am too lazy to change everything. This just stores the type of the vvariable
type Type struct {
	int
}

// Structure defining how should get method return
type varValue struct {
	Type
	float64
	string
	any
}

// Initializations of call stack
func initCallStack() *callStack {
	return &callStack{
		size:  1,
		Stack: map[string]*funcMemory{},
	}
}

// Deletes stack frame from call stacks
func (CS *callStack) deleteStackFrame(name string) {
	// fmt.Println("Trying to delete " + name)
	delete(CS.Stack, name)
}

// Initializations of the function memory
func (CS *callStack) initMemory(name string) *funcMemory {
	data := funcMemory{
		Mem:  map[string]node{},
		size: 0,
	}
	CS.Stack[name] = &data
	CS.size++
	return &data
}

// Declares or changes value of a variable
func (fnMem *funcMemory) setVar(name string, Type int, value any) {
	dataNode := node{
		Type: Type,
		Data: value,
	}

	fnMem.Mem[strings.Split(name, "%")[1]] = dataNode
	fnMem.size++
}

// Retrieves Value of a variable
func (fnMem *funcMemory) get(name string) varValue {
	varName := strings.Split(name, "%")[1]
	valueAndType := fnMem.Mem[varName]
	emptyNode := node{}
	if valueAndType == emptyNode {
		fmt.Println("Undefined Variable")
	}
	var varVal varValue
	if valueAndType.Type == 1 {
		varVal = varValue{Type{valueAndType.Type}, 0, string(valueAndType.Data.(string)), valueAndType.Data}
	} else if valueAndType.Type == 0 {

		varVal = varValue{Type{valueAndType.Type}, valueAndType.Data.(float64), "", valueAndType.Data}
	} else if valueAndType.Type == 2 {
		varVal = varValue{Type{valueAndType.Type}, 0, "", valueAndType.Data}

	}
	return varVal
}
