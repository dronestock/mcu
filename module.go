package main

import (
	`fmt`
	`strings`
)

type module struct {
	name    string
	version string
}

func (d *module) String() string {
	return fmt.Sprintf("%s@%s", d.name, d.version)
}

func parseMoules(originals ...string) (modules []*module) {
	// 防止有nil的依赖，因为originals始终有一个值（上层在用strings.Split时，空字符串也会分隔成一个有一个空字符串的数组）
	modules = make([]*module, 0)
	for _, original := range originals {
		_module := parseModule(original)
		if nil != _module {
			modules = append(modules, _module)
		}
	}

	return
}

func parseModule(original string) (module *module) {
	var _configs []string
	defer func() {
		module = newModule(_configs)
	}()

	if _configs = strings.Split(original, "@"); 2 == len(_configs) {
		return
	}
	if _configs = strings.Split(original, " "); 2 == len(_configs) {
		return
	}

	return
}

func newModule(configs []string) (_module *module) {
	length := len(configs)
	if nil != configs {
		_module = new(module)
		if 1 == length {
			_module.name = configs[0]
		}
		if 2 <= length {
			_module.version = configs[1]
		}
	}

	return
}
