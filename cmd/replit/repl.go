package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var clicommandMap = make(map[string]cliCommand)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func init(){
  clicommandMap["exit"] = cliCommand{
    name: "exit",
    description: "Exit the Pokedex",
    callback: commandExit,
  }
  clicommandMap["help"] = cliCommand{
    name: "help",
    description: "Displays a help message",
    callback: helpCommand,
  }
}



func cleanInput(text string) []string{
	output := strings.ToLower(text)
	s := strings.Fields(strings.TrimSpace(output))
	 return s
}


func startRepl(){
	scanner := bufio.NewScanner(os.Stdin)
    scannerLoop:
 for {
   fmt.Print("Pokedex > ")
   if scanner.Scan(){
    input := cleanInput(scanner.Text())
    if len(input) == 0 {continue}
    
    commandWord := input[0]
    callAction, exists := clicommandMap[commandWord]
    if !exists {
      fmt.Println("Unknown command")
      continue
    }
    err := callAction.callback()
    if err != nil {
      log.Println(fmt.Errorf("call Action failed: %w", err))
      continue
    }
   }
   err := scanner.Err(); if err != nil{
    log.Println(fmt.Errorf("scanner error: %w", err))
    break scannerLoop
   }
 }
}

func commandExit() error{
  fmt.Println("Closing the Pokedex... Goodbye!")
  os.Exit(0)
  return nil
}

func helpCommand() error {
  helpMessage := `Welcome to the Pokedex!
Usage:

`
  for _, cmd := range clicommandMap {
    command := cmd.name
    desc := cmd.description
    helpMessage += fmt.Sprintf("%s: %s\n", command, desc)
  }
  fmt.Println(helpMessage)
  return nil
}