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

func readYamlInput(cmd *cobra.Command, args []string) {
	yamlInput, err := cmd.Flags().GetString("file")
	if err != nil {
		log.Error(err)
	}
	log.Infof("Using %s as input", yamlInput)

	buf, err := ioutil.ReadFile(yamlInput)
	if err != nil {
		log.Error(err)
	}

	var result Input
	err = yaml.Unmarshal(buf, &result)
	if err != nil {
		log.Error("in file %q: %v", yamlInput, err)
	}

	parseTemplates(result.Templates)
	parseCommand(result.CoProcesses)
	//parseCommand(result.MainProcess)
}