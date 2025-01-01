package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	repl := &Repl{
		commands: getCommands(),
		config: &Config{},
	}

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		
		command, exists := repl.commands[commandName]
		if exists {
			err := command.callback(repl.config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name string
	description string
	callback func(cfg *Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"map": {
			name: "map",
			description: "Displays next location areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays previous location areas",
			callback: commandMapb,
		},
	}
}

type Config struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}

type Repl struct {
	commands map[string]cliCommand
	config *Config
}
