package main

func (p *plugin) update(typ string) (err error) {
	source := p.Sources[typ]
	dependencies := p.Dependencies[typ]
	switch typ {
	case typeGo, typeGolang:
		err = p.golang(source, dependencies...)
	case typeJava:
		err = p.java(source, dependencies...)
	case typeJs, typeJavascript:
		err = p.js(source)
	case typeDart:
		err = p.dart(source, dependencies...)
	}

	return
}
