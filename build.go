//+build mage

package main

import (
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Build mg.Namespace

var Aliases = map[string]interface{}{
	"build": Build.Win,
}

// Build linux binary
func (Build) Linux() error {
	envVar := make(map[string]string)
	envVar["GOOS"] = "linux"
	return sh.RunWithV(envVar, "go", "build", "-o", "bin/main", "./cmd")
}

// Build windows binary
func (Build) Win() error {
	envVar := make(map[string]string)
	sh.RunWithV(envVar, "go", "build", "-o", "bin/main.exe", "./cmd")
	return nil
}

// Build docker image and run as container
func Docker() error {
	sh.RunV("mage", strings.Split("build:linux", " ")...)
	sh.RunV("docker", strings.Split("build -t 192.168.0.20:5500/ken00535/fin-monitor .", " ")...)
	sh.RunV("docker", strings.Split("push 192.168.0.20:5500/ken00535/fin-monitor", " ")...)
	return nil
}
