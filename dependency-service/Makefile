GO_BIN=$(shell go env GOPATH)/bin
BUF_EXE=$(GO_BIN)/buf$(shell go env GOEXE)

SERVICE_NAME=dependency-service
SERVICE_PATH=ancore09/package-manager-backend/dependency-service

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
	go mod download

.PHONY: run
run:
	go run cmd/main.go

.PHONY: build
build: generate .build

.PHONY: .build
.build:
	go mod tidy
	CGO_ENABLED=0  go build \
		-tags='no_mysql no_sqlite3' \
		-ldflags=" \
			-X 'github.com/$(SERVICE_PATH)/internal/config.version=$(VERSION)' \
			-X 'github.com/$(SERVICE_PATH)/internal/config.commitHash=$(COMMIT_HASH)' \
		" \
		-o ./bin/grpc-server$(shell go env GOEXE) ./cmd/grpc-server/main.go