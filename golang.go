package main

import (
	`fmt`
	`os/exec`
	`path`

	`github.com/storezhang/glog`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

func golang(conf *config, logger glog.Logger) (err error) {
	if dir, dirErr := gox.IsDir(conf.filepath); nil != dirErr {
		panic(dirErr)
	} else if dir {
		conf.filepath = path.Join(conf.filepath, `go.mod`)
	}

	commands := []string{
		`mod`,
		`edit`,
	}
	for _, _dependency := range conf.dependencies {
		commands = append(commands, `-require`, fmt.Sprintf("%s@%s", _dependency.name, _dependency.version))
	}
	commands = append(commands, conf.filepath)

	// 执行命令
	cmd := exec.Command(`go`, commands...)
	if err = cmd.Run(); nil != err {
		output, _ := cmd.CombinedOutput()
		logger.Warn(
			`修改Dart模块描述文件出错`,
			field.String(`output`, string(output)),
			field.Strings(`command`, commands...),
			field.Error(err),
		)
	}

	return
}
