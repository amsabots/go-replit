package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
    user_command := cleanInput(scanner.Text())
    if len(user_command) == 0 {continue}
    fmt.Printf("Your command was: %s\n", user_command[0])
   }
   err := scanner.Err(); if err != nil{
    log.Println(fmt.Errorf("scanner error: %w", err))
    break scannerLoop
   }
 }
}