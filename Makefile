LOCAL_BIN := $(CURDIR)/bin

install golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... -- config .golangci.reference.yml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.35.1
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.1.0
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@v0.1.7

vendor-proto:
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi


generation-protoc-chat:
	mkdir -p $(CURDIR)/pkg/chat_v1
	protoc --proto_path grpc/chat_v1 --proto_path vendor.protogen \
            --go_out=pkg/chat_v1 --go_opt=paths=source_relative \
            --plugin=protoc-gen-go=bin/protoc-gen-go \
            --go-grpc_out=pkg/chat_v1 --go-grpc_opt=paths=source_relative \
            --plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
            --grpc-gateway_out=pkg/chat_v1 --grpc-gateway_opt=paths=source_relative \
            --plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
            --validate_out=lang=go:pkg/chat_v1 --validate_opt=paths=source_relative \
            --plugin=protoc-gen-validate=bin/protoc-gen-validate \
            grpc/chat_v1/chat.proto

generation-protoc-user:
	mkdir -p pkg/user_v1
	protoc --proto_path grpc/user_v1  --proto_path vendor.protogen \
            --go_out=pkg/user_v1 --go_opt=paths=source_relative \
            --plugin=protoc-gen-go=bin/protoc-gen-go \
            --go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
            --plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
            --validate_out=lang=go:pkg/user_v1 --validate_opt=paths=source_relative \
            --plugin=protoc-gen-validate=bin/protoc-gen-validate \
            --grpc-gateway_out=pkg/user_v1 --grpc-gateway_opt=paths=source_relative \
            --plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
            grpc/user_v1/user.proto

generation-protoc-auth:
	mkdir -p pkg/auth_v1
	protoc --proto_path grpc/auth_v1  --proto_path vendor.protogen \
            --go_out=pkg/auth_v1 --go_opt=paths=source_relative \
            --plugin=protoc-gen-go=bin/protoc-gen-go \
            --go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
            --plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
            --validate_out=lang=go:pkg/auth_v1 --validate_opt=paths=source_relative \
            --plugin=protoc-gen-validate=bin/protoc-gen-validate \
            --grpc-gateway_out=pkg/auth_v1 --grpc-gateway_opt=paths=source_relative \
            --plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
			grpc/auth_v1/auth.proto