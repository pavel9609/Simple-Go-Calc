package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calc(exp string) float64 {
	//Checking if the supplied string is a number
	pos := 0
	result, err := strconv.ParseFloat(exp, 64)
	//Scope parsing
	openPos := strings.Index(exp, "(")
	//If we found open scope try
	if openPos != -1 {
		closePos := strings.LastIndex(exp, ")")
		if closePos != -1 {
			scopeString := exp[openPos : closePos+1]
			lenString := len(scopeString)
			result = calc(scopeString[1 : lenString-1])
			numberString := fmt.Sprintf("%f", result)
			exp = strings.ReplaceAll(exp, scopeString, numberString)
			fmt.Println("Test exp:", exp)
		} else {
			fmt.Println("Scope parsing error on: ", exp)
			os.Exit(1)
		}
	}

	//If the string is not a number, then we are looking for the operation
	if err != nil {
		//Addition
		pos = strings.Index(exp, "+")
		if pos != -1 {
			arr := strings.Split(exp, "+")
			result = calc(arr[0])
			for i := 1; i < len(arr); i++ {
				result = result + calc(arr[i])
			}
			return result
		}
		//Substraction
		pos = strings.Index(exp, "-")
		if pos != -1 {
			arr := strings.Split(exp, "-")
			result = calc(arr[0])
			for i := 1; i < len(arr); i++ {
				result = result - calc(arr[i])
			}
			return result
		}
		//Multipliaction
		pos = strings.Index(exp, "*")
		if pos != -1 {
			arr := strings.Split(exp, "*")
			result = calc(arr[0])
			for i := 1; i < len(arr); i++ {
				result = result * calc(arr[i])
			}
			return result
		}
		//Division
		pos = strings.Index(exp, "/")
		if pos != -1 {
			arr := strings.Split(exp, "/")
			result = calc(arr[0])
			for i := 1; i < len(arr); i++ {
				result = result / calc(arr[i])
			}
			return result
		}
		fmt.Println("Parse Error on: ", exp)
		os.Exit(1)
	}
	return result

}
func main() {
	var result float64
	//Use Scanner for read string with spaces
	in := bufio.NewScanner(os.Stdin)
	fmt.Print("Input expression: ")
	in.Scan()
	exp := in.Text()
	//Remove spaces
	exp = strings.ReplaceAll(exp, " ", "")
	fmt.Println(exp)
	result = calc(exp)
	fmt.Println("Result is: ", result)
}
