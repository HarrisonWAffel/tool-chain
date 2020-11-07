package main

import (
	"flag"
	"fmt"
	"github.com/HarrisonWAffel/tool-chain/profile"
	"github.com/HarrisonWAffel/tool-chain/setup"
	"os"
)

type command struct {
	Name string
	handler func()
}

func main() {
	setup.EnsureBaseConfiguration()
	argCount := len(os.Args)
	commands := []command{ //more commands can be added here
		{Name:"profile", handler: func() {
			l := len(os.Args)
			if l < 3 {
				fmt.Println("usage: \n create a new profile: tool-chain profile new profileName \n activate a profile: tool-chain profile activate profileName")
				return
			}

			command := os.Args[2]
			switch command {
			case "new":
				setup.AddProfilePrompt(os.Args[3])

			case "activate":
				profile.ActivateProfile(os.Args[3], nil)

			default:
				fmt.Println("usage: \n create a new profile: tool-chain profile new profileName \n activate a profile: tool-chain profile activate profileName")
				return
			}
	}},
	{},
	{}}

	if argCount != 0 {
		for _, e := range commands {
			if len(os.Args) <= 1 {
				return
			}
			if os.Args[1] == e.Name {
				e.handler()
				break
			}
		}
		return
	}
	flag.PrintDefaults()
}
