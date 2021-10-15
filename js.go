package main

import (
	`fmt`
	`path`
	`strings`

	`github.com/storezhang/glog`
	`github.com/storezhang/gox`
)

func js(conf *config, _ glog.Logger) (err error) {
	if dir, dirErr := gox.IsDir(conf.filepath); nil != dirErr {
		panic(dirErr)
	} else if dir {
		conf.filepath = path.Join(conf.filepath, "package.json")
	}

	dependencies := make([]jsonElement, 0, len(conf.dependencies)+1)
	dependencies = append(dependencies, jsonElement{
		path:  "version",
		value: conf.version,
	})
	for _, _dependency := range conf.dependencies {
		_module := strings.ReplaceAll(_dependency.name, `.`, `\.`)
		dependencies = append(dependencies, jsonElement{
			path:  fmt.Sprintf("dependencies.%s", _module),
			value: _dependency.version,
		})
	}
	err = json(conf.filepath, dependencies...)

	return
}
