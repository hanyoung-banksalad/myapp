GOPATH:=$(shell go env GOPATH)
APP?=myapp

.PHONY: init
## init: initialize the application(myapp)
init:
	go mod download

.PHONY: build
## build: build the application(myapp)
build:
	go build -o build/${APP} cmd/main.go

.PHONY: run
## run: run the application(myapp)
run:
	go run -v -race cmd/main.go

.PHONY: format
## format: format files
format:
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/banksalad/go-banksalad/cmd/importsort@latest
	goimports -local github.com/banksalad -w .
	importsort -s github.com/banksalad -w $$(find . -name "*.go")
	gofmt -s -w .
	go mod tidy

.PHONY: test
## test: run tests
test:
	@go install github.com/rakyll/gotest@latest
	gotest -race -cover -v ./...

.PHONY: coverage
## coverage: run tests with coverage
coverage:
	@go install github.com/rakyll/gotest@latest
	gotest -race -coverprofile=coverage.txt -covermode=atomic -v ./...

.PHONY: lint
## lint: check everything's okay
lint:
	golangci-lint run ./...
	go mod verify

.PHONY: generate
## generate: generate source code for mocking
generate:
	@go install golang.org/x/tools/cmd/stringer@latest
	@go install github.com/golang/mock/mockgen@latest
	go generate ./...
	@$(MAKE) format

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':'

.PHONY: deploy-to-k8s
## deploy-to-k8s: deploy application to kubernetes
deploy-to-k8s:
	cat k8s.yaml.tmpl | gomplate -f - | kubectl apply -f -
