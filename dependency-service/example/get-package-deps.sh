#!/bin/sh

GRPC_HOST="localhost:7003"
GRPC_METHOD="ancore09.package_manager_backend.dependency_service.PackageService/GetPackageDeps"

payload=$(
  cat <<EOF
{
  "id": 3
}
EOF
)

grpcurl -plaintext -emit-defaults \
  -rpc-header 'x-app-name:dev' \
  -rpc-header 'x-app-version:1' \
  -d "${payload}" ${GRPC_HOST} ${GRPC_METHOD}