package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/goexl/gfx"
)

func (p *plugin) js(source string, labels []string) (err error) {
	modulePath := filepath.Join(source, jsModuleFilename)
	if _, exists := gfx.Exists(modulePath); !exists {
		return
	}

	dependencies := make([]jsonElement, 0, len(labels)+1)
	dependencies = append(dependencies, jsonElement{
		path:  `version`,
		value: p.Version,
	})

	// 处理依赖模块
	if err = p.each(labels, p.jsModules(dependencies)); nil != err {
		return
	}
	err = p.json(modulePath, dependencies...)

	return
}

func (p *plugin) jsModules(dependencies []jsonElement) moduleFunc {
	return func(module *module) {
		_module := strings.ReplaceAll(module.Name, `.`, `\.`)
		dependencies = append(dependencies, jsonElement{path: fmt.Sprintf("dependencies.%s", _module), value: module.Version})
	}
}
