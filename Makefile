
all: generate

generate:
	cp -r submodule/ ./
	go generate

.PHONY: all
