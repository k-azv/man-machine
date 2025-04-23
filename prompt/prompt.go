package prompt

import (
	"github.com/k-azv/man-machine/config"
)

var (
	noMarkdown = ",输出不要带任何markdown标记,"
	format = ",格式应当尽量简洁," + noMarkdown
	language   string
	Mam        string
)

// Initialize prompt, should run after LoadConfig
func Initialize() {
	language = ",回答使用的语种是" + config.Config.Language + ","
	Mam = "你是一个命令行程序专家,阅读我提供的命令行程序文档,向从未使用过这个程序的用户进行尽量简短的介绍,仅含常用内容,结合你的实践进行选择" +
	language + format
}
