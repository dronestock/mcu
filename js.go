package main

import (
	`path`

	`github.com/storezhang/gox`
	`github.com/storezhang/replace`
)

func js(conf *config) (err error) {
	if dir, dirErr := gox.IsDir(conf.path); nil != dirErr {
		panic(dirErr)
	} else if dir {
		conf.path = path.Join(conf.path, "package.json")
	}

	_replace := replace.NewJSONReplace(gox.GetFilename(conf.path), replace.JSONReplaceElement{
		Path:  "version",
		Value: conf.version,
	})
	err = _replace.Replace(path.Dir(conf.path))

	return
}
