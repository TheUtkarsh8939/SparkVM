package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Sturcture defining a function
type Function struct {
	Tokens [][][]string
	args   []string
}

/*
Function to run a given string of code. It also requires fuction name
and call stack for memory operation and is it running a function for
stopping at the end keyword
*/
func run(data [][][]string, functionName string, CS callStack, isRunningAFunc bool) (map[string]node, bool) {
	//Initializes the local memory and pushes stack frame to stack
	Memory := CS.initMemory(functionName)
	//Next 2 variables are required for funtion declarations
	inFunction := false
	currentFunc := ""
	//Loops over all the lines
	i := 0
	for i < len(data) {
		//Ignores all the empty lines
		if len(data[i]) == 0 {
			continue
		}
		//Error handling if the first word is not an instruction
		if data[i][0][0] != "instruction" {
			fmt.Println("\u001b[31mFirst Word Must be an instruction at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
			CS.deleteStackFrame(functionName)

			return Memory.Mem, true
		}
		//Defines the current instruction
		instruction := data[i][0][1]
		//Next comment is Required for debbugging
		// fmt.Println("Executing " + fmt.Sprintf("%d", i))
		if instruction == "fun" {
			if data[i][1][0] != "function" {
				fmt.Println("\u001b[31mSecond operand of function instruction should be a function with the syntax <name>(<args>) at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
				CS.deleteStackFrame(functionName)

				return Memory.Mem, true
			}
			tmp := strings.Split(data[i][1][1], "(")
			currentFunc = "%" + tmp[0]
			args := strings.Split(strings.TrimSuffix(tmp[1], ")"), ",")
			if args[0] != "" {
				for i := 0; i < len(args); i++ {
					if !strings.HasPrefix(args[i], "%") {
						fmt.Println("\u001b[31mFunction arguements could only be variables at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
						CS.deleteStackFrame(functionName)

						return Memory.Mem, true
					}
				}
			}
			Memory.setVar(currentFunc, 2, Function{[][][]string{}, args})
			inFunction = true
			i++
			continue
		} else if instruction == "end" {
			inFunction = false
			if isRunningAFunc {
				CS.deleteStackFrame(functionName)
				return Memory.Mem, true

			}
		}
		// fmt.Println("Running", functionName, "line", i+1)
		if (inFunction && currentFunc == "%main") || isRunningAFunc {

			switch instruction {
			case "eql", "cmp":
				if instruction == "eql" {
					operate(data, i, func(a float64, b float64) float64 { return float64(toInt(int(a) == int(b))) }, Memory, &CS, functionName)

				} else {
					operate(data, i, func(a float64, b float64) float64 { return float64(toInt(int(a) > int(b))) }, Memory, &CS, functionName)

				}
			case "mod":
				operate(data, i, func(a float64, b float64) float64 { return float64(int(a) % int(b)) }, Memory, &CS, functionName)
			case "deref":
				if len(data[i]) < 3 {
					fmt.Println("\u001b[31mNot enough Operands at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				if data[i][1][0] != "var" {
					fmt.Println("\u001b[31mExpected Variable for the first operand at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				if data[i][2][0] != "var" {
					fmt.Println("\u001b[31mExpected Variable for second operand at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				varData := Memory.get(data[i][1][1])
				if varData.Type.int != 3 {
					fmt.Println("\u001b[31mExpected pointer variable for first operand at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				val := CS.deref(varData.any.(string), functionName, i)
				Memory.setVar(data[i][2][1], val.Type, val.Data)
			case "setptr":
				if len(data[i]) < 3 {
					fmt.Println("\u001b[31mNot enough Operands at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				if data[i][1][0] != "var" {
					fmt.Println("\u001b[31mExpected Variable for second operand at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				if data[i][2][0] != "var" {
					fmt.Println("\u001b[31mExpected Variable for second operand at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				x := functionName + "*" + data[i][1][1]
				Memory.setVar(data[i][2][1], 3, x)
			case "import":
				if len(data[i]) < 4 {
					fmt.Println("\u001b[31mNot enough Operands at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				firstOpd := data[i][1]
				secondOpd := data[i][2]
				thirdOpd := data[i][3]
				if firstOpd[0] != "string" {
					fmt.Println("\u001b[31mImport path must be a string at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				if secondOpd[0] != "keyword" || secondOpd[1] != "as" {
					fmt.Println("\u001b[31mSecond Operand Must be \"as\" at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				if thirdOpd[0] != "var" {
					fmt.Println("\u001b[31mImport name should be a variable at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				Memory.setVar(thirdOpd[1], 2, firstOpd[1])
			case "call":
				if len(data[i]) < 2 {
					fmt.Println("\u001b[31mNot enough Operands at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				firstOpd := data[i][1]
				if firstOpd[0] != "function" {
					fmt.Println("\u001b[31mFirst operand must be a variable and a function type variable at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				funcName, _, _ := strings.Cut(firstOpd[1], "(")
				funcName = "%" + funcName
				fmt.Println(funcName)
				funcData := Memory.get(funcName)
				if funcData.Type.int != 2 {
					fmt.Println("\u001b[31mFirst operand must be a variable and a function type variable at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				Tokens := funcData.any.(Function).Tokens
				_, err := run(Tokens, funcName, CS, true)
				if err {
					return Memory.Mem, true
				}
			case "halt":
				CS.deleteStackFrame(functionName)

				return Memory.Mem, true
			case "set":
				Type := 0

				if len(data[i]) < 3 {
					fmt.Println("\u001b[31mNot Enough Operands at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				if data[i][2][0] == "string" {
					Type = 1
					Memory.setVar(data[i][1][1], Type, data[i][2][1])
				} else if data[i][2][0] == "instruction" || data[i][2][0] == "keyword" {
					fmt.Println("\u001b[31mVariable value couldn't be an instruction or keyword at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				} else if data[i][2][0] == "immediate" {
					x, _ := strconv.ParseFloat(data[i][2][1], 64)
					Memory.setVar(data[i][1][1], Type, x)

				}
				if data[i][1][0] != "var" {
					fmt.Println("\u001b[31mFirst Operand Must be an Variable at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
			case "add":
				operate(data, i, func(a float64, b float64) float64 { return a + b }, Memory, &CS, functionName)
			case "sub":
				operate(data, i, func(a float64, b float64) float64 { return a - b }, Memory, &CS, functionName)
			case "mult":
				operate(data, i, func(a float64, b float64) float64 { return a * b }, Memory, &CS, functionName)
			case "div":
				operate(data, i, func(a float64, b float64) float64 { return a / b }, Memory, &CS, functionName)
			case "sleep":
				if len(data[i]) < 2 {
					fmt.Println("\u001b[31mNot Enough Operands at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				if _, err := strconv.Atoi(data[i][1][1]); data[i][1][0] != "immediate" || err != nil {
					fmt.Println("\u001b[31mSleep Time can't be anything but a number at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				conv, _ := strconv.Atoi(data[i][1][1])
				time.Sleep(time.Duration(conv) * time.Millisecond)
			case "show":
				if len(data[i]) < 2 {
					fmt.Println("\u001b[31mNot Enough Operands at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				val := Memory.get(data[i][1][1])
				fmt.Println(val.any)
			case "jump":
				if len(data[i]) < 2 {
					fmt.Println("\u001b[31mNot Enough Operands at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")
					CS.deleteStackFrame(functionName)

					return Memory.Mem, true
				}
				lineNum, lineNumType := data[i][1][1], data[i][1][0]
				if _, err := strconv.Atoi(lineNum); lineNumType != "immediate" || err != nil {
					fmt.Println("\u001b[31mLine number must be a number (It can't be a float or a negative number) at line " + fmt.Sprintf("%d", i+1) + "\u001b[38;2;255;255;255m")

				}
				lineNoPlus1, _ := strconv.Atoi(lineNum)
				i--
				i = lineNoPlus1 - 1
				continue
			}
			if currentFunc == "%main" {
				x := Memory.get(currentFunc).any.(Function)
				newVal := append(x.Tokens, data[i])
				x.Tokens = newVal
				Memory.setVar(currentFunc, 2, x)
			}
		} else {

			x := Memory.get(currentFunc).any.(Function)
			newVal := append(x.Tokens, data[i])
			x.Tokens = newVal
			Memory.setVar(currentFunc, 2, x)
		}
		i++
	}
	CS.deleteStackFrame(functionName)
	return Memory.Mem, false
}
