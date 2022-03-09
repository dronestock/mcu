package main

import (
	`fmt`
	`path/filepath`

	`github.com/dronestock/drone`
	`github.com/storezhang/gfx`
)

func (p *plugin) golang(source string, dependencies ...dependency) (undo bool, err error) {
	modulePath := filepath.Join(source, goModuleFilename)
	if undo = !gfx.Exist(modulePath); undo {
		return
	}

	args := []interface{}{
		`mod`,
		`edit`,
	}
	for _, dep := range dependencies {
		args = append(args, `-require`, fmt.Sprintf(`%s@%s`, dep.Module, dep.Version))
	}
	// 写入模块文件
	args = append(args, modulePath)

	// 执行命令
	err = p.Exec(exeGo, drone.Args(args...), drone.Dir(source))

	return
}
