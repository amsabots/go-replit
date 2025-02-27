package commands

import (
	"fmt"
	"os"
)

func (cc *CommandConfig) commandExit() error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
  }
  
  func (cc CommandConfig) helpCommand() error {
	helpMessage := `Welcome to the Pokedex!
  Usage:
  
  `
	for _, cmd := range GetCommands() {
	  command := cmd.Name
	  desc := cmd.Description
	  helpMessage += fmt.Sprintf("%s: %s\n", command, desc)
	}
	fmt.Println(helpMessage)
	return nil
  }