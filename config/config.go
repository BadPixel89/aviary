package config

type JamfProConfig struct {
	JamfURL string
}
type ZenDeskConfig struct {
	ZenDeskURL string
	userlist   map[string]int
}
type ActiveDirectoryConfig struct {
	LDAPURL    string
	Domain     string
	SearchRoot string
}
type MasterConfig struct {
	JamfConfig JamfProConfig
	ZenConfig  ZenDeskConfig
	ADConfig   ActiveDirectoryConfig
}
