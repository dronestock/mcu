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
		"--inplace",
	}
	environments := make([]string, 0, len(conf.dependencies))

	updates := make([]string, 0, len(conf.dependencies))
	environments = append(environments, fmt.Sprintf("%s=%s", "version", conf.version))
	updates = append(updates, ".version=strenv(version)")
	for _, _dependency := range conf.dependencies {
		module := gox.RandString(16)
		version := gox.RandString(16)
		updates = append(updates, fmt.Sprintf(".dependencies.strenv(%s)=strenv(%s)", module, version))
		environments = append(
			environments,
			fmt.Sprintf("%s=%s", module, _dependency.module),
			fmt.Sprintf("%s=%s", version, _dependency.version),
		)
	}
	commands = append(commands, strings.Join(updates, " | "), conf.path)
	commands = append(commands, "--prettyPrint")

	// 执行命令
	cmd := exec.Command(`D:\Apps\Yq\yq.exe`, commands...)
	cmd.Env = append(cmd.Env, environments...)
	err = cmd.Wait()

	return
}
