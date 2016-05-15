package main

import (
	"os"

	cliApp "github.com/codegangsta/cli"
	"github.com/pythonandchips/ansible-share/cli"
)

func main() {
	app := cliApp.NewApp()
	app.Name = "Ansible Share"
	app.Usage = "ansible-share command [command options] path to folder"
	app.Commands = []cliApp.Command{}
	app.Commands = append(app.Commands, cli.Commands()...)
	app.Run(os.Args)
}
