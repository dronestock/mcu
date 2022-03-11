package main

import (
	`fmt`
	`strings`

	`github.com/dronestock/drone`
	`github.com/storezhang/gfx`
	`github.com/storezhang/gox`
)

func (p *plugin) dart(source string, dependencies ...dependency) (err error) {
	if !gfx.Exist(dartModuleFilename) {
		return
	}

	environments := make([]string, 0, len(dependencies))
	updates := make([]string, 0, len(dependencies))
	environments = append(environments, fmt.Sprintf(`%s=%s`, "version", p.Version))
	updates = append(updates, `.version=strenv(version)`)
	for _, dep := range dependencies {
		// 使用随机字符串是为了防止原始字符串里面出现环境变量不允许的字符
		version := gox.RandString(16)
		if to, replaced := p.isReplaced(dep, typeDart); replaced {
			updates = append(
				updates,
				fmt.Sprintf(`.dependencies.%s.git.url = %s`, dep.Module, to.Module),
				fmt.Sprintf(`.dependencies.%s.git.ref = %s`, dep.Module, to.Version),
			)
		} else {
			updates = append(updates, fmt.Sprintf(`.dependencies.%s = strenv(%s)`, dep.Module, version))
		}
		environments = append(
			environments,
			fmt.Sprintf(`%s=%s`, version, dep.Version),
		)
	}

	args := []interface{}{
		`eval`,
	}
	args = append(args, strings.Join(updates, ` | `), dartModuleFilename)
	args = append(args, `--inplace`, `--prettyPrint`)
	if p.Verbose {
		args = append(args, `--verbose`)
	}

	// 执行命令
	err = p.Exec(exeYq, drone.Args(args...), drone.Dir(source), drone.Envs(environments...))

	return
}
