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
  cc := NewCommandConfig()
  var cliCommandMap = make(map[string]CliCommand)

  cliCommandMap["exit"] = CliCommand{
    Name: "exit",
    Description: "Exit the Pokedex",
    Callback: cc.commandExit,
  }
  cliCommandMap["help"] = CliCommand{
    Name: "help",
    Description: "Displays a help message",
    Callback: cc.helpCommand,
  }
  return cliCommandMap
}

func NewCommandConfig() *CommandConfig{
  return &CommandConfig{}
}