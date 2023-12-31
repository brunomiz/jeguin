package domain

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Job struct {
	token    string
	commands []string
}

func (j *Job) GetToken() string {
	return j.token
}

func (j *Job) SetToken(token string) {
	j.token = token
}

func (j *Job) GetExecution() func() {
	var name string
	var args []string
	var commands []*exec.Cmd
	for _, command := range j.commands {
		components := strings.Split(command, " ")
		for i, component := range components {
			if i == 0 {
				name = component
			} else {
				args = append(args, component)
			}
		}
		commands = append(commands, exec.Command(name, args...))
		args = nil
	}
	return func() {
		for _, command := range commands {
			err := command.Run()
			fmt.Println("executing:", command)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (j *Job) AddExecutionCommand(executionCommand string) {
	j.commands = append(j.commands, executionCommand)
}
