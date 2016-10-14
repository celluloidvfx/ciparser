# Ciparser

This app is part of the [Celluloid VFX](http://celluloid-vfx.com/) Continuous Integration pipeline - a poor man's travis-ci clone. We use ciparser to build all our Go (golang) microservices and apps. Ciparser is a quite well working hack, to help our developers not to reinvent the wheel again. We're constantly adding things as we need it. Feel free to use - feedback and pull requests are absolutely welcomed.

## Core Features

1. Unique build entry point for all our Go code: just write *make*.
2. The developer has full control about the build behavior, without the need to configure the build pipeline.
2. Automatic compile flags generation for versioning and custom flags.
3. Single configuration file for all platforms
4. Switchable [musl-gcc](http://www.musl-libc.org/) integration
5. Switchable [upx](https://upx.github.io/) compression
6. Go dependency management
5. Extendable for all kind of parameters, platforms etc.

## What does this solve at Celluloid's:

We have an in-house developed CI/CD pipeline at Celluloid. Ciparser is used to read the cell-ci.yaml file and it provides values for *make* on build time. It also gathers necessary build environment information. Cell-ci.yaml and the Makefile are part of the source code and in git (think .travis.yml).

##  Installation

### Dependencies

- [musl-gcc](http://www.musl-libc.org/)
- [upx](https://upx.github.io/)
- [Gnu make](https://www.gnu.org/software/make/)
- [GCC](https://gcc.gnu.org/)

On Debian based systems use
```
sudo get update && apt-get install build-essential make upx-ucl musl
```

### Initial Installation
```
git clone https://github.com/celluloidvfx/ciparser
cd ciparser
go build *.go
cp ciparser $(GOPATH)/bin/ciparser
ciparser --version
```

(Optional) Once ciparser is build and in your path you can use it to build it a second time with ciparser itself. Then you get a proper output on *ciparser --version*.

```
git clone https://github.com/celluloidvfx/ciparser
cd ciparser
make
ciparser --version
```

## How to use:

### Five simple steps

1. Build ciparser and put it in your $GOPATH/bin as described above.
2. Copy the Makefile and cell-ci.yaml from the ciparser source repository as template to your newly created git project.
3. Open an editor and adjust the cell-ci.yaml to your needs. The Makefile normally shouldn't be touched.
4. Write your Go code.
5. Then simply navigate to your folder and type *make* or trigger this with a githook etc.
6. Sit back and watch make building your app based on your values provided in the cell-ci.yaml

### Annotated cell-ci.yaml

The cell-ci.yaml is the configuration file for your project build and should be in your source tree.

```
name: yourapp                                  # name of app
civersion: 1                                   # api version

build:
    active: true                               # not used
    output: yourapp.exe                        # output file name
    language: go                               # not used
    goversion: 1.7                             # go version, just spits out warning if mismatch.
    arch: amd64                                # Architecture
    platform: linux                            # supported windows/linux
    musl: true                                 # enable musl or use standard gcc
    upx: true                                  # enable upx
    deps:                                      # go library dependencies
      - "github.com/urfave/cli"
      - "github.com/ghodss/yaml"
    customvars:                                # customvars see below
       - name: ClientCert
         path: "/home/joe/cell-ls-24.celluloidvfx.inc.crt"
       - name: DomainName
         value: "Something_I_need_to_inject_into_my_app"
```

### Environment Variables

*Make* will read your makefile from your source tree. It uses ciparser to fetch several environmental information on compile time and encapsulates them as ldflags.

These "fixed" variable need to be declared in your app (i.e. main.go):

```
#Git information
var appVersion, appReleaseTag, appShortCommitID, appBranch string
```

We use a [cli library](https://github.com/urfave/cli) at Celluloid and you could display the versioning like this:

```
...
app.Version = mainVersion()
...
func mainVersion() string {
    s := ""
    s = s + appVersion + "\n"
    s = s + "Release-Tag: " + appReleaseTag + "\n"
    s = s + "Commit-ID: " + appShortCommitID + "\n"
    s = s + "Branch: " + appBranch + "\n"
    s = s + "CI API version: " + apiversion + "\n"
    return s
}
```

### Custom Variables

The Custom Vars functionality gives you the possibility to inject arbitrary information with ldflags to your build on compile time. There are two possible ways to declare a variable.

1. Value based Variables
In the cell-ci.yaml you can define a name and a value.
```
       - name: DomainName
         value: "Something_I_need_to_inject_into_my_app"
```

In this example the var will be available as *appDomainName*. The prefix app is added to the var name.

2. Path based Variable
In the cell-ci.yaml you can define a name and a path.
```
       - name: ClientCert
         path: "/path/to/your/file/you/need/to/inject"
```

In this example the var will be available as *appClientCert*. The prefix app is added to the var name. The difference to the value based vars is that the *content* of the path will be read and injected.

To do this savely all custom variables are base64 encoded and need to be decoded in runtime before use. This adds computation time to you app on startup.

In your code you have to declare the vars like this:
```
var appClientCert, appDomainName string
```

and decode them like this:
```
import (
    b64 "encoding/base64"
)
...
moo := string(decodeCompileFlags(appClientCert))
...
func decodeCompileFlags(enc string) []byte {
    d, err := b64.StdEncoding.DecodeString(enc)
    if err != nil {
        log.Fatal(err)
    }
    return d
}
```

## Running the app

ciparser will be used by **make** (see Makefile) i.e:

```
PLATFORM := $(shell ciparser get platform)
```
for more information see:
```
$(GOPATH)/bin/ciparser --help
```

## Todo


- implement more cross compilation darwin/bsd
- tidy up
- testing
- install go environments with godeb
- use docker/rkt as build environment
- more checks
- testing
- deploy
- remote build logs
- go vendor setup
- expose more compiler options
- deployment options
- autogenerate yaml files
- ....

If this is useful for you please add features and file a pull request.

## License

```
Johannes Amorosa, (C) 2016
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
