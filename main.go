package main

import (
	"os"

	cliApp "github.com/codegangsta/cli"
	"github.com/pythonandchips/ansible-share/cli"
	"github.com/pythonandchips/ansible-share/server"
)

func main() {
	app := cliApp.NewApp()
	app.Name = "Ansible Share"
	app.Usage = "ansible-share command [command options] path to folder"
	app.Commands = []cliApp.Command{}
	app.Commands = append(app.Commands, cli.Commands()...)
	app.Commands = append(app.Commands, server.Commands()...)
	app.Run(os.Args)
}
