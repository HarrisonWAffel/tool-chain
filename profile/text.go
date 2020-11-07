package profile

import (
	"fmt"
	"github.com/HarrisonWAffel/tool-chain/config"
	"io/ioutil"
	"os"
	"strings"
)

func ReadExportFile() ([]string, error) {
	cnf = config.ReadConfig()
	b, e := ioutil.ReadFile(cnf.ExportPath)
	if e != nil {

		if !strings.Contains(e.Error(), "no such file") {
			fmt.Println("\n\n"+e.Error()+"\n\n")
			return nil, e
		}
		 _, e = os.Create(cnf.ExportPath)
		 if e != nil {
		 	return nil, e
		 }
	}
	b, e = ioutil.ReadFile(cnf.ExportPath)
	if e != nil {
		return nil, e
	}
	return strings.Split(string(b), "\n"), nil
}


func SaveExportFile(contents []string, withHeader bool) error {
	c := strings.Join(contents, "\n")
	f, e := os.OpenFile(cnf.ExportPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		return e
	}
	f.Truncate(0)
	f.Seek(0,0)
	if withHeader {
		_, e = f.WriteString("#!/bin/bash\n")
		if e != nil {
			return e
		}
	}
	_, e = f.WriteString(c)
	if e != nil {
		return e
	}
	f.Close()
	return nil
}


func WriteProfileHeader(profileName string, exportFile []string) []string {
	return append(exportFile, Header(profileName))
}

func WriteProfileFooter(exportFile []string) []string {
	return append(exportFile, Footer())
}

func WriteProfileProperties(props map[string]string, exportFile []string) []string {
	for k, v := range props {
		s := fmt.Sprintf("# export %s=%s", k, v)
		exportFile = append(exportFile, s)
	}
	return exportFile
}

func Header(profileName string) string {
	return fmt.Sprintf("#==============* %s *==============", profileName)
}

func Footer() string {
	return fmt.Sprint("#============================================")
}


