package main

import (
	"strings"
)

/*
Given a string of token it will return a slice conatining
2 elements first is its type and second is the original string
*/
func findtype(tofindtypeof string) []string {
	//List of all the instructions
	var instructionlist []string = []string{
		"import",
		"call",
		"halt",
		"def",
		"show",
		"add",
		"sub",
		"mult",
		"div",
		"and",
		"or",
		"xor",
		"not",
		"set",
		"jump",
		"jiz",
		"jin",
		"sleep",
		"fun",
		"end",
	}
	//List of all the keywords
	keywordList := []string{
		"as",
	}
	//An auxillary array (Return value)
	var typearr []string
	if contains(instructionlist, tofindtypeof) {
		typearr = []string{"instruction", tofindtypeof}
	} else if strings.HasPrefix(tofindtypeof, "%") {
		typearr = []string{"var", tofindtypeof}
	} else if strings.HasPrefix(tofindtypeof, "'") && strings.HasSuffix(tofindtypeof, "'") {
		typearr = []string{"string", tofindtypeof}
	} else if contains(keywordList, tofindtypeof) {
		typearr = []string{"keyword", tofindtypeof}

	} else if tofindtypeof == "" { //Ignore all the empty tokens(Space in the program)

	} else if !strings.HasPrefix(tofindtypeof, "(") && strings.HasSuffix(tofindtypeof, ")") {
		typearr = []string{"function", tofindtypeof}

	} else {
		typearr = []string{"immediate", tofindtypeof}
	}
	return typearr
}

/*
Takes the program in form of an array of byte and returns a
[][][]string datatype that is a lexed form of the program
*/
func tokenize(content []byte) [][][]string {
	//isComment is variable to check if the forwarding tokens are comments
	isComment := false
	//Converts the byte array to string
	code := string(content)
	//Splits the code into an array of lines
	linearr := strings.Split(code, "\r\n")
	//Temprory/Auxilary
	var tokenmatrix [][]string
	var tokentensor [][][]string
	var tempmatrix [][]string
	i := 0
	//Acutally tokenizing
	for i < len(linearr) {
		tokenmatrix = append(tokenmatrix, strings.Split(linearr[i], " "))
		i++
	}
	for x := 0; x < len(tokenmatrix); x++ {
		for y := 0; y < len(tokenmatrix[x]); y++ {
			if !isComment {
				if strings.HasPrefix(tokenmatrix[x][y], ";") {
					isComment = true
				}
				Type := findtype(tokenmatrix[x][y])
				if len(Type) != 0 && !isComment {
					tempmatrix = append(tempmatrix, Type)
				}
			}
		}
		if isComment {
			isComment = false
		}
		tokentensor = append(tokentensor, tempmatrix)
		tempmatrix = [][]string{}

	}
	return tokentensor
}
