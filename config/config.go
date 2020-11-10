package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

type Config struct {
	ExportPath        string `json:"export_path"` //the path to the *rc file (bashrc, zshrc)
	CurrentAWSProfile string `json:"current_aws_profile"`
	//CurrentGCPProfile string `json:"current_gcp_profile"` //todo
	ProfileNames   []string `json:"profile_names"`
	CurrentProfile string   `json:"current_profile"`
}

//UpdateFile updates the config.json file
func UpdateFile(config Config) error {
	err := os.Remove(RootDir() + "/config/config.json")
	if err != nil {
		return err
	}
	configFile, err := os.OpenFile(RootDir()+"/config/config.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	j, err := json.Marshal(config)
	if err != nil {
		return err
	}
	_, err = configFile.Write(j)
	if err != nil {
		return err
	}
	configFile.Close()
	return nil
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadConfig() Config {
	content, err := ioutil.ReadFile(RootDir() + "/config/config.json")
	if err != nil {
		return Config{}
	}
	c := Config{}
	json.Unmarshal(content, &c)
	return c
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
