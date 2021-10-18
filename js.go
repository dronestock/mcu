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

	elements := make([]jsonElement, 0, len(conf.dependencies)+1)
	elements = append(elements, jsonElement{
		path:  "version",
		value: conf.version,
	})
	for _, _dependency := range conf.dependencies {
		_module := strings.ReplaceAll(_dependency.name, `.`, `\.`)
		elements = append(elements, jsonElement{
			path:  fmt.Sprintf("elements.%s", _module),
			value: _dependency.version,
		})
	}
	err = json(conf.filepath, elements...)

	return
}
