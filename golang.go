package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dronestock/drone"
	"github.com/goexl/gfx"
)

func (p *plugin) golang(source string, labels []string) (err error) {
	if _, exists := gfx.Exists(filepath.Join(source, goModuleFilename)); !exists {
		return
	}

	args := []any{
		`mod`,
		`edit`,
	}

	// 处理依赖模块
	if err = p.each(labels, p.golangModules(args)); nil != err {
		return
	}
	// 写入模块文件
	args = append(args, goModuleFilename)
	// 执行命令
	err = p.Exec(exeGo, drone.Args(args...), drone.Dir(source))

	return
}

func (p *plugin) golangModules(args []any) moduleFunc {
	return func(module *module) {
		version := module.Version
		if !strings.HasPrefix(version, goVersionPrefix) {
			version = fmt.Sprintf(`%s%s`, goVersionPrefix, version)
		}
		args = append(args, `-require`, fmt.Sprintf(`%s@%s`, module.Name, version))
	}
}
