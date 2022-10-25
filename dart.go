package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dronestock/drone"
	"github.com/goexl/gfx"
	"github.com/goexl/gox"
)

func (p *plugin) dart(source string, labels []string) (err error) {
	if _, exists := gfx.Exists(filepath.Join(source, dartModuleFilename)); !exists {
		return
	}

	environments := make([]string, 0, len(labels))
	updates := make([]string, 0, len(labels))
	environments = append(environments, fmt.Sprintf(`%s=%s`, "version", p.Version))
	updates = append(updates, `.version=strenv(version)`)

	// 处理依赖模块
	if err = p.each(labels, p.dartModules(updates, environments)); nil != err {
		return
	}

	args := []any{
		`eval`,
	}
	args = append(args, strings.Join(updates, ` | `), dartModuleFilename)
	args = append(args, `--inplace`, `--prettyPrint`)
	if p.Verbose {
		args = append(args, `--verbose`)
	}

	// 执行命令
	err = p.Exec(exeYq, drone.Args(args...), drone.Dir(source), drone.StringEnvs(environments...))

	return
}

func (p *plugin) dartModules(updates []string, environments []string) moduleFunc {
	return func(module *module) {
		// 使用随机字符串是为了防止原始字符串里面出现环境变量不允许的字符
		version := gox.RandString(16)
		updates = append(updates, fmt.Sprintf(`.modules.%s = strenv(%s)`, module.Name, version))
		environments = append(environments, fmt.Sprintf(`%s=%s`, version, module.Version))
	}
}
