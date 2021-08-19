package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type Credentials struct {
	Host string
	Password  string
	User string
	Port int
}

func GetCredentials(profile string) Credentials {
	profilesPath, exists := os.LookupEnv("FILESY_PROFILE_PATH")
	if !exists {
		panic(fmt.Sprintf("FILESY_PROFILE_PATH environment variable is missing"))
	}
	cfg, err := ini.Load(profilesPath)
	if err != nil {
		panic(fmt.Sprintf("Fail to read file: %v", err))
	}
	section, err := cfg.GetSection(profile)
	if section == nil {
		panic(fmt.Sprintf("profile missing %v", err))
	}

	hasHost := section.HasKey("host")
	hasPassword := section.HasKey("password")
	hasUser := section.HasKey("user")
	hasPort := section.HasKey("port")

	if !hasHost || !hasPassword || !hasUser || !hasPort {
		panic(fmt.Sprintf("Some Credentials Missing"))
	}

	var port int = section.Key("port").MustInt()
	if port == 0  {
		panic(fmt.Sprintf("Cannot Parse Port value"))
	}
	return Credentials {
		Host: section.Key("host").Value(),
		Password: section.Key("password").Value(),
		User: section.Key("user").Value(),
		Port: port,
	}
}

