package main
import (
    "fmt"
    "os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error

   
}

var commands = map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
	"help": {
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
    },
}

//Fuction command that will return only an error
func commandExit() error{
        fmt.Println("Closing the Pokedex... Goodbye!")
        os.Exit(0)
        return fmt.Errorf("")
    }

func commandHelp() error{
    fmt.Println("Welcome to the Pokedex! \nUsage: \nhelp: Displays a help message\n exit: Exit the Pokedex")
    return nil
}
