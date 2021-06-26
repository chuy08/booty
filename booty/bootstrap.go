package booty

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
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
		log.Error(err)
	}

	var result Input
	err = yaml.Unmarshal(buf, &result)
	if err != nil {
		log.Error("in file %q: %v", filename, err)
	}

	//fmt.Println(result)
	//fmt.Println(result.Templates)

	parseTemplates(result.Templates)
}

func entry(cmd *cobra.Command, args []string) {
	//fmt.Println(cmd)
	//fmt.Println(args)

	fileName, _ := cmd.Flags().GetString("file")
	log.Info("Filename: ", fileName)

	readConf(fileName)
}
