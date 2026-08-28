GOCTL ?= goctl

STYLE := go_zero
MODULE := $(shell go list -m)

API_FILE := api/desc/main.api
API_DIR := api

RPC_PROTO := p9_core.proto
RPC_PROTO_PATH := rpc/proto

.PHONY: api rpc gen

api:
	$(GOCTL) api go --api $(API_FILE) --dir $(API_DIR) --style=$(STYLE)

rpc:
	cd $(RPC_PROTO_PATH) && $(GOCTL) rpc protoc $(RPC_PROTO) --go_out=../.. --go-grpc_out=../.. --zrpc_out=.. --go_opt=module=$(MODULE) --go-grpc_opt=module=$(MODULE) --module=$(MODULE) -I . --multiple --style=$(STYLE)

gen: api rpc