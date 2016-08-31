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
	"io/ioutil"

	"github.com/ghodss/yaml"
)

type CiConfig struct {
	Name      string `json:"name"`
	Civersion string `json:"civersion"`
	Build     Build  `json:"build"`
}

type Build struct {
	Active    bool         `json:"active"`
	Output    string       `json:"output"`
	Language  string       `json:"language"`
	Goversion string       `json:"goversion"`
	Arch      string       `json:"arch"`
	Goos      string       `json:"platform"`
	Musl      bool         `json:"musl"`
	Deps      []string     `json:"deps"`
	Cvars     []Customvars `json:"customvars"`
	//Cvars     map[string]string `json:"customvars"`
}

type Customvars struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	Value string `json:"value"`
}

var GetTasks = []string{"name", "civersion", "active", "output", "language", "goversion", "musl"}

func ReadConfig(path string) (*CiConfig, error) {
	var ci *CiConfig
	b, err := ioutil.ReadFile(path)

	if err == nil {
		err = yaml.Unmarshal(b, &ci)
	}
	return ci, err
}

func (c CiConfig) GetValueName(value string) interface{} {
	switch {
	// Ciconfig
	case value == "name":
		return c.Name
	case value == "civersion":
		return c.Civersion
	// Build
	case value == "active":
		if c.Build.Active {
			return "true"
		} else {
			return "false"
		}
	case value == "output":
		return c.Build.Output
	case value == "language":
		return c.Build.Language
	case value == "goversion":
		return c.Build.Goversion
	case value == "platform":
		return c.Build.Goos
	case value == "musl":
		if c.Build.Musl {
			return "true"
		} else {
			return "false"
		}
	case value == "deps":
		return c.Build.Deps
	case value == "cvars":
		return c.Build.Cvars
	}

	r := "Usage: "
	for _, t := range GetTasks {
		r = r + t + " "
	}
	return r
	//return errors.New("No value provided")
}
