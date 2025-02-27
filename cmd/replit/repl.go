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
    user_command := cleanInput(scanner.Text())
    if len(user_command) == 0 {continue}
    
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