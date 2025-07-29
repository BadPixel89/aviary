package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// should make a blank conf
var Config MasterConfig

const ConfigDirName string = "aviary"

const ConfigFileName string = "aviary-config.json"

var aviaryconfigdir string
var aviaryconfigfile string

func LoadMasterConfig(path string) error {
	userconfdir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("[error] could not find config directory")
		os.Exit(1)
	}
	aviaryconfigdir = filepath.Join(userconfdir, ConfigDirName)
	aviaryconfigfile = filepath.Join(aviaryconfigdir, ConfigFileName)

	_, err = os.Stat(aviaryconfigfile)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println(err.Error())
			return err
		}
		err = createConfig(aviaryconfigdir, ConfigFileName)
		if err != nil {
			fmt.Println(err.Error())
		}
		return nil
	}

	config, err := os.ReadFile(aviaryconfigfile)
	if err != nil {
		fmt.Println("errored to read config from: " + aviaryconfigfile)
		fmt.Println(err.Error())
		return nil
	}
	err = json.Unmarshal(config, &Config)
	if err != nil {
		fmt.Println("[error] unmarshalling config")
		return nil
	}
	return nil
}

// small wrapper func for external calls to save config
func SaveConfig() error {
	userconfdir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("[error] could not find config directory")
		os.Exit(1)
	}
	aviaryconfigdir = filepath.Join(userconfdir, ConfigDirName)
	aviaryconfigfile = filepath.Join(aviaryconfigdir, ConfigFileName)

	jsondata, err := json.MarshalIndent(Config, "", "    ")
	if err != nil {
		fmt.Println("[error] marshalling config")
		return err
	}
	err = os.WriteFile(aviaryconfigfile, jsondata, 0644)
	if err != nil {
		fmt.Println("[error] unable to write config file")
		return err
	}
	return nil
}

func createConfig(path string, file string) error {
	//	this should be blank if we are creating the conf
	jsondata, err := json.MarshalIndent(Config, "", "    ")
	if err != nil {
		fmt.Println("[error] marshalling config")
		return err
	}
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	_, err = os.Create(filepath.Join(path, file))
	if err != nil {
		fmt.Println("[error] unable to create config file")
		return err
	}
	err = os.WriteFile(filepath.Join(path, ConfigFileName), jsondata, 0644)
	if err != nil {
		fmt.Println("[error] unable to write config file")
		return err
	}
	return nil
}

// conf is intended to be in os.UserConfigDir()
//
// UserConfigDir returns the default root directory to use for user-specific configuration data.
// Users should create their own application-specific subdirectory within this one and use that.
//
// On Unix systems, it returns $XDG_CONFIG_HOME as specified by
// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html if non-empty
// else $HOME/.config.
// On Darwin, it returns $HOME/Library/Application Support.
//
// On Windows, it returns %AppData%.
//
// On Plan 9, it returns $home/lib.
//
// If the location cannot be determined (for example, $HOME is not defined), then it will return an error.
