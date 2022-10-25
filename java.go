package main

import (
	"path/filepath"

	"github.com/goexl/gfx"
)

func (p *plugin) java(source string, labels []string) (err error) {
	if mp, maven := gfx.Exists(filepath.Join(source, mavenModuleFilename)); maven {
		err = p.maven(mp, labels)
	} else if gp, gradle := gfx.Exists(filepath.Join(source, gradleModuleFilename)); gradle {
		err = p.gradle(gp, labels)
	}

	return
}
