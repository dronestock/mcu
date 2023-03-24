package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/args"
	"github.com/goexl/gox/rand"
)

func (p *plugin) dart(source string, labels []string) (err error) {
	if _, exists := gfx.Exists(filepath.Join(source, dartModuleFilename)); !exists {
		return
	}

	environments := make([]string, 0, len(labels))
	updates := make([]string, 0, len(labels))
	environments = append(environments, fmt.Sprintf("%s=%s", "version", p.Version))
	updates = append(updates, ".version=strenv(version)")

	// 处理依赖模块
	if err = p.each(labels, p.dartModules(updates, environments)); nil != err {
		return
	}

	ab := args.New().Build()
	ab.Subcommand("eval")
	ab.Add(strings.Join(updates, " | "), dartModuleFilename)
	ab.Flag("inplace")
	ab.Flag("prettyPrint")
	if p.Verbose {
		ab.Flag("verbose")
	}

	// 执行命令
	cb := p.Command(p.Binary.Yq)
	cb.Args(ab.Build())
	cb.Dir(source)
	cb.Environment().Kv("version", p.Version).Build()
	_, err = cb.Build().Exec()

	return
}

func (p *plugin) dartModules(updates []string, environments []string) moduleFunc {
	return func(module *module) {
		// 使用随机字符串是为了防止原始字符串里面出现环境变量不允许的字符
		version := rand.New().String().Length(16).Build().Generate()
		updates = append(updates, fmt.Sprintf(`.modules.%s = strenv(%s)`, module.Name, version))
		environments = append(environments, fmt.Sprintf(`%s=%s`, version, module.Version))
	}
}
