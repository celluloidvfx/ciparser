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
	"os"

	"github.com/urfave/cli"
)

const apiversion = "1"

var appVersion, appReleaseTag, appShortCommitID, appBranch, path string

func main() {

	app := cli.NewApp()
	app.Name = "Celluloid YamlParser"
	app.Usage = "This Application reads cell-ci.yaml files, manages build environments."
	app.Author = "Johannes Amorosa"
	app.Email = "johannesa@celluloid-vfx.com"
	app.Version = mainVersion()

	registerFlags(app)

	registerCommands(app)

	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelp(c)
		return cli.NewExitError("", 0)
	}
	_ = app.Run(os.Args)
}

func mainVersion() string {
	s := ""
	s = s + appVersion + "\n"
	s = s + "Release-Tag: " + appReleaseTag + "\n"
	s = s + "Commit-ID: " + appShortCommitID + "\n"
	s = s + "Branch: " + appBranch + "\n"
	s = s + "CI API version: " + apiversion + "\n"
	return s
}

func civersion() bool {
	cfg, err := readConfig(path)
	switch {
	case err == nil:
		cv := cfg.getValueName("civersion")
		if cv == apiversion {
			return true
		}
	}
	return false
}
