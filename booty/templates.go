package booty

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
	//log "github.com/sirupsen/logrus"
)

func chomp(filename string) string {
	return strings.TrimSuffix(filename, ".erb")
}

func readPath(filename string, values map[string]interface{}) {
	//fmt.Println("Hi chuy from readPath: ", filename)

	t, err := template.ParseFiles(filename)
	if err != nil {
		fmt.Println(err)
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
	fmt.Println("Hi chuy from exectuable bool: ", b)
}

func parseTemplates(input Templates) {
	fmt.Println("Parsing config templates")
	//fmt.Println(input)

	for _, i := range input {
		//fmt.Println(i)
		readPath(i.Path, i.Values)

		//if i.Executable {
		//	executable(i.Executable)
		//}
	}

}
