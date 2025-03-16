# 目标二进制文件名
BINARY = gman

# 默认目标操作系统和架构（Linux 64 位）
GOOS ?= linux
GOARCH ?= amd64

# Go 编译命令
BUILD_CMD = GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINARY) main.go

.PHONY: all build clean windows

# 默认目标（编译 Linux 64 位）
all: build

# 编译 Linux
build:
	$(BUILD_CMD)

# 编译 Windows 64 位
windows:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY).exe main.go

# 清理编译文件
clean:
	rm -f $(BINARY) $(BINARY).exe
