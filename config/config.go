package config

//	namefix structs start
type NamefixConfig struct {
	Replacements []Replacement `json:"Replacements"`
}
type Replacement struct {
	Match       string `json:"Match"`
	Replacement string `json:"Replacement"`
}

//	namefix structs end
type MasterConfig struct {
	NamefixConf NamefixConfig
}
