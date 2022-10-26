package main

import (
	"fmt"

	"github.com/beevik/etree"
)

const (
	keyProject      = `project`
	keyVersion      = `version`
	keyDependencies = `dependencies`
	keyDependency   = `dependency`
	keyGroupId      = `groupId`
	keyArtifactId   = `artifactId`
	groupFormat     = `dependency[groupId='%s']`
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
		_dependency := p.findByGroup(dependencies, module.Group, module.Artifact)
		if nil == _dependency {
			_dependency = dependencies.CreateElement(keyDependency)
			_dependency.CreateElement(keyGroupId).SetText(module.Group)
			_dependency.CreateElement(keyArtifactId).SetText(module.Artifact)
			_dependency.CreateElement(keyVersion).SetText(module.Version)
		} else {
			_dependency.SelectElement(keyVersion).SetText(module.Version)
		}
	}
}

// 解决不了XPath同时选取两个子节点值的问题，特意加的方法，如果找到解决办法可用XPath解决
func (p *plugin) findByGroup(dependencies *etree.Element, group string, artifact string) (dependency *etree.Element) {
	elements := dependencies.FindElements(fmt.Sprintf(groupFormat, group))
	for _, element := range elements {
		artifactId := element.SelectElement(keyArtifactId)
		if nil != artifactId && artifact == artifactId.Text() {
			dependency = element
		}

		if nil != dependency {
			return
		}
	}

	return
}
