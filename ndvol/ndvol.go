package main

import (
	"github.com/qeas/nedge-docker-volume/ndvol/ndvolapi"
	"github.com/qeas/nedge-docker-volume/ndvol/ndvolcli"
	"os"
)

const (
	VERSION = "0.0.1"
)

var (
	client *ndvolapi.Client
)

func main() {
	ncli := ndvolcli.NewCli(VERSION)
	ncli.Run(os.Args)
}
