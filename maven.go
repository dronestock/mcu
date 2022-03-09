package main

import (
	`fmt`

	`github.com/beevik/etree`
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

func (p *plugin) maven(filename string, dependencies ...dependency) (err error) {
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

	dependenciesElement := project.SelectElement(keyDependencies)
	if nil == dependenciesElement {
		dependenciesElement = project.CreateElement(keyDependencies)
	}

	// 修改依赖
	for _, dep := range dependencies {
		path := etree.MustCompilePath(fmt.Sprintf(dependencyFormat, dep.Group, dep.Artifact))
		dependencyElement := dependenciesElement.FindElementPath(path)
		if nil == dependencyElement {
			dependencyElement = dependenciesElement.CreateElement(keyDependency)
			dependencyElement.CreateElement(keyGroupId).SetText(dep.Group)
			dependencyElement.CreateElement(keyArtifactId).SetText(dep.Artifact)
			dependencyElement.CreateElement(keyVersion).SetText(dep.Version)
		} else {
			dependencyElement.SelectElement(keyGroupId).SetText(dep.Group)
			dependencyElement.SelectElement(keyArtifactId).SetText(dep.Artifact)
			dependencyElement.SelectElement(keyVersion).SetText(dep.Version)
		}
	}

	// 写入文件
	doc.Indent(xmlSpaces)
	err = doc.WriteToFile(filename)

	return
}
