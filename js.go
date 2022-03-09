package main

import (
	`fmt`
	`path/filepath`
	`strings`

	`github.com/storezhang/gfx`
)

func (p *plugin) js(source string, dependencies ...dependency) (undo bool, err error) {
	modulePath := filepath.Join(source, jsModuleFilename)
	if undo = !gfx.Exist(modulePath); undo {
		return
	}

	elements := make([]jsonElement, 0, len(conf.dependencies)+1)
	elements = append(elements, jsonElement{
		path:  `version`,
		value: p.Version,
	})
	for _, dep := range dependencies {
		_module := strings.ReplaceAll(dep.Module, `.`, `\.`)
		elements = append(elements, jsonElement{
			path:  fmt.Sprintf("elements.%s", _module),
			value: dep.Version,
		})
	}
	err = p.json(modulePath, elements...)

	return
}
