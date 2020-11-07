package main

import (
	"fmt"
	config "github.com/HarrisonWAffel/tool-chain/config"
	"github.com/HarrisonWAffel/tool-chain/profile"
	"github.com/HarrisonWAffel/tool-chain/setup"
	"os"
)

type command struct {
	Name    string
	handler func()
}

func main() {
	setup.EnsureBaseConfiguration()
	argCount := len(os.Args)
	commands := []command{ //more commands can be added here
		{Name: "profile", handler: func() {
			l := len(os.Args)
			if l < 3 {
				fmt.Println("usage: \n create a new profile: tool-chain profile new profileName \n activate a profile: tool-chain profile activate profileName\n list all profiles: tool-chain profile list")
				return
			}

			command := os.Args[2]
			switch command {
			case "new":
				setup.AddProfilePrompt(os.Args[3])

			case "activate":
				profile.ActivateProfile(os.Args[3], nil)

			case "remove":
				curConf := config.ReadConfig()
				err := profile.UpdateConfig(curConf)
				if err != nil {
					panic(err)
				}

				for _, e := range config.ReadConfig().ProfileNames {
					if e == os.Args[3] {
						profile.DeleteProfile(os.Args[3], nil)
						fmt.Println("Profile " + os.Args[3] + " Has been deleted.")
						return
					}
				}
				fmt.Println("Profile " + os.Args[3] + " not found.")

			case "list":
				fmt.Print(profile.ListProfiles())

			default:
				fmt.Println("usage: \n create a new profile: tool-chain profile new profileName \n activate a profile: tool-chain profile activate profileName\n list all profiles: tool-chain profile list")
			}
		}},
		{},
		{}}

	if argCount != 0 {
		for _, e := range commands {
			if len(os.Args) <= 1 {
				break
			}
			if os.Args[1] == e.Name {
				e.handler()
				break
			}
		}
	} else {
		fmt.Println(setup.BoxMessage("       Command Line Tool-Chain.     \n                                    \n The below commands are available:  \nprofile: manage environment profiles"))
	}
}
