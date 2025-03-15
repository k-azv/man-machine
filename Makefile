# 目标二进制文件名
BINARY = gman

# Go 编译命令
BUILD_CMD = go build -o $(BINARY) main.go

.PHONY: all build clean

# 默认目标
all: build

# 编译
build:
	$(BUILD_CMD)

# 清理编译文件
clean:
	rm -f $(BINARY)
