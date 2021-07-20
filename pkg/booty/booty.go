package booty

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"gopkg.in/yaml.v3"
)

type Things struct {
	Args      []string
	Cmd       *cobra.Command
	Verbosity string
}

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

//func ReadYamlInput(cmd *cobra.Command, args []string) {
func ReadYamlInput(p Things) {

	log.Info("Bootstraping with booty and away we go!!!")

	yamlInput, err := p.Cmd.Flags().GetString("file")
	if err != nil {
		log.Error(err)
	}
	log.Infof("Using %s as input", yamlInput)

	buf, err := ioutil.ReadFile(yamlInput)
	if err != nil {
		log.Error(err)
		return
	}

	var result Input
	err = yaml.Unmarshal(buf, &result)
	if err != nil {
		log.Error("in file %q: %v", yamlInput, err)
	}

	p.parseTemplates(result.Templates)
	log.Debugf("Template: %s", result.Templates)
	p.parseCommand(result.CoProcesses)
	log.Debugf("Co-process: %s", result.CoProcesses)
	p.parseCommand(result.MainProcess)
	log.Debugf("Main-process: %s", result.MainProcess)

	log.Info("And done")
}
