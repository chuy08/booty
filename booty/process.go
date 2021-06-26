package booty

import (
	"bytes"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func parseCommand(process Process) {
	var command bytes.Buffer

	for _, v := range process {
		command.WriteString(v.Command)
		command.WriteString(" ")
		for _, vv := range v.Args {
			command.WriteString(vv)
			command.WriteString(" ")
		}

		runCommand(command)
	}
}

func runCommand(command bytes.Buffer) {
	log.Info("Command: ", command.String())

	out, err := exec.Command(command.String()).Output()

	if err != nil {
		log.Fatal(err)
	}

	log.Info(string(out))

}
