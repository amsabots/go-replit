package main

import (
	"strings"
)


func cleanInput(text string) []string{
	output := strings.ToLower(text)
	s := strings.Fields(strings.TrimSpace(output))
	 return s
}