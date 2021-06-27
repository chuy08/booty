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

	f, err := os.Create(chomp(filename))
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
		executable()
	}
}

func executable() {
	log.Info("Hi chuy from exectuable")
}

func parseTemplates(input Templates) {
	log.Info("Parsing config templates")

	for _, i := range input {
		readPath(i.Path, i.Values, i.Executable)
	}
}
