package main

import (
	"fmt"
	"strconv"
	"strings"
)

// A function to check if a slice contains an element
func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

/*
Given the tokens and the i to access the current line. It usses
function's memory struct pointer to access first and second operand
and executes the given fuction on them providing required first and
second arguement and sets the third operand (If present otherwise the _res var) to the
return of the function. pointer to the callstack is is used to delete the stack frame if an error is thrown
*/
func operate(data [][][]string, i int, operation func(a float64, b float64) float64, MemoryAddr *funcMemory, CS *callStack, functionName string) {
	Memory := *MemoryAddr
	firstOperandType := data[i][1][0]
	secondOperandType := data[i][2][0]
	thirdOperand := "%_res"
	thirdOperandType := "var"
	secondOperandInt, _ := strconv.Atoi(data[i][2][1])
	firstOperandInt, _ := strconv.Atoi(data[i][1][1])
	firstOperand := float64(firstOperandInt)
	secondOperand := float64(secondOperandInt)

	if len(data[i]) == 4 {
		thirdOperand = data[i][3][1]
		thirdOperandType = data[i][3][0]
	}
	if firstOperandType == "var" {
		firstOperand = Memory.get(data[i][1][1]).float64

	}
	if secondOperandType == "var" {
		secondOperand = Memory.get(data[i][2][1]).float64

	}
	if !((firstOperandType == "var" || firstOperandType == "immediate") && (secondOperandType == "var" || secondOperandType == "immediate")) {
		fmt.Println("\u001b[31mFirst Operand and Second Operand Must be an Variable or a Immediate at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
		CS.deleteStackFrame(functionName)

		return
	}
	if thirdOperandType != "var" {
		fmt.Println("\u001b[31mThird Operand Must be an Variable at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
		CS.deleteStackFrame(functionName)

		return
	}
	if secondOperandType == "var" && Memory.get(data[i][2][1]).Type.int != 0 {
		fmt.Println("\u001b[31mIf Second Operand is Variable they must be int at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
		CS.deleteStackFrame(functionName)

		return
	}
	if firstOperandType == "var" && Memory.get(data[i][1][1]).Type.int != 0 {
		fmt.Println("\u001b[31mIf First Operand is Variable they must be int at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
		CS.deleteStackFrame(functionName)

		return
	}
	res := operation(firstOperand, secondOperand)
	Memory.setVar(thirdOperand, 0, res)
}

// Changes the file extension (Required for saving the lexed file)
func changeExtension(file string, ext string) string {
	split := strings.Split(file, ".")
	toReturn := ""
	for i := 0; i < len(split); i++ {
		if i == (len(split) - 1) {
			toReturn += ext

		} else {
			toReturn += (split[i] + ".")

		}
	}
	return toReturn
}
