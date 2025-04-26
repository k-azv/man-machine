package prompt

import (
	"strings"

	"github.com/k-azv/man-machine/config"
)

var (
	noMarkdown = ",输出不要带任何markdown标记"
	format     = ",格式应当尽量简洁" + noMarkdown
	role       = "你是一个命令行程序专家,阅读我提供的命令行程序文档,"
	prompts    []string
)

// Mam get prompt based on type
func Mam() string {
	return strings.Join(prompts, ",")
}

// GenerateBasic generate prompt for basic mam <command>, should run after LoadConfig
func GenerateBasic() {
	language := generateLanguage()
	Basic := role + "向从未使用过这个程序的用户进行尽量简短的介绍,仅含常用内容,结合你的实践进行选择" +
		language + format
	prompts = append(prompts, Basic)
}

// GenerateIwant generate prompt for using flag `-i` or `--want`, should run after LoadConfig
func GenerateIwant(argument string) {
	language := generateLanguage()
	Iwant := role +
		"结合文档并分析用户的需求:\"" + argument + "\",生成对应的命令,同时说明你的命令的行为,注明其可能存在的风险;" +
		"#格式:你的回答应只有命令本身以及命令的行为与风险" +
		language + format
	prompts = append(prompts, Iwant)
}

// generateLanguage returns the language used for responses from config
func generateLanguage() string {
	return ",回答使用的语种是:\"" + config.Config.Language + "\""
}
