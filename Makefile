# civersion 1

# App
APP_NAME := $(shell ciparser get output)

# Musl
MUSL := $(shell ciparser get musl)
ifeq ($(MUSL), true)
	COMPILER := '/usr/local/musl/bin/musl-gcc'
else
ifeq ($(MUSL),)
	COMPILER := '/usr/bin/gcc'
endif
endif

LDFLAGS := $(shell ciparser ldflags)
BUILD_LDFLAGS := '$(LDFLAGS)'

PWD := $(shell pwd)
GOPATH := $(shell ciparser go path)
HOST := $(shell ciparser get platform)

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
	@CC=$(COMPILER) /usr/bin/go build --ldflags $(BUILD_LDFLAGS) -o $(GOPATH)/bin/$(APP_NAME) && echo "Installing to $(GOPATH)/bin/$(APP_NAME)"

compress:
	@echo "Running $@:"
	@upx --brute $(GOPATH)/bin/$(APP_NAME) && echo "Compressed to $(GOPATH)/bin/$(APP_NAME)"
