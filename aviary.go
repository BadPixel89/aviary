package main

import (
	"aviary/command"
	"aviary/config"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//program always receives the folder it's contained in as the first argument?
	//not under linux, passes ./aviary
	if len(os.Args) == 1 {
		welcome()
		configpath, wderr := os.Getwd()
		if wderr != nil {
			fmt.Println("[error] working dir not found")
			fmt.Println(wderr.Error())
		}

		conferr := config.LoadMasterConfig(configpath)
		if conferr != nil {
			fmt.Println("error loading config\n" + conferr.Error())
		}
		saveerr := config.SaveConfig(configpath)
		if saveerr != nil {
			fmt.Println(saveerr.Error())
		}

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == "q" {
				os.Exit(0)
			}

			if scanner.Text() == "help" {
				command.List()
				continue
			}

			input := strings.Split(scanner.Text(), " ")

			err := command.Parse(input)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		panic("[fail] command loop escaped")
	}

	err := command.Parse(os.Args[1:])

	if err != nil {
		fmt.Println(err.Error())
	}
}

func welcome() {
	fmt.Println("welcome to aviary... name pending update")

}
