package booty

import (
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

func parseCommand(process Process) {
	for _, v := range process {
		var a []string

		a = append(a, v.Command)
		a = append(a, v.Args...)
		runCommand(v.Command, a)
	}
}

func runCommand(c string, args []string) {
	cmd, err := exec.LookPath(c)
	if err != nil {
		log.Error("Look Path ", err)
	}

	// cmd structure
	do := &exec.Cmd{
		Path:   cmd,
		Args:   args,
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	log.Info("Running: ", strings.Join(args, " "))
	if err := do.Run(); err != nil {
		log.Error(err)
	}
}
