package setup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/HarrisonWAffel/tool-chain/config"
	"github.com/HarrisonWAffel/tool-chain/profile"
	"io/ioutil"
	"os"
	"strings"
)

//EnsureBaseConfiguration checks that the config.json file exists,
//and whether or not the tool-chain has been configured. If the tool-chain
//has not been configured, a configuration prompt will occur.
func EnsureBaseConfiguration() {
	needsConfiguration, err := ConfigIsMissing()
	if err != nil {
		panic(err)
	}
	if needsConfiguration {
		Prompt()
	}
}


//Prompt asks the user for the absolute path of their
// zshrc/bashrc file
func Prompt() {
	fmt.Print(BoxMessage("Tool-Chain Setup"))
	fmt.Println("\n Step 1. add exports file to zshrc. exports file will be generated at: " + config.ReadConfig().ExportPath)
	wd, e := os.Getwd()
	if e != nil {
		panic(e)
	}
	nc := config.Config{
		ExportPath:        wd+"/exports",
		CurrentAWSProfile: "",
		ProfileNames:      nil,
	}

	e2 := config.UpdateFile(nc)
	if e2 != nil {
		panic(e2)
	}
	pth := ReadFromUser("Please enter the absolute(!) path to your zshrc or bashrc file.", "Please enter an absolute path (No ~)")
	if pth == "" {
		return
	}

	hasConfig, e := RCHasConfig(pth)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	if !hasConfig {
		err := AddConfigToRC(pth)
		if err != nil {
			panic(err)
		}
	}

	name := ReadFromUser("Step 3. Setup an profile \n Please enter a profile name: ", "Please enter a profile name (type exit or press control c to exit)")
	if name == "" {
		return
	}
	fmt.Println("Please enter your desired profile exports. \nTo set these values type the following into your console. \n For each desired export type <exportName>=<exportValue> \n when you are finished type 'done' ")
	props := make(map[string]string)
	for {
		pro := ""
		fmt.Scanln(&pro)
		if pro == "done" {
			break
		}
		split := strings.Split(pro, "=")
		var k,v string
		k = split[0]
		v = split[1]
		props[k] = v
	}

	var first = false
	if len(config.ReadConfig().ProfileNames)==0 {
		first = true
	}

	mustWork(profile.AddProfile(name, props, first))
	curConf := config.ReadConfig()
	curConf.ProfileNames = append(curConf.ProfileNames, name)
	mustWork(config.UpdateFile(curConf))
	p, e := profile.ReadExportFile()
	if e != nil {
		panic(e)
	}


	profile.ActivateProfile(name, p)
	fmt.Println("\n" + BoxMessage("Setup Complete!"))
}

func AddProfilePrompt(name string) {
	fmt.Println("Please enter your desired profile exports. \nTo set these values type the following into your console. \n For each desired export type <exportName>=<exportValue> \n when you are finished type 'done' ")
	props := make(map[string]string)
	for {
		pro := ""
		fmt.Scanln(&pro)
		if pro == "done" {
			break
		}
		split := strings.Split(pro, "=")
		var k,v string
		k = split[0]
		v = split[1]
		props[k] = v
	}

	var first = false
	if len(config.ReadConfig().ProfileNames)==0 {
		first = true
	}

	mustWork(profile.AddProfile(name, props, first))
	curConf := config.ReadConfig()
	curConf.ProfileNames = append(curConf.ProfileNames, name)
	mustWork(config.UpdateFile(curConf))
	p, e := profile.ReadExportFile()
	if e != nil {
		panic(e)
	}
	fmt.Println("Would you like to activate this profile? (Y/N)")
	yn := ""
	fmt.Scanln(&yn)
	if strings.ToLower(yn) == "y" {
		profile.ActivateProfile(name, p)
		fmt.Println("\n" + BoxMessage("Profile Added and Activated!"))
		return
	}
	fmt.Println("\n" + BoxMessage("Profile Added!"))
}

func mustWork(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFromUser(prompt, rePrompt string) string {
	input := ""
	fmt.Println(prompt + " type 'exit' or press control c to exit ")
	for {
		fmt.Scanln(&input)
		if input == "" {
			fmt.Println(rePrompt)
			continue
		}
		if input == "exit" {
			return input
		}
		break
	}
	return input
}

//ConfigIsMissing ensures that config.json exists, creates it if necessary.
// a boolean is returned indicating if the config.json file exists, and no changes were made
func ConfigIsMissing() (bool, error) {
	if !checkForFile("config/config.json") {
		configFile, err := os.OpenFile("config/config.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			 return false, err
		}
		emptyConfig := config.Config{}
		j, e := json.Marshal(emptyConfig)
		if e != nil {
			return false, e
		}
		_, err = configFile.Write(j)
		if err != nil {
			return false, err
		}
		return true, nil
	}else{
		return false, nil
	}
}

func AddConfigToRC(rcFilePath string) error {
	f, e := os.OpenFile(rcFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if e != nil {
		panic(e)
	}

	_, e = f.WriteString("sh "+ config.ReadConfig().ExportPath+"\n")
	if e != nil {
		panic(e)
	}
	return e
}

func RCHasConfig(rcFilePath string) (bool,error) {
	l, err := ioutil.ReadFile(rcFilePath)
	if err != nil {
		return false, err
	}
	lines := strings.Split(string(l), "\n")
	added := fmt.Sprintf("sh "+config.ReadConfig().ExportPath)
	for _, e := range lines {
				if e == added {
					return true, nil
				}
	}
	return false, nil
}
//BoxMessage returns a string containing the given profile surrounded by a border comprised of *'s
func BoxMessage(text string) string {
	if len(text) == 0 {return ""}
	var buf []byte
	b := bytes.NewBuffer(buf)
	lines := strings.Split(text, "\n")
	largest := 0
	for _, e := range lines {
		chars := strings.Split(e, "")
		if len(chars) >= largest {
			largest = len(chars)
		}
	}
	for i := 0; i < largest + 5; i++ {b.Write([]byte("*"))}
	b.Write([]byte(" "))
	for i, e := range lines {
		chars := strings.Split(e, "")
		if i == len(lines) - 1 {
			b.Write([]byte(( "\n* " + e)))
		}else{
			b.Write([]byte(( "\n* " + e + "*")))
		}
		if len(chars) <= largest {
			for j := len(chars); j < largest; j++ {
				b.Write([]byte(" "))
			}
			b.Write([]byte(" *"))
		}
	}
	b.Write([]byte("\n"))
	for i := 0; i < largest + 5; i ++ {
		b.Write([]byte("*"))
	}
	return b.String()
}


//checkForFile checks if a file exists.
//it assumes the working directory is the
//project root.
func checkForFile(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return true
}
