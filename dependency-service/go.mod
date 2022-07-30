module github.com/ancore09/dependency-service

go 1.18

require (
	github.com/ancore09/dependency-service/pkg/dependency-service v0.0.0-00010101000000-000000000000
	github.com/ancore09/package-manager-backend/package-service/pkg/package-service v0.0.0-20220730170021-d6c781578bc4
	github.com/dgraph-io/dgo/v200 v200.0.0-20210401091508-95bfd74de60e
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/rs/zerolog v1.27.0
	google.golang.org/grpc v1.48.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.7 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220725144611-272f38e5d71b // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace github.com/ancore09/dependency-service/pkg/dependency-service => ./pkg/dependency-service
