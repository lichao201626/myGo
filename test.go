package main

import (
	"fmt"
	"regexp"
	"strings"
)

var pattern = `[0-9]{4}-?/?[0-9]{2}-?/?[0-9]{2}`
var negate = false

var content = `20180809
	2018-08-12
	2018/08/23
	line 3`

func main() {
	regex, err := regexp.Compile(pattern)
	if err!=nil {
		fmt.Println("Failed to compild pattern:", err)
		return
	}
	
	lines := strings.Split(content, "\n")
	fmt.Printf("matches\tline\n")
	for _, line := range lines {
		matches := regex.MatchString(line)
		if negate {
			matches = !matches
		}
		fmt.Printf("%v\t%v\n", matches, line)
	}
	
	fmt.Println("Hello, playground")
}