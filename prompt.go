package main

import (
	"strings"

	"github.com/k-azv/man-machine/config"
)

const (
	noMarkdown = ",输出不要带任何markdown标记"
	format     = ",格式应当尽量简洁" + noMarkdown
	role       = "你是一个命令行程序专家,阅读我提供的命令行程序文档,"
)

// PromptGenerator manages prompt generation and storage
type PromptGenerator struct {
	Prompts  []string
	Language string
}

// NewPromptGenerator creates a new PromptGenerator instance
func NewPromptGenerator(cfg config.Config) *PromptGenerator {
	p := &PromptGenerator{
		Prompts: []string{},
		Language: cfg.Language,
	}
	return p
}

// Mam returns the combined prompt string
func (p *PromptGenerator) Mam() string {
	return strings.Join(p.Prompts, ",")
}

// GenerateBasic generates a basic prompt for basic mam <command>
func (p *PromptGenerator) GenerateBasic() {
	language := p.Language
	Basic := role + "向从未使用过这个程序的用户进行尽量简短的介绍,仅含常用内容,结合你的实践进行选择" +
		language + format
	p.Prompts = append(p.Prompts, Basic)
}

// GenerateIwant generates a prompt when using flag `-i` or `--want`
func (p *PromptGenerator) GenerateIwant(argument string) {
	language := p.Language
	Iwant := role +
		"结合文档并分析用户的需求:\"" + argument + "\",生成对应的命令,同时说明你的命令的行为,注明其可能存在的风险;" +
		"#格式:你的回答应只有命令本身以及命令的行为与风险" +
		language + format
	p.Prompts = append(p.Prompts, Iwant)
}