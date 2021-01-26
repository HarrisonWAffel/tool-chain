package main

import (
	"fmt"
	"github.com/HarrisonWAffel/tool-chain/commands"
	"github.com/HarrisonWAffel/tool-chain/setup"
	"github.com/HarrisonWAffel/tool-chain/types"
	"os"
)

func main() {

	setup.EnsureBaseConfiguration()
	argCount := len(os.Args)
	if argCount <= 1 {
		printGeneralUsage()
		return
	}

	//more commands can be added here
	commands := []types.Command{
		commands.Profile,
		commands.Net,
	}

	for _, e := range commands {
		if os.Args[1] == e.Name {
			e.Handler()
			return
		}
	}

	printGeneralUsage()
}

func printGeneralUsage() {
	fmt.Println(setup.BoxMessage("       Command Line Tool-Chain.     \n                                    \n The below commands are available: " +
		" \nprofile: manage environment profiles\nnet: networking tools\nExample: 'tool-chain profile new'"))
}
