package commands

type cliCommand struct {
	name        string
	description string
	Callback    func() error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "closes the Pokedex program",
			Callback:    commandExit,
		},
	}
}
