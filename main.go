package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Starting Point of the Program
func main() {
	args := os.Args
	if len(args) >= 2 {
		//Running the VM with 'run' sub-command will execute the file given afterwards in the format 'sparkvm run <filename>'
		//*One must give the filename with its file extension
		if args[1] == "run" {
			content, err := os.ReadFile(args[2])
			if err != nil {
				fmt.Println(string("\033[31m"), err)
				os.Exit(1)
			}
			callStack, globalPointer := initCallStack()
			data := tokenize(content)
			atLastMem, _ := run(data, "%global", *callStack, false)
			fmt.Printf("%v\n", atLastMem)
			fmt.Printf("%v\n", *callStack)
			fmt.Printf("%v\n", *globalPointer)
		} else if args[1] == "lex" { //Running the VM with lex sub-command would print all the tokens generated by the lexer and save them to the file <filename>.spc FORMAT: sparkvm lex <filename>
			content, err := os.ReadFile(args[2])
			if err != nil {
				fmt.Println(string("\033[31m"), err)
				os.Exit(1)
			}
			data := tokenize(content)
			fmt.Printf("%v", data)
			jsond, _ := json.Marshal(data)
			os.WriteFile(changeExtension(args[2], "csp"), []byte(jsond), os.FileMode(60))
		}
		//TODO: Add a way to run already lexed .spc files to run directly without rerunning the lexer
	} else {
		for {
			//TODO: Add a command line way to execute code from terminal line by line like that of Javascript console
			fmt.Printf("%v", "Sparky VM>")
			var cmd []byte
			reader := bufio.NewWriter(os.Stdout)
			reader.Write(cmd)
			// tokens := tokenize([]byte(cmd))
			// atLastMem := run(tokens)

		}
	}
}
