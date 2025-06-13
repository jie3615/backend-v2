# 定义变量
PROTOC=protoc
GOCTL=goctl
RM=rm -f
RMDIR=rm -rf

# 检查命令是否存在
CHECK_PROTOC := $(shell command -v $(PROTOC) 2> /dev/null)
CHECK_GOCTL := $(shell command -v $(GOCTL) 2> /dev/null)

# 定义目标
.PHONY: all clean proto

all: proto

# 生成代码
proto:
ifndef CHECK_PROTOC
	$(error "protoc is not available, please install it")
endif
ifndef CHECK_GOCTL
	$(error "goctl is not available, please install it")
endif
	$(GOCTL) rpc protoc ./pb/meta_land_worker.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --style=go_zero

# 清理生成的文件
clean:
	$(RM) metaland_worker.go
	$(RMDIR) metalandworker
	$(RM) etc/metalandworker.yaml