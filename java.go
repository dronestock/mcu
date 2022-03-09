package main

import (
	`path/filepath`

	`github.com/storezhang/gfx`
)

func (p *plugin) java(source string, dependencies ...dependency) (undo bool, err error) {
	if modulePath := filepath.Join(source, mavenModuleFilename); gfx.Exist(modulePath) {
		err = p.maven(modulePath, dependencies...)
	} else if modulePath = filepath.Join(source, gradleModuleFilename); gfx.Exist(modulePath) {
		err = p.gradle(modulePath, dependencies...)
	} else {
		undo = true
	}

	return
}
