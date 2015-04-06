package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Ansible Share"
	app.Usage = "ansible-share command [command options] path to folder"
	app.Commands = []cli.Command{
		{
			Name:   "push",
			Usage:  "ansible-share push -t name .",
			Action: Push,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "tag, t",
					Value: "nginx",
					Usage: "name of role, default to folder name",
				},
			},
		},
		{
			Name:   "clone",
			Usage:  "ansible-share clone",
			Action: Clone,
		},
	}
	app.Run(os.Args)
}

func checkerror(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Push(c *cli.Context) {
	tag := c.String("tag")
	path := c.Args()[0]
	role := NewRole(tag)
	walker := FileWalker{}
	compressor := Tar{}
	httpTransport := HttpTransport{url: role.Url()}
	PushRole(path, walker, compressor, httpTransport)
}

func Clone(c *cli.Context) {
	tag := c.Args()[0]
	tar := Tar{}
	role := NewRole(tag)
	transport := HttpTransport{}
	file := transport.DownloadFile(role.Url())
	basePath := "./role" + role.name
	tar.Uncompress(file, basePath)
}
