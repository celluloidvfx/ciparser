// +build ignore

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
	"fmt"
	"github.com/codegangsta/cli"
	ci "gitlab/dev-op/ciparser"
	"os"
	"time"
)

var appVersion, appReleaseTag, appShortCommitID string

func main() {
	app := cli.NewApp()
	app.Name = "Celluloid YamlParser"
	app.Usage = "This Application reads cell-ci.yamls"
	app.Author = "Johannes Amorosa"
	app.Email = "johannesa@celluloid-vfx.com"
	app.Version = "XXX" + "\n" + mainVersion()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "p, path",
			Value: "cell-ci.yaml",
			Usage: "Path to input file",
		},
	}
	//a := os.Args[1:]

	app.Commands = []cli.Command{
		// One User query
		{
			Name:  "get",
			Usage: "Fetches a value",
			Before: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return errors.New("No value provided")
				}
				return nil
			},
			Action: func(c *cli.Context) error {
				cfg, err := ci.ReadConfig(c.String("path"))
				switch {
				case err == nil:
					a := c.Args().First()
					fmt.Println(cfg.GetValueName(a))
					return cli.NewExitError("", 0)
				default:
					return cli.NewExitError("", 1)
				}

			},
		},
		{
			Name:  "ldflags",
			Usage: "Create ldflags",
			Action: func(c *cli.Context) error {
				fmt.Println(ci.GenLDFlags(time.Now().UTC().Format(time.RFC3339)))
				return cli.NewExitError("", 0)
			},
		},
	}
	app.Run(os.Args)
}

func mainVersion() string {
	s := ""
	s = s + "Version: " + appVersion + "\n"
	s = s + "Release-Tag: " + appReleaseTag + "\n"
	s = s + "Commit-ID: " + appShortCommitID + "\n"
	return s
}

//fmt.Println(fetchFieldinYaml(a[0], a[1]))
/*
var CI *CiConfig
var B *Build

func ReadConfig(path string) (*CiConfig, error) {
	b, err := ioutil.ReadFile(path)

	err = yaml.Unmarshal(b, &CI)

	return CI, err
}

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

func main() {

	a := os.Args[1:]

	if len(os.Args) > 1 {
		fmt.Println(fetchFieldinYaml(a[0], a[1]))
		os.Exit(0)
	}

	os.Exit(1)
}
*/
