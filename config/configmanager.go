package config

func LoadConfig(path string) (Config, error) {
	return Config{
		JamfUrl: "www.google.com",
	}, nil
}

// this needs to create a dummy file with empty fields if no conf is present
// otherwise read in values
// conf is next to executable
// path passed in includes exe name, might be worth figuring out a way to reilably get
// just the folder path we are in
