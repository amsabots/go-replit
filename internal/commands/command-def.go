package commands


type CommandConfig struct {
  Next int
  Previous int
}

type CliCommand struct {
	Name        string
	Description string
	Callback    ReplCallBakFunc
}

type ReplCallBakFunc func() error


func GetCommands() map[string]CliCommand{
  commands := NewCommandConfig()
  var cliCommandMap = make(map[string]CliCommand)

  // Exit
  cliCommandMap["exit"] = CliCommand{
    Name: "exit",
    Description: "Exit the Pokedex",
    Callback: commands.commandExit,
  }
// Help
  cliCommandMap["help"] = CliCommand{
    Name: "help",
    Description: "Displays a help message",
    Callback: commands.helpCommand,
  }

  // Map
  cliCommandMap["map"] = CliCommand{
    Name: "map",
    Description: "Displays map info",
    Callback: commands.pokeAPI,
  }

  // Map
  cliCommandMap["mapb"] = CliCommand{
    Name: "mapb",
    Description: "Displays previous map info",
    Callback: commands.pokeAPIBack,
  }

  return cliCommandMap
}

func NewCommandConfig() *CommandConfig{
  return &CommandConfig{}
}