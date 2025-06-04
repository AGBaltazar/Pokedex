package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Welcome to the Pokedex! \n")
    for {
        var formattedInput string
        if scanner.Scan(){
            scannedInput := scanner.Text()
            formattedInput = cleanInput(scannedInput)
        }

        foundCommand, exists := commands[formattedInput]
        if exists == true{
            err := foundCommand.callback()
            if err != nil{
                fmt.Printf("Looks like there was an error : %v", err)
            } 
        }else if exists == false{
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

