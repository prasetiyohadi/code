.PHONY: check_install swagger generate_client tidy proto

check_install:
	which swagger >/dev/null || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	mkdir -p api
	swagger generate spec -o ./api/swagger.yaml --scan-models

generate_client: check_install swagger
	mkdir -p api
	cd api && swagger generate client -f ./swagger.yaml -A go-usvc

tidy:
	go mod tidy

proto:
	export GO111MODULE=on
	go get google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc
	protoc --proto_path=./internal/server/proto \
	--go_out=./internal/server/proto --go_opt=paths=source_relative \
	--go-grpc_out=./internal/server/proto --go-grpc_opt=paths=source_relative \
	internal/server/proto/currency/currency.proto
