package main

import (
	`fmt`
	`os/exec`
	`path`
	`strings`

	`github.com/storezhang/glog`
	`github.com/storezhang/gox`
)

func dart(conf *config, _ glog.Logger) (err error) {
	if dir, dirErr := gox.IsDir(conf.path); nil != dirErr {
		panic(dirErr)
	} else if dir {
		conf.path = path.Join(conf.path, "pubspec.yaml")
	}

	commands := []string{
		"e",
		"-i",
	}

	updates := make([]string, 0, len(conf.dependencies))
	updates = append(updates, fmt.Sprintf(".version=%s", conf.version))
	for _, _dependency := range conf.dependencies {
		updates = append(updates, fmt.Sprintf(".dependencies.%s=%s", _dependency.module, _dependency.version))
	}
	commands = append(commands, strings.Join(updates, "|"), conf.path)

	// 执行命令
	err = exec.Command(`D:\Apps\Yq\yq.exe`, commands...).Wait()

	return
}
