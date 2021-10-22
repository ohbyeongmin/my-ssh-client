package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type ServerList struct {
	List []Server `yaml:"list"`
}

type Server struct {
	Name    string `yaml:"name"`
	Account string `yaml:"account"`
	Ip      string `yaml:"ip"`
	Port    int    `yaml:"port"`
	Key     string `yaml:"key"`
}

type PathType struct {
	Path string `yaml:"path"`
}

type PreRunType struct {
	PreRun PathType `yaml:"preRun"`
}

var serverViper *viper.Viper

func GetPreRun() string {
	var preRun PreRunType
	err := serverViper.Unmarshal(&preRun)
	if err != nil {
		fmt.Printf("fatal error config file: unmarshal server config, err: %s \n", err)
		os.Exit(1)
	}
	return preRun.PreRun.Path
}

func GetServerList() []string {
	serverList := getServerList()
	var serverNames []string
	for _, v := range serverList.List {
		serverNames = append(serverNames, v.Name)
	}
	return serverNames
}

func GetServerByIndex(index int) (Server, error) {
	serverList := getServerList()
	if len(serverList.List) < int(index) {
		return Server{}, errors.New("out of index")
	}

	return serverList.List[index-1], nil
}

func GetServerByName(name string) (Server, error) {
	serverList := getServerList()
	for _, v := range serverList.List {
		if strings.Contains(v.Name, name) {
			return v, nil
		}
	}
	return Server{}, errors.New("not exist server")
}

func getServerList() *ServerList {
	var list ServerList
	err := serverViper.Unmarshal(&list)
	if err != nil {
		fmt.Printf("fatal error config file: unmarshal server config, err: %s \n", err)
		os.Exit(1)
	}
	return &list
}

func SetFileConfig(file string) {
	serverViper = viper.New()
	serverViper.SetConfigFile(file)
	err := serverViper.ReadInConfig()
	if err != nil {
		fmt.Printf("fatal error config file: Server List, err: %s \n", err)
		os.Exit(1)
	}
}
