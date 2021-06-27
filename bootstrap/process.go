package booty

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func parseCommand(process Process) {
	log.Info("Process: ", process)
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

	//fmt.Println(do.String())
	if err := do.Run(); err != nil {
		log.Error(err)
	}
}
