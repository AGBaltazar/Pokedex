///The registry will hold all the cliCommands and logic

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
   
}

var commands = map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback: func(cfg *config, args []string) error {
            return commandExit(cfg)
        },
    },
	"help": {
        name:        "help",
        description: "Displays a help message",
        callback: func(cfg *config, args []string) error {
            return commandHelp(cfg)
        },
    },
    "map": {
        name:        "map",
        description: "Displays 20 locations from the PokeWorld",
        callback: func(cfg *config, args []string) error {
            return commandMap(cfg)
        },
    },
    "mapb": {
        name:        "mapb",
        description: "Displays the previous 20 locations from the Pokeworld",
        callback: func(cfg *config, args []string) error {
            return commandMapb(cfg)
        },
    },
    "explore":{
        name:       "explore",
        description: "Returns a list of pokemon within a given PokeTown",
        callback: func(cfg *config, args []string) error {
            if len(args) < 1 {
                return fmt.Errorf("please provide a location name")
            }
            return commandExplore(cfg, args[0])
        },
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
	Results []LocationArea
}
type LocationArea struct{
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationPokemon struct{
    Encounter []struct{}
    PokeEncounter []struct {
        Pokemon Pokemon `json:"pokemon"`
    } `json:"pokemon_encounters"`

}

type Pokemon struct{
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

func commandExplore(cfg *config, name string) error{
    baseUrl := "https://pokeapi.co/api/v2/location-area/"
    locationName := name
    Url := baseUrl + locationName

    res, err := http.Get(Url)
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
    pokemon := LocationPokemon{}
    error := json.Unmarshal(body, &pokemon)
    if error != nil{
        fmt.Println(error)
    }

    //Now once we have the data we will only pull what is needed with a for loop
    fmt.Printf("Exploring %v...\n" ,locationName)
    fmt.Println("Found Pokemon:\n")
    for _, locationPokemon := range pokemon.PokeEncounter{
        fmt.Printf("-%v\n", locationPokemon.Pokemon.Name)
    }
    return nil
}



