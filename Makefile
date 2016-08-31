# civersion 1

# App
APP_NAME := $(shell ciparser get output)

# Musl
MUSL := $(shell ciparser get musl)
ifeq ($(MUSL), "true")
	CC := /usr/local/musl/bin/musl-gcc
else
ifeq ($(MUSL), "false")
	CC := $(shell which gcc)
endif
endif

LDFLAGS := $(shell ciparser ldflags)
BUILD_LDFLAGS := '$(LDFLAGS)'

PWD := $(shell pwd)
GOPATH := $(shell ciparser go path)
HOST := $(shell ciparser get platform)

all: install

getdeps:
	@go get -u gitlab.celluloidvfx.inc/dev-op/golint && echo "Installed golint"
	@go get -u gitlab.celluloidvfx.inc/dev-op/gocyclo && echo "Installed gocyclo"
	@go get -u gitlab.celluloidvfx.inc/dev-op/go-misc/deadcode && echo "Installed deadcode"
	@go get -u gitlab.celluloidvfx.inc/dev-op/misspell/cmd/misspell && echo "Installed misspell"
	@go get -u honnef.co/go/simple/cmd/gosimple && echo "Installed gosimple"
	@ciparser go deps
verifiers: vet fmt simple lint cyclo spelling

vet:
	@echo "Running $@:"
	@GO15VENDOREXPERIMENT=1 go tool vet -all *.go
	@GO15VENDOREXPERIMENT=1 go tool vet -shadow=true *.go

fmt:
	@echo "Running $@:"
	@GO15VENDOREXPERIMENT=1 gofmt -s -l *.go

simple:
	@echo "Running $@:"
	@GO15VENDOREXPERIMENT=1 gosimple

lint:
	@echo "Running $@:"
 	@GO15VENDOREXPERIMENT=1 ${GOPATH}/bin/golint *.go

cyclo:
	@echo "Running $@:"
	@GO15VENDOREXPERIMENT=1 ${GOPATH}/bin/gocyclo -over 65 *.go

build: getdeps verifiers

deadcode:
	@GO15VENDOREXPERIMENT=1 ${GOPATH}/bin/deadcode

spelling:
	@GO15VENDOREXPERIMENT=1 ${GOPATH}/bin/misspell *.go *.rm

check:
	@echo "Running $@:"
	@ciparser check

test: build
 	@echo "Running $@:"
 	@GO15VENDOREXPERIMENT=1 go test .

gomake-all: build
	@echo "Installing to $(GOPATH)/bin/$(APP_NAME)"
	@GO15VENDOREXPERIMENT=1 go build --ldflags $(BUILD_LDFLAGS) -o $(GOPATH)/bin/$(APP_NAME)

install: gomake-all

# clean:
# 	@echo "Cleaning up all the generated files:"
# 	@rm -fv minio minio.test cover.out
# 	@find . -name '*.test' | xargs rm -fv
# 	@rm -rf isa-l
# 	@rm -rf build
# 	@rm -rf release
