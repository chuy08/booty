package booty

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"gopkg.in/yaml.v3"
	//"os"
)

type Process []struct {
	Command string   `yaml:"command"`
	Args    []string `yaml:"args"`
}

type Templates []struct {
	Path       string                 `yaml:"path"`
	Executable bool                   `yaml:"executable,omitempty"`
	Values     map[string]interface{} `yaml:"values"`
}

type Input struct {
	Templates   Templates `yaml:"templates"`
	MainProcess Process   `yaml:"main-process"`
	CoProcesses Process   `yaml:"co-processes"`
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

	log.Info("Templates: ", result.Templates)
	parseTemplates(result.Templates)
	log.Info("Co: ", result.CoProcesses)
	parseCommand(result.CoProcesses)
	log.Info("Main: ", result.MainProcess)
}

func entry(cmd *cobra.Command, args []string) {
	//fmt.Println(cmd)
	//fmt.Println(args)

	fileName, _ := cmd.Flags().GetString("file")
	extension, _ := cmd.Flags().GetString("extension")
	log.Info("Yaml Input: ", fileName)
	log.Info("Template Extension: ", extension)

	readConf(fileName)
}
