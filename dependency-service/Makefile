GO_BIN=$(shell go env GOPATH)/bin
BUF_EXE=$(GO_BIN)/buf$(shell go env GOEXE)

.PHONY: generate
generate:
	$(BUF_EXE) generate

.PHONY: bufinit
bufinit:
	$(BUF_EXE) mod init

.PHONY: deps
deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
	go install github.com/envoyproxy/protoc-gen-validate@v0.6.1
	go install github.com/bufbuild/buf/cmd/buf@v0.56.0
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking@v0.56.0
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@v0.56.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1

.PHONY: run
run:
	go run cmd/main.go