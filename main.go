package main

import (
	"fmt"
	"strings"
)

func main() {
    fmt.Println("Hello, World!")
}

///This function willtake the input and strip/split it based off the delimiter(whitespaces)
func cleanInput(text string) []string{
    word := strings.Fields(text)
    var outputSlice []string
    
    for i := 0; i < len(word); i++ {
        lowerCase := strings.ToLower(word[i])
        outputSlice = append(outputSlice, lowerCase)
       
    }
    return outputSlice
}