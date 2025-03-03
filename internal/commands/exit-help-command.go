package commands

import (
	"fmt"
	"os"
)

func  commandExit(cfg *CommandConfig) error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
  }
  
  func  helpCommand(cfg *CommandConfig) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
    fmt.Println()
	for _, cmd := range GetCommands() {
	  command := cmd.Name
	  desc := cmd.Description
	  fmt.Printf("%s: %s\n", command, desc)
	}
	fmt.Println("")
	return nil
  }