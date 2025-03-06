package commands

import (
	"time"

	httpclient "github.com/amsabots/go-replit.git/internal/http-client"
)


type CommandConfig struct {
  Next string
  Previous string
  HtttpClient httpclient.HTTPClient
}

type CliCommand struct {
	Name        string
	Description string
	Callback    ReplCallBakFunc
}

type ReplCallBakFunc func(cc *CommandConfig) error

func NewCommandConfig() *CommandConfig{
  commandConfig := new(CommandConfig)

  httpClient := httpclient.NewClient(5 * time.Second)
  commandConfig.HtttpClient = httpClient
  
  return commandConfig
}


func GetCommands() map[string]CliCommand{
  var cliCommandMap = make(map[string]CliCommand)

  // Exit
  cliCommandMap["exit"] = CliCommand{
    Name: "exit",
    Description: "Exit the Pokedex",
    Callback: commandExit,
  }
// Help
  cliCommandMap["help"] = CliCommand{
    Name: "help",
    Description: "Displays a help message",
    Callback: helpCommand,
  }

  // Map
  cliCommandMap["map"] = CliCommand{
    Name: "map",
    Description: "Displays map info",
    Callback: pokeAPI,
  }

  // Map
  cliCommandMap["mapb"] = CliCommand{
    Name: "mapb",
    Description: "Displays previous map info",
    Callback: pokeAPIBack,
  }

  return cliCommandMap
}