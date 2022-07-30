#!/bin/sh

GRPC_HOST="localhost:7002"
GRPC_METHOD="ancore09.package_manager_backend.package_service.PackageService/DeletePackage"

payload=$(
  cat <<EOF
{
  "id": 1
}
EOF
)

grpcurl -plaintext -emit-defaults \
  -rpc-header 'x-app-name:dev' \
  -rpc-header 'x-app-version:1' \
  -d "${payload}" ${GRPC_HOST} ${GRPC_METHOD}