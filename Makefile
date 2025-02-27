GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: api
api:
	protoc --proto_path=./api \
		   --proto_path=./third_party \
		   --go_out=paths=source_relative:./api \
		   --go-grpc_out=paths=source_relative:./api \
		   --go-http_out=paths=source_relative:./api \
           --openapi_out=fq_schema_naming=true,default_response=false:. \
		   $(API_PROTO_FILES)

.PHONY: wire
# wire
wire:
	wire
