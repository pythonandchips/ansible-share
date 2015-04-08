package main

import (
	"os"

	"github.com/1partcarbon/ansible_share/file"
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

func Push(c *cli.Context) {
	tag := c.String("tag")
	path := c.Args()[0]
	role := NewPushRole(tag)
	walker := file.FileWalker{}
	compressor := Tar{}
	httpTransport := HttpTransport{url: role.Url()}
	files := walker.ListFiles(path)
	tarfile := compressor.Compress(path, files)
	httpTransport.UploadFile(tarfile, "role", "")
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
