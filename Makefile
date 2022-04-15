GO    := GO111MODULE=on go
PROMU := $(GOPATH)/bin/promu
pkgs   = $(shell $(GO) list ./... | grep -v /vendor/)

PREFIX                  ?= $(shell pwd)
BIN_DIR                 ?= $(shell pwd)
DOCKER_IMAGE_NAME       ?= ipvs-exporter
DOCKER_IMAGE_TAG        ?= $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))
TAG 					:= $(shell echo `if [ "$(TRAVIS_BRANCH)" = "master" ] || [ "$(TRAVIS_BRANCH)" = "" ] ; then echo "latest"; else echo $(TRAVIS_BRANCH) ; fi`)

all: format build test

info:
	@echo ">> show project info"
	@$(PROMU) info

style:
	@echo ">> checking code style"
	@! gofmt -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

test:
	@echo ">> running tests"
	@$(GO) test -mod=vendor -short $(pkgs)

format:
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

vet:
	@echo ">> vetting code"
	@$(GO) vet $(pkgs)

build: promu
	@echo ">> building binaries"
	@$(PROMU) build --prefix $(PREFIX)

crossbuild: promu
	@echo ">> crossbuilding binaries"
	@$(PROMU) crossbuild

tarball: promu
	@echo ">> building release tarball"
	@$(PROMU) tarball --prefix $(PREFIX) $(BIN_DIR)

docker:
	@echo ">> building docker image"
	@docker build -t "$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" .

rpm:
	@echo ">> building rpm package"
	@$(PREFIX)/rpm/makerpm.sh rpm

srpm:
	@echo ">> building rpm package"
	@$(PREFIX)/rpm/makerpm.sh srpm

push:
	@echo ">> pushing docker image, $(DOCKER_USER),$(DOCKER_IMAGE_NAME),$(TAG)"
	@echo $(DOCKER_PASS) | docker login -u "$(DOCKER_USER)" --password-stdin
	@docker tag "$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" "$(DOCKER_USER)/$(DOCKER_IMAGE_NAME):$(TAG)"
	@docker push "$(DOCKER_USER)/$(DOCKER_IMAGE_NAME):$(TAG)"

release: promu github-release
	@echo ">> pushing binary to github with ghr"
	@$(PROMU) crossbuild tarballs
	@$(PROMU) release .tarballs

promu:
	@GOOS=$(shell uname -s | tr A-Z a-z) \
		GOARCH=$(subst x86_64,amd64,$(patsubst i%86,386,$(shell uname -m))) \
		$(GO) install github.com/prometheus/promu@latest

github-release:
	@GOOS=$(shell uname -s | tr A-Z a-z) \
		GOARCH=$(subst x86_64,amd64,$(patsubst i%86,386,$(shell uname -m))) \
		$(GO) install github.com/aktau/github-release@latest

.PHONY: all style format build test vet tarball docker promu rpm srpm
