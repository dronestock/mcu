package main

func (p *plugin) update(dependency dependency) (err error) {
	switch dependency.Type {
	case typeGo, typeGolang:
		err = p.golang(dependency.Source, dependency.Modules)
	case typeJava:
		err = p.java(dependency.Source, dependency.Modules)
	case typeJs, typeJavascript:
		err = p.js(dependency.Source, dependency.Modules)
	case typeDart:
		err = p.dart(dependency.Source, dependency.Modules)
	}

	return
}
