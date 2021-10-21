package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type ServerList struct {
	List []Server `yaml:"list"`
}

type Server struct {
	Name string `yaml:"name"`
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
	Key  string `yaml:"key"`
}

type PathType struct {
	Path string `yaml:"path"`
}

type PreRunType struct {
	PreRun PathType `yaml:"preRun"`
}

func GetVersion() string {
	versionViper := viper.New()
	versionViper.SetConfigName("version")
	versionViper.SetConfigType("yaml")
	versionViper.AddConfigPath(".")
	versionViper.AddConfigPath("./config/")
	err := versionViper.ReadInConfig()
	if err != nil {
		fmt.Printf("fatal error config file: version, err: %s \n", err)
		os.Exit(1)
	}
	return versionViper.GetString("version")
}

func GetPreRun(file string) string {
	serverListViper := viper.New()
	serverListViper.SetConfigFile(file)
	serverListViper.AddConfigPath(".")
	serverListViper.AddConfigPath("./config/")
	err := serverListViper.ReadInConfig()
	if err != nil {
		fmt.Printf("fatal error config file: Server List, err: %s \n", err)
		os.Exit(1)
	}
	var preRun PreRunType
	err = serverListViper.Unmarshal(&preRun)
	if err != nil {
		fmt.Printf("fatal error config file: unmarshal server config, err: %s \n", err)
		os.Exit(1)
	}
	return preRun.PreRun.Path
}

func GetServerList(file string) []string {
	serverListViper := viper.New()
	serverListViper.SetConfigFile(file)
	serverListViper.AddConfigPath(".")
	serverListViper.AddConfigPath("./config/")
	err := serverListViper.ReadInConfig()
	if err != nil {
		fmt.Printf("fatal error config file: Server List, err: %s \n", err)
		os.Exit(1)
	}
	var serverList ServerList
	var serverNames []string
	err = serverListViper.Unmarshal(&serverList)
	if err != nil {
		fmt.Printf("fatal error config file: unmarshal server config, err: %s \n", err)
		os.Exit(1)
	}
	for _, v := range serverList.List {
		serverNames = append(serverNames, v.Name)
	}
	return serverNames
}

func GetServer(file string, index int32) Server {
	serverListViper := viper.New()
	serverListViper.SetConfigFile(file)
	serverListViper.AddConfigPath(".")
	serverListViper.AddConfigPath("./config/")
	err := serverListViper.ReadInConfig()
	if err != nil {
		fmt.Printf("fatal error config file: Server List, err: %s \n", err)
		os.Exit(1)
	}
	var serverList ServerList
	err = serverListViper.Unmarshal(&serverList)
	if err != nil {
		fmt.Printf("fatal error config file: unmarshal server config, err: %s \n", err)
		os.Exit(1)
	}
	return serverList.List[index-1]
}
