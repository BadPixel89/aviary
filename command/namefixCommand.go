package command

import (
	"aviary/config"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var _ = RegisterCommand(NamefixCommand{})

type NamefixCommand struct{}

func (n NamefixCommand) Run(args []string) error {
	directory, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("namefix sees path as: " + directory)
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//not sure I like how many thing.things we have here
	previewRename(files, config.Config.NamefixConf.Replacements)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("[input] proceed with changes?:(y/n)")
	text, _ := reader.ReadString('\n')
	//account for newline of windows and unix
	if text == "y\n" || text == "y\r\n" {
		actuallyRename(files, directory, config.Config.NamefixConf.Replacements)
		os.Exit(0)
	}
	fmt.Println("[exit ] no changes made")
	return nil
}
func previewRename(files []os.DirEntry, replacements []config.Replacement) {
	fmt.Println("[info ] current file names")
	for _, file := range files {
		if file.Type().IsDir() {
			continue
		}
		fmt.Println("\t" + file.Name())
	}
	fmt.Println("[info ] names after renaming * indicates change")
	for _, file := range files {
		if file.Type().IsDir() {
			continue
		}
		rename := fixName(file.Name(), replacements)
		if rename == file.Name() {
			fmt.Println("\t" + rename)
		} else {
			fmt.Println("*\t" + rename)
		}

	}
}

// needs the file names and path so the full filepath can be used in old and new names
func actuallyRename(files []os.DirEntry, path string, replacements []config.Replacement) {
	for _, file := range files {
		if file.Type().IsDir() {
			continue
		}
		//more allocations but way more readable than one line full of filepath joins in brackets
		oldname := filepath.Join(path, file.Name())
		newname := filepath.Join(path, fixName(file.Name(), replacements))
		os.Rename(oldname, newname)
	}
}
func fixName(filename string, replacers []config.Replacement) string {
	filename = strings.ToLower(filename)
	filenamesplit := strings.Split(filename, ".")
	ext := ""
	if len(filenamesplit) > 1 {
		ext = "." + filenamesplit[len(filenamesplit)-1]
	}
	filename = strings.Replace(filename, ext, "", -1)

	for _, replace := range replacers {
		filename = strings.Replace(filename, replace.Match, replace.Replacement, -1)
	}

	//	keep regex here until we implement into config
	//	matches sequencs of 2 or more '.' '-' '_' or spaces
	// 	these patterns happen if many replacements match, "showname.1080p.webrip.rarbg.x264" becomes "showname...."
	removechars := regexp.MustCompile(`([.\s-_]){2,}`)
	remove := removechars.FindAllString(filename, -1)

	for _, substr := range remove {
		//match one to stop replace interfering with itself
		//each substring should occurr once if regex matched it
		filename = strings.Replace(filename, substr, "-", 1)
	}
	filename += ext
	//clean up in case last character is a -, -.mp4 looks weird
	filename = strings.Replace(filename, "-.", ".", -1)
	return filename
}

func (n NamefixCommand) Help() {
	fmt.Println("description:")
	fmt.Println("\trenames files according to the config file, currently towards my prefered convention")
	fmt.Println("\tnamefix operates on the current directory, move into a dir and run 'aviary namefix' to fix that directory")
	fmt.Println("\tspecifying a path is on the task list - see todolist.txt")
	fmt.Println("usage:")
	fmt.Println("\tnamefix")
	fmt.Println("no flags yet")
	fmt.Println("\t")
}
func (n NamefixCommand) Name() string {
	return "namefix"
}
