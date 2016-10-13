# civersion 1

# App
APP_NAME := $(shell ciparser get output)

# Platform
PLATFORM := $(shell ciparser get platform)

# Musl
MUSL := $(shell ciparser get musl)
ifeq ($(MUSL), true)
	ifeq ($(PLATFORM), "linux")
		COMPILER := $(shell which musl-gcc)
	endif
else

	COMPILER := $(shell which gcc)
endif

LDFLAGS := $(shell ciparser ldflags)
BUILD_LDFLAGS := '$(LDFLAGS)'

PWD := $(shell pwd)
GOPATH := $(shell ciparser go path)

all: gomake-all

gomake-all: getdeps verifiers build compress

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

build:
	@echo "Running $@:"
	@echo "Using Linker: $(COMPILER)"
	@echo "Building for: $(PLATFORM)"
	@GOOS=$(PLATFORM) CC=$(COMPILER) /usr/bin/go build --ldflags $(BUILD_LDFLAGS) -o $(GOPATH)/bin/$(APP_NAME) && echo "Installing to $(GOPATH)/bin/$(APP_NAME)"

compress:
	@echo "Running $@:"
	@upx --brute $(GOPATH)/bin/$(APP_NAME) && echo "Compressed to $(GOPATH)/bin/$(APP_NAME)"
