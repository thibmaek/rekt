.DEFAULT_GOAL := build

REKT_CLI_VERSION=1.1.0-alpha
LDFLAGS="-X 'main.Version=$(REKT_CLI_VERSION)'"

dependencies:
	asdf install
	brew install jadx
	pip install --upgrade git+https://github.com/P1sec/hermes-dec

build_cli:
	cd rekt-cli && \
	go build -ldflags=$(LDFLAGS) -o ./bin/rekt-v$(REKT_CLI_VERSION) .
	cd ..

build_go: build_cli

build_docker:
	docker build -t rekt .

build: build_go build_docker
