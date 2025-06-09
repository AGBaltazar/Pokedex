package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    cfg := &config{
        PokeCollection: make(map[string]Pokemon),
    }
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Welcome to the Pokedex! \n")
    for {
        fmt.Print("Pokedex > ")
        var formattedInput []string
        if scanner.Scan(){
            scannedInput := scanner.Text()
            formattedInput = cleanInput(scannedInput)
        }

        foundCommand, exists := commands[formattedInput[0]]
        if exists {
            if len(formattedInput) > 1 {
            strings.Join(formattedInput[1:], " ")
            }
            err := foundCommand.callback(cfg, formattedInput[1:])
            if err != nil{
                fmt.Printf("Looks like there was an error : %v", err)
            } 
        }else if !exists{
                fmt.Println("Unknown command")
            }
    }

}

///This function will take the input and strip/split it based off the delimiter(whitespaces)
func cleanInput(text string) []string{
    word := strings.Fields(text)
    var lowerCase []string
    for i := range word {
        
        lowerCase = append(lowerCase, strings.ToLower(word[i]))
    
}

    return lowerCase
    }
