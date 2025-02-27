package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/amsabots/go-replit.git/internal/commands"
)



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
    callAction, exists := commands.GetCommands()[commandWord]
    if !exists {
      fmt.Println("Unknown command")
      continue
    }
    err := callAction.Callback()
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