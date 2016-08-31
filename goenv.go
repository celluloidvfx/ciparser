/*
 * Johannes Amorosa, (C) 2016
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"errors"
	//"fmt"
	"os"
	"os/exec"
	"strings"
)

type Goenv struct {
	Path    string
	Version string
}

func getGoPath() (string, error) {
	p := os.Getenv("GOPATH")
	if p != "" {
		return p, nil
	}
	return "", errors.New("GOPATH not set")
}

func getInstalledGoVersion() (string, error) {
	var (
		v []byte
		e error
	)
	cmdName := "go"
	cmdArgs := []string{"version"}
	if v, e = exec.Command(cmdName, cmdArgs...).Output(); e != nil {
		return "", e
	}
	s := strings.Trim(strings.Split(string(v), " ")[2], "go")
	return strings.TrimSpace(string(s)), nil
}

func getGoBin() (string, error) {
	var (
		v []byte
		e error
	)
	cmdName := "which"
	cmdArgs := []string{"go"}
	if v, e = exec.Command(cmdName, cmdArgs...).Output(); e != nil {
		return "", e
	}
	return strings.TrimSpace(string(v)), nil
}

func installGoDeps(d string) error {
	var (
		//v []byte
		e error
	)

	cmdName := "go"
	cmdArgs := []string{"get", d}
	if _, e = exec.Command(cmdName, cmdArgs...).Output(); e != nil {
		return e
	} else {
		return nil
	}
}
