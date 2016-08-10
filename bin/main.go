package main

import (
	"github.com/dictybase-docker/docker-machine-kvm"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(kvm.NewDriver("default", "path"))
}
