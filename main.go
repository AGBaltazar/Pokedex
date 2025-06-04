package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    cfg := &config{}
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Welcome to the Pokedex! \n")
    for {
        fmt.Print("Pokedex > ")
        var formattedInput string
        if scanner.Scan(){
            scannedInput := scanner.Text()
            formattedInput = cleanInput(scannedInput)
        }

        foundCommand, exists := commands[formattedInput]
        if exists{
            err := foundCommand.callback(cfg)
            if err != nil{
                fmt.Printf("Looks like there was an error : %v", err)
            } 
        }else if !exists{
                fmt.Println("Unknown command")
            }
    }

}

///This function will take the input and strip/split it based off the delimiter(whitespaces)
func cleanInput(text string) string{
    word := strings.Fields(text)
    
    lowerCase := strings.ToLower(word[0])

    return lowerCase
    }
