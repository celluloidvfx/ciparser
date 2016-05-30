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

package ciparser

import (
	//"github.com/codegangsta/cli"
	"github.com/ghodss/yaml"
	"io/ioutil"
	//"os"
	//"reflect"
	//"strings"
)

type CiConfig struct {
	Name      string `json:"name"`
	Civersion string `json:"civersion"`
	Build     Build  `json:"build"`
}

type Build struct {
	Active    bool     `json:"active"`
	Output    string   `json:"output"`
	Language  string   `json:"language"`
	Goversion string   `json:"goversion"`
	Arch      string   `json:"arch"`
	Goos      string   `json:"platform"`
	Musl      bool     `json:"musl"`
	Deps      []string `json:"deps"`
}

func (c CiConfig) GetValueName(value string) string {
	switch {
	case value == "name":
		return c.Name
	case value == "civersion":
		return c.Civersion
	case value == "output":
		return c.Build.Output
	case value == "musl":
		if c.Build.Musl {
			return "true"
		} else {
			return "false"
		}
	case value == "active":
		if c.Build.Active {
			return "true"
		} else {
			return "false"
		}
	}
	return "N/A"
}

/*
func (c CiConfig) AppName() string {
	return c.Name
}

func (c CiConfig) CiVersion() string {
	return c.Civersion
}

func (c CiConfig) DoBuild() bool {
	return c.Build.Active
}

func (c CiConfig) UseMusl() bool {
	return c.Build.Musl
}
*/

//var B *Build
/*
func (c *CiConfig) BuildActive() bool {
	fmt.Println(c.Build.active)
	return c.Build.active
}

func (b Build) Active() bool {
	return b.active
}

func (b Build) Output() string {
	return b.output
}

func (b Build) Language() string {
	return b.language
}

func (b Build) Goversion() string {
	return b.goversion
}

func (b Build) Arch() string {
	return b.arch
}

func (b Build) Goos() string {
	return b.goos
}

func (b Build) Musl() bool {
	return b.musl
}

func (b Build) Deps() []string {
	return b.deps
}
*/
func ReadConfig(path string) (*CiConfig, error) {
	var ci *CiConfig
	b, err := ioutil.ReadFile(path)

	if err == nil {
		err = yaml.Unmarshal(b, &ci)
	}
	return ci, err
}

/*
func fetchFieldinYaml(path, key string) string {
	CI, err := ReadConfig(path)
	if err != nil {
		os.Exit(1)
	}
	B := CI.Build

	r := reflect.ValueOf(B)
	f := reflect.Indirect(r).FieldByName(strings.Title(key))

	return f.String()
}
*/
