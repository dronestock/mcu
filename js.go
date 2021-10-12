package main

import (
	`path`

	`github.com/storezhang/gox`
)

func js(conf *config) (err error) {
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
