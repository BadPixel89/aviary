package config

func LoadMasterConfig(path string) (MasterConfig, error) {
	return DefaultConf(), nil
}

func DefaultConf() MasterConfig {
	return MasterConfig{}
}

// this needs to create a dummy file with empty fields if no conf is present
// otherwise read in values
// conf is next to executable
// path passed in includes exe name, might be worth figuring out a way to reilably get
// just the folder path we are in
