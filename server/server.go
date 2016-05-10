package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/gorilla/mux"
	"github.com/mitchellh/go-homedir"
)

var ErrSSLMissConfigure = errors.New("Public and Private key must be provided to use ssl mode")
var ErrPrivateKeyMissing = errors.New("Private key was not found at specified location.")
var ErrPublicKeyMissing = errors.New("Public key was not found at specified location.")
var directoryRoot string

func Commands() []cli.Command {
	return []cli.Command{
		{
			Name:  "web",
			Usage: "start web application",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "port, p",
					Value: "8080",
					Usage: "port to run ansible-share on",
				},
				cli.StringFlag{
					Name:  "home",
					Usage: "home folder to store roles",
				},
				cli.BoolFlag{
					Name:  "ssl",
					Usage: "turn ansible share ssl on. Must be used with --private_key and --public_key",
				},
				cli.StringFlag{
					Name:  "private_key",
					Usage: "location of private ssl key. This will turn on ssl",
				},
				cli.StringFlag{
					Name:  "public_key",
					Usage: "location of public ssl cert. This will turn on ssl",
				},
			},
			Action: start,
		},
	}
}

func start(c *cli.Context) {
	configuration, err := newConfiguration(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	createDirectory(configuration.home)
	r := mux.NewRouter()
	r.HandleFunc("/roles/{role}/{tag}", UploadRoleHandler).Methods("POST")
	r.HandleFunc("/roles/{role}/{tag}", DownloadRoleHandler).Methods("GET")
	r.HandleFunc("/_ping", Ping).Methods("GET")
	http.Handle("/", r)
	if configuration.ssl {
		err := http.ListenAndServeTLS(":443", configuration.privateKey, configuration.publicKey, nil)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		err := http.ListenAndServe(configuration.port, nil)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type configuration struct {
	ssl        bool
	privateKey string
	publicKey  string
	port       string
	home       string
}

func newConfiguration(c *cli.Context) (configuration, error) {
	config := configuration{
		port:       ":" + c.String("port"),
		home:       c.String("home"),
		ssl:        c.Bool("ssl"),
		publicKey:  c.String("public_key"),
		privateKey: c.String("private_key"),
	}
	if config.home == "" {
		dir, _ := homedir.Dir()
		config.home = filepath.Join(dir, "happer")
	}
	if config.privateKey != "" {
		if _, err := os.Stat(config.privateKey); os.IsNotExist(err) {
			return config, ErrPrivateKeyMissing
		}
	}
	if config.privateKey != "" {
		if _, err := os.Stat(config.privateKey); os.IsNotExist(err) {
			return config, ErrPublicKeyMissing
		}
	}
	return config, nil
}

func Ping(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "ping")
}
