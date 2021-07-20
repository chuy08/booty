package booty

import (
	"os"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
)

func chomp(filename string) string {
	return strings.TrimSuffix(filename, ".erb")
}

func (p Things) readPath(filename string, values map[string]interface{}, e bool) {

	t, err := template.ParseFiles(filename)
	if err != nil {
		log.Error(err)
		return
	}

	configFile := chomp(filename)
	f, err := os.Create(configFile)
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
		p.executable(configFile)
	}
}

func (p Things) executable(configFile string) {
	log.Infof("Making %s executable", configFile)
	c := "chmod"
	a := []string{c, "+x", configFile}
	p.runCommand(c, a)
}

func (p Things) parseTemplates(input Templates) {
	for _, i := range input {
		p.readPath(i.Path, i.Values, i.Executable)
	}
}
