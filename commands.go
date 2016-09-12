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
	"fmt"

	"github.com/urfave/cli"
)

func registerCommands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:  "get",
			Usage: "Fetches a value from cell-ci.yaml",
			Action: func(c *cli.Context) error {
				cfg, err := readConfig(path)
				switch {
				case err == nil:
					a := c.Args().First()
					b := cfg.getValueName(a)
					fmt.Println(b.(string))
					return cli.NewExitError("", 0)
				default:
					return cli.NewExitError("Can't read file: "+path, 1)
				}

			},
		},
		{
			Name:  "ldflags",
			Usage: "Create ldflags",
			Action: func(c *cli.Context) error {
				cfg, err := readConfig(path)
				switch {
				case err == nil:
					fmt.Println(genLDFlags(cfg))
					return cli.NewExitError("", 0)
				default:
					return cli.NewExitError("Can't read file: "+path, 1)
				}
			},
		},
		{
			Name:  "check",
			Usage: "Check for build environment",
			Action: func(c *cli.Context) error {
				if !civersion() {
					fmt.Println("Ciparser Api Version mismatch this project uses apiversion " + apiversion)
				}

				cfg, err := readConfig(path)
				switch {
				case err == nil:
					// Go path
					if gopath, e := getGoPath(); e != nil {
						return cli.NewExitError("Can't get gopath", 1)
					} else {
						fmt.Println("Found GOPATH: " + gopath)
					}

					// Go version
					if v, e := getInstalledGoVersion(); e != nil {
						return cli.NewExitError("Can't determine go version", 1)
					} else {
						n := cfg.getValueName("goversion")
						if n.(string) != v {
							return cli.NewExitError("Go Version mismatch: "+v+" is installed, but need "+n.(string), 1)
						} else {
							fmt.Println("Found Go Version: " + v)
						}
					}

					return cli.NewExitError("", 0)
				default:
					return cli.NewExitError("Can't read file: "+path, 1)
				}
			},
		},
		{
			Name:  "go",
			Usage: "Handle go environment",
			Action: func(c *cli.Context) error {
				_ = cli.ShowSubcommandHelp(c)
				return cli.NewExitError("", 0)
			},
			Subcommands: []cli.Command{
				{
					Name:  "path",
					Usage: "Fetches go path",
					Action: func(c *cli.Context) error {
						if v, e := getGoPath(); e != nil {
							return cli.NewExitError("GOPATH not set", 1)
						} else {
							fmt.Println(v)
							return cli.NewExitError("", 0)
						}
					},
				},
				{
					Name:  "version",
					Usage: "Fetches needed go version string",
					Action: func(c *cli.Context) error {
						if v, e := getInstalledGoVersion(); e != nil {
							return cli.NewExitError("GOPATH not set or go executable not found", 1)
						} else {
							fmt.Println(v)
							return cli.NewExitError("", 0)
						}
					},
				},
				{
					Name:  "bin",
					Usage: "Fetches path of go executable",
					Action: func(c *cli.Context) error {
						if v, e := getGoBin(); e != nil {
							return cli.NewExitError("Go executable not found", 1)
						} else {
							fmt.Println(v)
							return cli.NewExitError("", 0)
						}
					},
				},
				{
					Name:  "deps",
					Usage: "Installs project dependencies",
					Action: func(c *cli.Context) error {
						cfg, err := readConfig(path)
						switch {
						case err == nil:
							d := cfg.getValueName("deps")
							for _, v := range d.([]string) {
								fmt.Printf("Installing: " + v + " ... ")
								if e := installGoDeps(v); e != nil {
									fmt.Println(e)
									return cli.NewExitError("Failed to install go dependency", 1)
								} else {
									fmt.Printf("Done!\n")
								}
							}
							return cli.NewExitError("", 0)
						default:
							return cli.NewExitError("Can't read config: "+path, 1)
						}
					},
				},
			},
		},
	}
}
