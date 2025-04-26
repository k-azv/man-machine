# 🤖 Man-Machine

[中文](./docs/README_zh.md)

Man-Machine is a command-line tool that uses LLM to help users read documentation for command-line programs easily.

## Features
- Let LLM read command manuals for you.

## Installation

Download the prebuilt binary from the [Releases page](https://github.com/k-azv/man-machine/releases) and add it to your PATH environment variable.

## Initialization

```shell
mam setup # Create and open the configuration file(~/.config/mam/config.yaml)
```

Complete the setup in `config.yaml`. Currently only services compatible with the OpenAI API protocol are supported

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

Basic usage
```shell
mam mam # Get the usage of mam
```

## Roadmap
- Add `-i`, `--iwant` options, generate commands as needed after reading the documentation
- Add `-l`, `--level` options to customize the level of detail in the output
- Cache functionality for `mam <command>`
 Add `-o`, `--output` options, redirect output to a file

## License

MIT