package main

import (
	`path`

	`github.com/storezhang/glog`
	`github.com/storezhang/gox`
)

func js(conf *config, _ glog.Logger) (err error) {
	if dir, dirErr := gox.IsDir(conf.path); nil != dirErr {
		panic(dirErr)
	} else if dir {
		conf.path = path.Join(conf.path, "package.json")
	}

	err = json(conf.path, jsonElement{
		path:  "version",
		value: conf.version,
	})

	return
}
