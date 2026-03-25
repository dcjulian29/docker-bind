/*
Copyright © 2026 Julian Easterling julian@julianscorner.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dcjulian29/go-toolbox/docker"
	"github.com/dcjulian29/go-toolbox/filesystem"
	"github.com/dcjulian29/go-toolbox/textformat"
)

var imageVersion string

func main() {

	args := filesystem.EnsureUnixPathArguments()

	if len(args) == 1 {
		if args[0] == "--image-version" {
			fmt.Println(imageVersion)
			os.Exit(0)
		}
	}

	prefix := "/usr/bin"
	binary := strings.ReplaceAll(filepath.Base(os.Args[0]), ".exe", "")

	switch binary {
	case "ddns-confgen":
		prefix = "/usr/sbin"
	case "rndc":
		prefix = "/usr/sbin"
	case "rndc-confgen":
		prefix = "/usr/sbin"
	case "tsig-keygen":
		prefix = "/usr/sbin"
	}

	binary = fmt.Sprintf("%s/%s", prefix, binary)
	data, work := docker.HostContainerVolume()

	opts := docker.ContainerOptions{
		AdditionalArgs:   strings.Join(args, " "),
		EntryPoint:       binary,
		Image:            "dcjulian29/bind",
		Interactive:      true,
		Tag:              imageVersion,
		Volumes:          []string{data},
		WorkingDirectory: work,
	}

	if _, err := docker.Run(opts); err != nil {
		fmt.Println(textformat.Fatal(err.Error()))
		os.Exit(1)
	}
}
