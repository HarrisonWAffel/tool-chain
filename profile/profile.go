package profile

import (
	"fmt"
	"github.com/HarrisonWAffel/tool-chain/config"
	"strings"
)

type profile struct {
	Name    string `json:"name"`
	Exports map[string]string
}

//a profile block looks like this
// ==============* profile name *==============
// export PROPERTY=VALUE
// ...
// ============================================
//

var cnf = config.ReadConfig()

func AddProfile(profileName string, properties map[string]string, isFirstProfile bool) error {
	c, e := ReadExportFile()
	if e != nil {

		return e
	}
	c = WriteProfileHeader(profileName, c)
	c = WriteProfileProperties(properties, c)
	c = WriteProfileFooter(c)
	e = SaveExportFile(c, isFirstProfile)
	if e != nil {
		fmt.Println(e)
		return e
	}
	return UpdateConfig(cnf)
}

func ActivateProfile(profileName string, c []string) {
	var err error
	if c == nil {
		c, err = ReadExportFile()
		if err != nil {
			panic(err)
		}
	}
	triggered := false
	var newset []string
	for _, t := range c {
		if strings.Contains(t, Header(profileName)) {
			triggered = true
		}
		if strings.Contains(t, Footer()) {
			triggered = false
		}

		if triggered {
			chars := strings.Split(t, "")
			if len(chars) > 0 {
				if chars[0] == "#" && (!strings.Contains(t, Footer()) && !strings.Contains(t, Header(profileName))) {
					newset = append(newset, strings.ReplaceAll(t, "#", ""))
				} else {
					newset = append(newset, t)
				}
			}

		}
		if !triggered {
			if strings.Contains(t, "#") {
				newset = append(newset, t)
			} else {
				newset = append(newset, "#"+t)
			}
		}
	}
	e := SaveExportFile(newset, false)
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
	e = UpdateConfig(cnf)
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func ListProfiles() []string {
	c := config.ReadConfig()
	return c.ProfileNames
}

//checks current config against exports file
//to flush any deleted profiles
func UpdateConfig(conf config.Config) error {
	s, e := ReadExportFile()
	if e != nil {
		panic(e)
	}
	nc := config.Config{
		conf.ExportPath,
		"",
		nil,
	}
	for _, e := range s {

		for _, k := range conf.ProfileNames {
			if strings.Contains(e, Header(k)) {
				nc.ProfileNames = append(nc.ProfileNames, k)
			}
		}
	}
	return config.UpdateFile(nc)
}
