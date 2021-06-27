package booty

import (
	"html/template"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func chomp(filename string) string {
	return strings.TrimSuffix(filename, ".erb")
}

func readPath(filename string, values map[string]interface{}, e bool) {

	t, err := template.ParseFiles(filename)
	if err != nil {
		log.Error(err)
		return
	}

	configFile := chomp(filename)
	f, err := os.Create(chomp(configFile))
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = t.Execute(f, values)
	if err != nil {
		log.Print("execute: ", err)
		return
	}

	f.Close()

	if e {
		executable(configFile)
	}
}

func executable(configFile string) {
	log.Infof("Making %s executable", configFile)
	c := "chmod"
	a := []string{c, "+x", configFile}
	runCommand(c, a)
}

func parseTemplates(input Templates) {
	for _, i := range input {
		readPath(i.Path, i.Values, i.Executable)
	}
}
