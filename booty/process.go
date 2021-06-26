package booty

import (
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

func parseCommand(process Process) {
	log.Info("Process: ", process)
	for _, v := range process {
		a := strings.Join(v.Args, ", ")
		runCommand(v.Command, a)
	}
}

func runCommand(c, a string) {
	cmd, err := exec.LookPath(c)
	if err != nil {
		log.Error("Look Path ", err)
	}

	// cmd structure
	do := &exec.Cmd{
		Path:   cmd,
		Args:   []string{cmd, a},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	//fmt.Println(do.String())
	if err := do.Run(); err != nil {
		log.Error(err)
	}
}
