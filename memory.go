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
	Mem           map[string]node
	size          int
	globalPointer *GlobalMemory
}

// Structure of the call stack
type callStack struct {
	size          int
	Stack         map[string]*funcMemory
	globalPointer *GlobalMemory
}

// Structure Defining the global memory
type GlobalMemory struct {
	size   int
	Memory map[string]node
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

// Initializer of Gloal Memory
func initGlobalMemory() *GlobalMemory {
	return &GlobalMemory{
		size:   0,
		Memory: map[string]node{},
	}
}

// Initializations of call stack
func initCallStack() (*callStack, *GlobalMemory) {
	globalPointer := initGlobalMemory()
	return &callStack{
		size:          0,
		Stack:         map[string]*funcMemory{},
		globalPointer: globalPointer,
	}, globalPointer
}

// Deletes stack frame from call stacks
func (CS *callStack) deleteStackFrame(name string) {
	// fmt.Println("Trying to delete " + name)
	delete(CS.Stack, name)
	CS.size--
}

// Derefrencer Function
func (CS *callStack) deref(ptr string, accesser string, i int) node {
	functionName, variableName, isCorrectSyntax := strings.Cut(ptr, "*")
	if !isCorrectSyntax {
		fmt.Println("\u001b[31mIncorrect pointer syntax at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")

		CS.deleteStackFrame(accesser)
		return node{}
	}
	funcMemory := CS.Stack[functionName]
	if funcMemory == nil {
		fmt.Println("\u001b[31mNil Pointer Derefrence at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
		CS.deleteStackFrame(accesser)
		return node{}
	}
	data := funcMemory.get(variableName)
	returnNode := node{
		Type: data.Type.int,
		Data: data.any,
	}
	return returnNode
}

// Initializations of the function memory
func (CS *callStack) initMemory(name string) *funcMemory {
	data := funcMemory{
		Mem:           map[string]node{},
		size:          0,
		globalPointer: CS.globalPointer,
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

	if strings.HasPrefix(name, "%") {
		fnMem.Mem[strings.Split(name, "%")[1]] = dataNode
		fnMem.size++
	} else if strings.HasPrefix(name, "$") {
		global := fnMem.globalPointer
		global.Memory[strings.Split(name, "$")[1]] = dataNode
		global.size++
	}
}

// Retrieves Value of a variable
func (fnMem *funcMemory) get(name string) varValue {
	var varVal varValue
	var valueAndType node
	if strings.HasPrefix(name, "%") {
		varName := strings.Split(name, "%")[1]
		valueAndType = fnMem.Mem[varName]
	} else {
		global := fnMem.globalPointer
		varName := strings.Split(name, "$")[1]
		valueAndType = global.Memory[varName]
	}
	emptyNode := node{}
	if valueAndType == emptyNode {
		fmt.Println("Undefined Variable")
	}
	if valueAndType.Type == 1 {
		varVal = varValue{Type{valueAndType.Type}, 0, string(valueAndType.Data.(string)), valueAndType.Data}
	} else if valueAndType.Type == 0 {

		varVal = varValue{Type{valueAndType.Type}, valueAndType.Data.(float64), "", valueAndType.Data}
	} else if valueAndType.Type == 2 {
		varVal = varValue{Type{valueAndType.Type}, 0, "", valueAndType.Data}

	} else if valueAndType.Type == 3 {
		varVal = varValue{Type{valueAndType.Type}, 0, "", valueAndType.Data}

	}
	return varVal
}
