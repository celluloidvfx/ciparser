# civersion 1

# App
APP_NAME := $(shell ciparser get output)

# Platform
PLATFORM := $(shell ciparser get platform)

# Arch
ARCH := $(shell ciparser get arch)

# Musl
MUSL := $(shell ciparser get musl)


# Upx
UPX := $(shell ciparser get upx)

# Ldflags
LDFLAGS := $(shell ciparser ldflags)
BUILD_LDFLAGS := '$(LDFLAGS)'

PWD := $(shell pwd)
GOPATH := $(shell ciparser go path)

all: gomake-all

gomake-all:
ifeq ($(UPX), true)
gomake-all: getdeps verifiers build compress
else
gomake-all: getdeps verifiers build
endif

build:
	@echo "Running $@:"
	@echo "Building for: $(PLATFORM)"
	@echo "Building for: $(ARCH)"
ifeq ($(MUSL), true)
MUSLPATH := $(shell which musl-gcc)
build: musl-build
else
build: gobuild
endif

musl-build:
	@echo "Using Musl: $(MUSLPATH)"
	@GOOS=$(PLATFORM) GOARCH=$(ARCH) CC=$(MUSLPATH) /usr/bin/go build --ldflags $(BUILD_LDFLAGS) -o $(GOPATH)/bin/$(APP_NAME) && echo "Installing to $(GOPATH)/bin/$(APP_NAME)"

gobuild:
	@echo "Not using Musl"
	@GOOS=$(PLATFORM) GOARCH=$(ARCH) /usr/bin/go build --ldflags $(BUILD_LDFLAGS) -o $(GOPATH)/bin/$(APP_NAME) && echo "Installing to $(GOPATH)/bin/$(APP_NAME)"

compress:
	@echo "Running $@:"
	@upx --brute $(GOPATH)/bin/$(APP_NAME) && echo "Compressed to $(GOPATH)/bin/$(APP_NAME)"


getdeps:
	@echo "Running $@:"
	@${GOPATH}/bin/ciparser go deps

verifiers: gometalinter spelling test

gometalinter:
	@echo "Running $@:"
	@${GOPATH}/bin/gometalinter --cyclo-over=12 --errors ./...

spelling:
	@echo "Running $@:"
	@${GOPATH}/bin/misspell *.go *.md *.MD

test:
	@echo "Running $@:"
	@/usr/bin/go test --cover ./...
