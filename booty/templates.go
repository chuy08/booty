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

func readPath(filename string, values map[string]interface{}) {

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
}

func executable(b bool) {
	log.Info("Hi chuy from exectuable bool: ", b)
}

func parseTemplates(input Templates) {
	log.Info("Parsing config templates")

	for _, i := range input {
		readPath(i.Path, i.Values)

		//if i.Executable {
		//	executable(i.Executable)
		//}
	}
}
