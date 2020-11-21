package main

import (
	"fmt"
	config "github.com/HarrisonWAffel/tool-chain/config"
	mynet "github.com/HarrisonWAffel/tool-chain/net"
	"github.com/HarrisonWAffel/tool-chain/profile"
	"github.com/HarrisonWAffel/tool-chain/setup"
	"os"
	"strings"
)

type command struct {
	Name    string
	handler func()
}

func main() {
	setup.EnsureBaseConfiguration()
	argCount := len(os.Args)
	if argCount == 1 {
		fmt.Println(setup.BoxMessage("       Command Line Tool-Chain.     \n                                    \n The below commands are available: " +
			" \nprofile: manage environment profiles\nnet: networking tools\nExample: 'tool-chain profile new'"))
		return
	}
	commands := []command{ //more commands can be added here
		{Name: "profile", handler: func() {
			l := len(os.Args)
			if l <= 3 {
				printProfileUsage()
				return
			}

			command := os.Args[2]
			switch command {
			case "new":
				setup.AddProfilePrompt(os.Args[3])

			case "activate":
				profile.ActivateProfile(os.Args[3], nil)

			case "remove":
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
				printProfileUsage()
			}
		}},
		{Name: "net",
			handler: func() {
				l := len(os.Args)
				if l != 3 {
					printNetUsage()
					return
				}
				cmd := os.Args[2]
				switch cmd {
				case "scan":
					c := config.ReadConfig()
					if c.Subnet == "" {
						fmt.Println("Please enter subnet. do not leave a period at the end")
						sub := ""
						fmt.Scanln(&sub)
						h := mynet.GetAllHostsOnNetwork(sub)
						for k, v := range h {
							fmt.Printf("Host Name: %s IP Address: %s\n", k, v)
						}
						c.Subnet = sub
						config.UpdateFile(c)
					}
					h := mynet.GetAllHostsOnNetwork(config.ReadConfig().Subnet)
					for k, v := range h {
						fmt.Printf("Host Name: %s IP Address: %s\n", k, v)
					}

				default:
					printNetUsage()
				}
			}},
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
	}
}

func printNetUsage() {
	fmt.Println("usage: \n Find All Hosts on network: 'tool-chain net scan' ")
}

func printProfileUsage() {
	fmt.Println("usage: \n create a new profile: tool-chain profile new profileName " +
		"\n activate a profile: tool-chain profile activate profileName" +
		"\n list all profiles: tool-chain profile list")

	fmt.Println("\n")
	pro := config.ReadConfig().CurrentProfile
	fmt.Printf("Current Profile: %s\n", pro)
	x := profile.GetProfileExports(pro)
	fmt.Println("Current Exports:")
	for k, v := range x {
		if strings.Contains(k, "ACCESS") || strings.Contains(k, "SECRET") {
			c := strings.Split(v, "")
			shown := c[len(c)-4:]
			var end []string
			end = append(end, strings.Repeat("*", len(c)-4))
			end = append(end, strings.Join(shown, ""))
			fmt.Printf("%s = %s\n", k, strings.Join(end, ""))
		} else {
			fmt.Printf("%s = %s\n", k, v)
		}
	}
}
