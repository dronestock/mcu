package main

import (
	"fmt"

	"github.com/beevik/etree"
)

const (
	keyProject       = `project`
	keyVersion       = `version`
	keyDependencies  = `dependencies`
	keyDependency    = `dependency`
	keyGroupId       = `groupId`
	keyArtifactId    = `artifactId`
	dependencyFormat = `dependency[groupId=%s and artifactId=%s]`
)

func (p *plugin) maven(filename string, labels []string) (err error) {
	doc := etree.NewDocument()
	if err = doc.ReadFromFile(filename); nil != err {
		return
	}

	project := doc.SelectElement(keyProject)
	if nil == project {
		project = doc.CreateElement(keyProject)
	}

	version := project.SelectElement(keyVersion)
	if nil == version {
		version = project.CreateElement(keyVersion)
	}
	version.SetText(p.Version)

	dependencies := project.SelectElement(keyDependencies)
	if nil == dependencies {
		dependencies = project.CreateElement(keyDependencies)
	}

	// 处理依赖模块
	if err = p.each(labels, p.mavenModules(dependencies)); nil != err {
		return
	}

	// 写入文件
	doc.Indent(xmlSpaces)
	err = doc.WriteToFile(filename)

	return
}

func (p *plugin) mavenModules(dependencies *etree.Element) moduleFunc {
	return func(module *module) {
		path := etree.MustCompilePath(fmt.Sprintf(dependencyFormat, module.Group, module.Artifact))
		_dependency := dependencies.FindElementPath(path)
		if nil == _dependency {
			_dependency = dependencies.CreateElement(keyDependency)
			_dependency.CreateElement(keyGroupId).SetText(module.Group)
			_dependency.CreateElement(keyArtifactId).SetText(module.Artifact)
			_dependency.CreateElement(keyVersion).SetText(module.Version)
		} else {
			_dependency.SelectElement(keyGroupId).SetText(module.Group)
			_dependency.SelectElement(keyArtifactId).SetText(module.Artifact)
			_dependency.SelectElement(keyVersion).SetText(module.Version)
		}
	}
}
