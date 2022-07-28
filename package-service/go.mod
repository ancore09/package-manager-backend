module github.com/ancore09/package-service

go 1.18

require (
	github.com/ancore09/package-manager-backend/package-service/pkg/package-service v0.0.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/rs/zerolog v1.27.0
	google.golang.org/grpc v1.48.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.7 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sys v0.0.0-20220610221304-9f5ed59c137d // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220725144611-272f38e5d71b // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace github.com/ancore09/package-manager-backend/package-service/pkg/package-service => ./pkg/package-service
