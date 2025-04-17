package config

import (
	"encoding/json"
	"os"
)

// should make a blank
var Config MasterConfig

func LoadMasterConfig(path string) error {
	//	point at file to overwrite blank above
	return nil
}

func SaveConfig(path string) error {
	//saves blank config because it is not loaded with data from file currnetly
	jsondata, err := json.MarshalIndent(Config, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path+"/aviary-config.json", jsondata, 0644)

	if err != nil {
		return err
	}
	return nil
}

// this needs to create a dummy file with empty fields if no conf is present
// otherwise read in values
// conf is next to executable
// path passed in includes exe name, might be worth figuring out a way to reilably get
// just the folder path we are in
