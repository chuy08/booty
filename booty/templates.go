package booty

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

func chomp(filename string) string {
	return strings.TrimSuffix(filename, ".erb")
}

func readPath(filename string) {
	fmt.Println("Hi chuy from readPath: ", filename)

	n := map[string]string{"foo": "foobar", "bar": "barfoo"}

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

	err = t.Execute(f, n)
	if err != nil {
		log.Print("execute: ", err)
		return
	}

	f.Close()
}

func executable(b bool) {
	fmt.Println("Hi chuy from exectuable bool: ", b)
}

func parseValues(v map[string]interface{}) {
	fmt.Println("Hi chuy from parseValues", v)
}

func parseTemplates(input Templates) {
	fmt.Println("Parsing config templates")
	//fmt.Println(input)

	for _, i := range input {
		//fmt.Println(i)
		//parseValues(i.Values)
		readPath(i.Path)

		//if i.Executable {
		//	executable(i.Executable)
		//}
	}

}
