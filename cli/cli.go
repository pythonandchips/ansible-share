package cli

import (
	cliApp "github.com/codegangsta/cli"
	"github.com/pythonandchips/ansible-share/file"
)

func Commands() []cliApp.Command {
	return []cliApp.Command{
		{
			Name:   "push",
			Usage:  "ansible-share push -t name .",
			Action: Push,
			Flags: []cliApp.Flag{
				cliApp.StringFlag{
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
}

func Push(c *cliApp.Context) {
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

func Clone(c *cliApp.Context) {
	tag := c.Args()[0]
	tar := Tar{}
	role := NewRole(tag)
	transport := HttpTransport{}
	file := transport.DownloadFile(role.Url())
	basePath := "./role" + role.name
	tar.Uncompress(file, basePath)
}
