# 🤖 Man-Machine

[English](../README.md)

Man-Machine 是一个通过 LLM 来轻松阅读命令行程序文档的命令行工具

## Features

- 让 LLM 帮你读手册
- 让 LLM 按你的需求生成命令

## Installation

从 [Releases page](https://github.com/k-azv/man-machine/releases) 下载编译好的二进制文件, 并添加到环境变量中

## Initialization

```shell
mam setup # 创建并打开配置文件(~/.config/mam/config.yaml)
```

在 `config.yaml` 中完成设置, 目前仅支持与 OpenAI API 协议兼容的服务

```yml
# config.yaml template

# API Key - Replace with your API key from the provider
apiKey: <YOUR_API_KEY_HERE>

# API Base URL - Fill in with your API service address
# If using OpenAI official API, use https://api.openai.com/v1
# For third-party API services, enter the complete base URL
baseURL: <YOUR_BASE_URL_HERE>

# Model Name - Specify which LLM model to use
model: <YOUR_MODEL_HERE>

# Language Setting - Specify which language the AI should use for responses
language: <LANGUAGE_HERE>
```

## Usage

基础用法

```shell
mam mam # 获取 mam 的用法
```

### Options

#### `--iwant`, `-i`

按你的需求生成对应命令

```shell
mam rm -i "删除 / 目录下的所有文件" 
# or
mam rm --iwant "删除 / 目录下的所有文件"
```

#### `--bare`, `-b`

直接执行提供的命令以获取帮助文档，而不是通过 mam 内部的逻辑获取

下面的示例中会直接将 `go help build` 的输出提供给 LLM

```shell
mam -b go help build
# or
mam --bare go help build
```
#### `--chat`, `-c`

进入连续对话模式

```shell
mam --chat
# or
mam -c
```

#### `--help`, `-h`

获取帮助文档

```shell
mam --help
# or
mam -h
```

## Roadmap

- 添加 `-q`, `--query` 选项，用于查询程序文档
- 为 `mam <command>` 添加缓存功能
- 添加 `-o`, `--output` 选项，用于重定向输出到文件

## License

MIT