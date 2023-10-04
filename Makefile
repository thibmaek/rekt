.DEFAULT_GOAL := build

REKT_CLI_VERSION=1.0.0
LDFLAGS="-X 'main.Version=$(REKT_CLI_VERSION)'"

build_cli:
	cd rekt-cli && \
	go build -ldflags=$(LDFLAGS) -o ./bin/rekt-v$(REKT_CLI_VERSION) .
	cd ..

build_go: build_cli

build_docker:
	docker build -t rekt .

build: build_go build_docker
