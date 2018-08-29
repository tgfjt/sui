GO ?= go
app := sui

build:
	@echo "Yes! now buiding ${app}"
	@$(GO) build -o $(app)
.PHONY: build

install:
	@echo "Yes! Installing ${app} ${GOPATH}/bin/sui"
	@$(GO) install
.PHONY: install
