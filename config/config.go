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
//	namefix structs start
type NamefixConfig struct{
	Replacements []Replacement `json:"Replacements"`
}
type Replacement struct {
	Match       string `json:"Match"`
	Replacement string `json:"Replacement"`
}
//	namefix structs end
type MasterConfig struct {
	JamfConf JamfProConfig
	ZenConf ZenDeskConfig
	ADConf   ActiveDirectoryConfig
	NamefixConf NamefixConfig
}