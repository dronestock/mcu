package main

import (
	`fmt`
	`strings`
)

type module struct {
	name    string
	version string
}

func (m *module) String() string {
	return fmt.Sprintf("%s@%s", m.name, m.version)
}

func parseMoules(key string) (modules []module) {
	_config := env(key)
	if "" == _config {
		return
	}

	originals := strings.Split(_config, `,`)
	modules = make([]module, 0, len(originals))
	for _, original := range originals {
		modules = append(modules, parseModule(original))
	}

	return
}

func parseModule(original string) (module module) {
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

func newModule(configs []string) (_module module) {
	length := len(configs)
	if nil != configs {
		if 1 == length {
			_module.name = configs[0]
		} else if 2 <= length {
			_module.name = configs[0]
			_module.version = configs[1]
		}
	}

	return
}
