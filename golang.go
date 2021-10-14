package main

import (
	`fmt`
	`os/exec`
	`path`

	`github.com/storezhang/glog`
	`github.com/storezhang/gox`
)

func golang(conf *config, _ glog.Logger) (err error) {
	if dir, dirErr := gox.IsDir(conf.filepath); nil != dirErr {
		panic(dirErr)
	} else if dir {
		conf.filepath = path.Join(conf.filepath, "go.mod")
	}

	commands := []string{
		"mod",
		"edit",
	}
	for _, _dependency := range conf.dependencies {
		commands = append(commands, "-require", fmt.Sprintf("%s@%s", _dependency.module, _dependency.version))
	}
	commands = append(commands, conf.filepath)

	// 执行命令
	err = exec.Command("go", commands...).Run()

	return
}
