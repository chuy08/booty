package booty

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	//"os"
)

type Templates []struct {
	Path       string                 `yaml:"path"`
	Executable bool                   `yaml:"executable,omitempty"`
	Values     map[string]interface{} `yaml:"values"`
}

type Input struct {
	Templates   Templates
	MainProcess []map[string]interface{} `yaml:"main-process"`
	CoProcesses []map[string]interface{} `yaml:"co-processes"`
}

func readConf(filename string) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	var result Input
	err = yaml.Unmarshal(buf, &result)
	if err != nil {
		fmt.Println(fmt.Errorf("in file %q: %v", filename, err))
	}

	//fmt.Println(result)
	//fmt.Println(result.Templates)

	parseTemplates(result.Templates)

	//fmt.Println(result)
	//fmt.Println(result["templates"])

	//for key := range result {
	//fmt.Println("key: ", key, "result: ", result[key])
	//switch vv := result[key].(type) {
	//case string:
	//	fmt.Println(key, "is a string", vv)
	//case float64:
	//	fmt.Println(key, "is a float64: ", vv)
	//case []interface{}:
	//	fmt.Println(key, "is an interface: ", vv)
	//	for k, v := range vv {
	//		fmt.Println("key: ", k, "value: ", v)
	//	}
	//default:
	//	fmt.Println(key, "is type is unknown: ", vv)
	//}
	//}
	//return parsed
}

func entry(cmd *cobra.Command, args []string) {
	//fmt.Println(cmd)
	//fmt.Println(args)

	fileName, _ := cmd.Flags().GetString("file")
	fmt.Printf("Filename: %s\n", fileName)

	//c := readConf(fileName)
	readConf(fileName)
	//fmt.Println(c)

	//t, err := template.New("todos").Parse("You have a task named \"{{ .Name}}\" with description: \"{{ .Description}}\"")
	//if err != nil {
	//	panic(err)
	//}
	//err = t.Execute(os.Stdout, c)
	//if err != nil {
	//	panic(err)
	//}
}
