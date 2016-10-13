ciparser
====

This app is part of the [Celluloid VFX](http://celluloid-vfx.com/) Continuous Integration pipeline - a poor man travis-ci clone. We use this to build all our go microservices. This is a quite well working hack to help our developers not to reinvent everything always again. We're constantly changing things as we need it. Feel free to use - feedback is welcomed.

How to use:
----

1. Build this go app manually and put it in your GOPATH/bin.
2. Copy the Makefile and the cell-ci.yaml file in your newly created git folder.
3. Point the dependencies in the Makefile to the official github repos
3. Edit the yaml file to your needs
4. Type "make"

TODO
----

- implement cross compile
- tidy up
- testing
- install go environments with godeb
- docker environment
- more checks
- autonmatic testing
- deploy
- notification

What does this solve at Celluloid's:
---

1. Unique build entrypoint for all our go code: just write "make"
2. Automatic compile flags generation for git versioning
3. Automatic compile flags generation for custom flags
3. Single configuration file
4. Musl build
5. Extendable for all kind of parameters

License
---

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
