///The registry will hold all the cliCommands and logic

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
    "encoding/json"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error

   
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
        callback: commandHelp,
    },
    "map": {
        name:        "map",
        description: "Displays 20 locations from the PokeWorld",
        callback: commandMap,
    },
    "mapb": {
        name:        "mapb",
        description: "Displays the previous 20 locations from the Pokeworld",
        callback: commandMapb,
    },
}

//Fuction command that will return only an error/nil
func commandExit(cfg *config) error{
        fmt.Println("Closing the Pokedex... Goodbye!")
        os.Exit(0)
        return fmt.Errorf("")
    }

func commandHelp(cfg *config) error{
    fmt.Println("Welcome to the Pokedex! \nUsage: \n Map \nhelp: Displays a help message\n exit: Exit the Pokedex")
    return nil
}


////////////Location struct and Get Request///////
type Map struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous *string   `json:"previous"`
	Results  []LocationArea
}
type LocationArea struct{
	Name string `json:"name"`
	URL  string `json:"url"`
	
}

type config struct {
    nextURL     string
    previousURL string
}

func commandMap(cfg *config) error{
    url := "https://pokeapi.co/api/v2/location-area/"
    if cfg.nextURL != ""{
        url = cfg.nextURL
    }

    res, err := http.Get(url)
    if err != nil{
        log.Fatal(err)
    }
    body, err := io.ReadAll(res.Body)
    res.Body.Close()

    if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

    //Unmarshaling the data and "copying" it to location
    location := Map{}
    error := json.Unmarshal(body, &location)
    if error != nil{
        fmt.Println(error)
    }

    cfg.nextURL = location.Next
    if location.Previous != nil{
        cfg.previousURL = *location.Previous
    }else{
        cfg.previousURL = ""
    }

    //Now once we have the data we will only pull what is needed with a for loop
    for _, locationArea := range location.Results{
        fmt.Println(locationArea.Name)
    }
    return nil
}

func commandMapb(cfg *config) error{
    if cfg.previousURL == ""{
        fmt.Println("you're on the first page")
        return nil
    }
    res, err := http.Get(cfg.previousURL)
    if err != nil{
        log.Fatal(err)
    }
    body, err := io.ReadAll(res.Body)
    res.Body.Close()

    if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

    //Unmarshaling the data and "copying" it to location
    location := Map{}
    error := json.Unmarshal(body, &location)
    if error != nil{
        fmt.Println(error)
    }

    //Now once we have the data we will only pull what is needed with a for loop
    for _, locationArea := range location.Results{
        fmt.Println(locationArea.Name)
    }
    return nil
}


