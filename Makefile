GO ?= go
GOCTL ?= goctl

# 仓库独立配置, 新建仓库时需要修改
MODULE := github.com/ns9ryan/p9_core
RPC_PROTO := p9_core.proto

STYLE := go_zero

# API
API_FILE := api/desc/main.api
API_DIR := api

# RPC
RPC_PROTO_PATH := rpc/proto

# Ent
ENT_DIR := rpc/ent

.PHONY: api
api:
	$(GOCTL) api go --api $(API_FILE) --dir $(API_DIR) --style=$(STYLE)

.PHONY: rpc
rpc:
	cd $(RPC_PROTO_PATH) && $(GOCTL) rpc protoc $(RPC_PROTO) --go_out=../.. --go-grpc_out=../.. --zrpc_out=.. --go_opt=module=$(MODULE) --go-grpc_opt=module=$(MODULE) --module=$(MODULE) -I . --multiple --style=$(STYLE)

.PHONY: ent-new
ent-new:
ifndef name
	$(error name is required, example: make ent-new name=Language)
endif
	cd rpc && $(GO) run entgo.io/ent/cmd/ent new $(name)

.PHONY: ent
ent:
	$(GO) generate ./$(ENT_DIR)

.PHONY: gen
gen: api rpc ent