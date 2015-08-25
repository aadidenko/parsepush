INSTALL_DEPS=go get github.com/mitchellh/gox && \
go get github.com/facebookgo/parse

init:
	@echo "Init development environment..." && \
	make deps

deps:
	@echo "Install dependencies..."
	@$(INSTALL_DEPS)

run:
	go run main.go

build:
	gox -osarch=linux/amd64 -output=bin/parsepush
