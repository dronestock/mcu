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

	commands := []string{"eval"}
	environments := make([]string, 0, len(conf.dependencies))

	updates := make([]string, 0, len(conf.dependencies))
	environments = append(environments, fmt.Sprintf("%s=%s", "version", conf.version))
	updates = append(updates, ".version=strenv(version)")
	for _, _dependency := range conf.dependencies {
		// 使用随机字符串是为了防止原始字符串里面出现环境变量不允许的字符
		version := gox.RandString(16)
		updates = append(updates, fmt.Sprintf(".dependencies.%s = strenv(%s)", _dependency.module, version))
		environments = append(
			environments,
			fmt.Sprintf("%s=%s", version, _dependency.version),
		)
	}
	commands = append(commands, strings.Join(updates, " | "), conf.path)
	commands = append(commands, "--inplace", "--prettyPrint")

	// 执行命令
	cmd := exec.Command(`yq`, commands...)
	cmd.Env = append(cmd.Env, environments...)
	err = cmd.Wait()

	return
}
