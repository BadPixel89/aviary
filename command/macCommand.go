package command

import "strings"

// use this API to get json info about a MAC
// https://api.maclookup.app/v2/macs/<MAC ADDRESS HERE>

var _ = RegisterCommand(MacCommand{})

type MacCommand struct{}

func (m MacCommand) Run(args []string) error {
	return nil

	// check args
	//		index 0 or -m -mac should be input
	//		type specifier should be checked and set as output
	//		no type specified = all types output
	//		-o = put on clipboard, default is to not add, config can overwrite this
	// make mac raw
	// 		len() > 12 cannot be raw
	//    		check what separator it contains
	//         	remove separator
	//     	    return raw
	// validate raw
	//		raw as input into creation func
	//		output desired formats to console
	//		put on clipboard if specified

}

func DetectFormat(inputMac string) string {
	if strings.Contains(inputMac, ":") {
		return "colon"
	}
	if strings.Contains(inputMac, "-") {
		return "colon"
	}
	if strings.Contains(inputMac, ".") {
		return "colon"
	}
	return "raw"
}

func (m MacCommand) Help() {

}

func (m MacCommand) Name() string {
	return "mac"
}
