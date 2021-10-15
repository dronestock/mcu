package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/storezhang/glog"
	"github.com/storezhang/gox"
	"github.com/storezhang/gox/field"
)

func dart(conf *config, logger glog.Logger) (err error) {
	if dir, dirErr := gox.IsDir(conf.filepath); nil != dirErr {
		panic(dirErr)
	} else if dir {
		conf.filepath = path.Join(conf.filepath, `pubspec.yaml`)
	}

	commands := []string{`eval`}
	environments := make([]string, 0, len(conf.dependencies))
	updates := make([]string, 0, len(conf.dependencies))
	environments = append(environments, fmt.Sprintf(`%s=%s`, "version", conf.version))
	updates = append(updates, `.version=strenv(version)`)
	replacesMap := toReplacesMap(conf.replaces)
	for _, _dependency := range conf.dependencies {
		// 使用随机字符串是为了防止原始字符串里面出现环境变量不允许的字符
		version := gox.RandString(16)
		if _replace, ok := replacesMap[_dependency.name]; ok == true {
			updates = append(updates, fmt.Sprintf(`.dependencies.%s.git.url = %s`, _dependency.name, _replace.to.name))
			updates = append(updates, fmt.Sprintf(`.dependencies.%s.git.ref = %s`, _dependency.name, _replace.to.version))
		} else {
			updates = append(updates, fmt.Sprintf(`.dependencies.%s = strenv(%s)`, _dependency.name, version))
		}
		environments = append(
			environments,
			fmt.Sprintf(`%s=%s`, version, _dependency.version),
		)
	}
	commands = append(commands, strings.Join(updates, ` | `), conf.filepath)
	commands = append(commands, `--inplace`, `--prettyPrint`)

	// 执行命令
	cmd := exec.Command(`yq`, commands...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, environments...)
	if err = cmd.Run(); nil != err {
		output, _ := cmd.CombinedOutput()
		logger.Warn(`修改Dart模块描述文件出错`, field.String(`output`, string(output)), field.Error(err))
	}

	return
}

func toReplacesMap(replaces []replace) (replacesMap map[string]replace) {
	replacesMap = make(map[string]replace)
	for _, _replace := range replaces {
		replacesMap[_replace.from.name] = _replace
	}
	return
}
