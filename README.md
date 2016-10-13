ciparser
====

This app is part of the [Celluloid VFX](http://celluloid-vfx.com/) Continuous Integration pipeline - a poor man travis-ci clone. We use this to build all our go microservices. This is a quite well working hack to help our developers not to reinvent the wheel again. We're constantly changing things as we need it. Feel free to use - feedback is absolutely welcomed.

What does this solve at Celluloid's:
---

1. Unique build entrypoint for all our go code: just write *make*. This is useful if you trigger build jobs via git hooks. The developer has full control about the job behavior, without the need to change the build pipeline.
2. Automatic compile flags generation for versioning and custom flags.
3. Single configuration file
4. [musl-gcc](http://www.musl-libc.org/) builds
5. [upx](https://upx.github.io/) compression
5. Extendable for all kind of parameters, platforms etc.


How to use:
----

1. Build this go app manually and put it in your GOPATH/bin.
2. Copy the Makefile and cell-ci.yaml as template to your newly created git folder.
3. Edit the cell-ci. yaml file to your needs
4. Type "make"
5. Sit back and watch your machine build


Annotated cell-ci.yaml
----
The cell-ci.yaml is the configuration file for your project build.

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

Customvars and git hashes are base64 encoded and added as ldflags to your app in compile time. The prefix app is added to the var name.
In our above cell-ci.yaml the var *clientCert* and *domainName* will be available in your app as *appClientCert* and *appDomainName*. (You have to declare them of course in your code.)
In the example content of *appClientCert* will be the read from /home/joe/cell-ls-24.celluloidvfx.inc.crt (base64 encoded), *appDomainName* has the value: Something_I_need_to_inject_into_my_app.

In your code you have to declare the vars like this
```
var appVersion, appReleaseTag, appShortCommitID, appBranch string
var appClientCert, appDomainName string
```

RUNNING THE APP
==================

ciparser will be used by **make** (see Makefile) i.e:

```
PLATFORM := $(shell ciparser get platform)
```
for more information see:
```
$(GOPATH)/bin/ciparser --help
```

TODO
----

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
- generate yaml files
- ....

If this is useful for you please add features and file a pull request.

License
---

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
