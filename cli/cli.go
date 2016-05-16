package cli

import (
	"fmt"
	"path/filepath"

	cliApp "github.com/codegangsta/cli"
	"github.com/pythonandchips/ansible-share/ansifile"
	"github.com/pythonandchips/ansible-share/file"
	"github.com/pythonandchips/ansible-share/role"
	"github.com/pythonandchips/ansible-share/storage"
)

func Commands() []cliApp.Command {
	return []cliApp.Command{
		{
			Name:   "push",
			Usage:  "ansible-share push -t {bucket_name}/{role name}:{tag} .",
			Action: push,
			Flags: []cliApp.Flag{
				cliApp.StringFlag{
					Name:  "tag, t",
					Usage: "name of role, default to folder name",
				},
			},
		},
		{
			Name:   "pull",
			Usage:  "ansible-share pull {bucket name}/{role name}:{tag}",
			Action: pull,
		},
		{
			Name:   "list",
			Usage:  "ansible-share list {bucket name}",
			Action: list,
		},
	}
}

func push(c *cliApp.Context) {
	tag := c.String("tag")
	path := c.Args()[0]
	role := role.NewPushRole(tag)
	walker := file.FileWalker{}
	compressor := Tar{}
	files := walker.ListFiles(path)
	tarfile := compressor.Compress(path, files)
	storage := storage.NewS3Storage(role.Host)
	storage.Put(role.Name, role.Version, tarfile)
}

func list(c *cliApp.Context) {
	host := c.Args()[0]
	storage := storage.NewS3Storage(host)
	roles, err := storage.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, role := range roles {
		fmt.Println(role)
	}
}

func pull(c *cliApp.Context) {
	if !c.Args().Present() {
		pullAll()
		return
	}
	tag := c.Args()[0]
	role := role.NewRole(tag)
	storage := storage.NewS3Storage(role.Host)
	version := role.Version
	if version == "latest" {
		latestVersion, err := storage.LatestVersion(role.Name)
		if err != nil {
			fmt.Println("Cannot pull role: ", err)
			return
		}
		role.Version = latestVersion
	}
	downloadRole(role)
}

func downloadRole(role role.Role) {
	storage := storage.NewS3Storage(role.Host)
	file, err := storage.Get(role.Name, role.Version)
	if err != nil {
		fmt.Println(err)
		return
	}
	basePath := filepath.Join("./roles", role.Name)
	tar := Tar{}
	tar.Uncompress(file, basePath)
	ansiWriteErr := ansifile.Write(role)
	if ansiWriteErr != nil {
		fmt.Println(err)
	}
}

func pullAll() {
	roles, err := ansifile.Read()
	if err != nil {
		fmt.Println("Error reading ansifile: ", err)
		return
	}
	for _, role := range roles {
		downloadRole(role)
	}
}
