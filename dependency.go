package main

import (
	`fmt`
	`strings`
)

type dependency struct {
	module  string
	version string
}

func (d *dependency) String() string {
	return fmt.Sprintf("%s@%s", d.module, d.version)
}

func parseDependencies(originals ...string) (dependencies []*dependency) {
	// 防止有nil的依赖，因为originals始终有一个值
	dependencies = make([]*dependency, 0)
	for _, original := range originals {
		_dependency := parseDependency(original)
		if nil != _dependency {
			dependencies = append(dependencies, parseDependency(original))
		}
	}

	return
}

func parseDependency(original string) (dependency *dependency) {
	var _configs []string
	defer newDependency(_configs)

	if _configs = strings.Split(original, "@"); 2 == len(_configs) {
		return
	}
	if _configs = strings.Split(original, " "); 2 == len(_configs) {
		return
	}

	return
}

func newDependency(configs []string) (_dependency *dependency) {
	if nil != configs && 2 == len(configs) {
		_dependency = &dependency{
			module:  configs[0],
			version: configs[1],
		}
	}

	return
}
