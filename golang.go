package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/args"
)

func (p *plugin) golang(source string, labels []string) (err error) {
	if _, exists := gfx.Exists(filepath.Join(source, goModuleFilename)); !exists {
		return
	}

	ab := args.New().Long(strikethrough).Build()
	ab.Subcommand("mod", "edit")
	// 处理依赖模块
	if err = p.each(labels, p.golangModules(ab)); nil != err {
		return
	}

	// 写入模块文件
	ab.Add(goModuleFilename)
	// 执行命令
	_, err = p.Command(p.Binary.Go).Args(ab.Build()).Dir(source).Build().Exec()

	return
}

func (p *plugin) golangModules(args *args.Builder) moduleFunc {
	return func(module *module) {
		version := module.Version
		if !strings.HasPrefix(version, goVersionPrefix) {
			version = fmt.Sprintf("%s%s", goVersionPrefix, version)
		}
		args.Arg("require", fmt.Sprintf("%s@%s", module.Name, version))
	}
}
