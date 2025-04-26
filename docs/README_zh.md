# 🤖 Man-Machine

[English](../README.md)

Man-Machine 是一个通过 LLM 来轻松阅读命令行程序文档的命令行程序

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

## Roadmap

- 为 `mam <command>` 添加缓存功能
- 添加 `-o`，`--output` 选项，将输出重定向到文件

## License

MIT