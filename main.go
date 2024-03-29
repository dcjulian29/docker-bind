/*
Copyright © 2023 Julian Easterling julian@julianscorner.com

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
	"os/exec"
	"path/filepath"
	"strings"
)

var apk string

func main() {
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
	args := os.Args[1:]
	pwd, _ := os.Getwd()
	data := strings.ReplaceAll(fmt.Sprintf("%s:/data", pwd), "\\", "/")
	docker := []string{
		"run",
		"--rm",
		"-it",
		"-v",
		data,
		fmt.Sprintf("dcjulian29/bind:%s", apk),
		binary,
	}

	if len(args) > 0 {
		docker = append(docker, args...)

		if args[0] == "--image-version" {
			fmt.Println(apk)
			os.Exit(0)
		}
	}

	cmd := exec.Command("docker", docker...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Printf("\n\033[1;31m%s: \033[1;33m%s\033[0m\n", "An error occurred", err)
		os.Exit(1)
	}

	os.Exit(0)
}
