# 🤖 Man-Machine

[中文](./docs/README_zh.md)

Man-Machine is a command-line tool that uses LLM to help users read documentation for command-line programs easily.

## Features

- Let LLM read command manuals for you.
- Let LLM generate commands according to your needs.

## Installation

Download the prebuilt binary from the [Releases page](https://github.com/k-azv/man-machine/releases) and add it to your PATH environment variable.

## Initialization

```shell
mam setup # Create and open the configuration file (~/.config/mam/config.yaml)
```

Complete the setup in `config.yaml`. Currently, only services compatible with the OpenAI API protocol are supported.

```yaml
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

Basic usage:

```shell
mam mam # Get the usage of mam
```

### Options

#### `--iwant`, `-i`

Generate commands according to your needs.

```shell
mam rm -i "Delete all files under the / directory"
# or
mam rm --iwant "Delete all files under the / directory"
```

#### `--bare`, `-b`

Execute the provided command literally to fetch help documentation, bypassing mam's internal attempts.

In the following example, the output of `go help build` is directly provided to the LLM.

```shell
mam -b go help build
# or
mam --bare go help build
```

#### `--help`, `-h`

Display help information.

```shell
mam --help
# or
mam -h
```

## Roadmap

- Add `-q`, `--query` options to query the LLM about program documentation.
- Add cache functionality for `mam <command>`.
- Add `-o`, `--output` options to redirect output to a file.
- Add `-c`, `--conversation` option for continuous conversation mode.

## License

MIT