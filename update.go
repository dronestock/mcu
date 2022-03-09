package main

func (p *plugin) update(lang string) (err error) {
	source := p.Sources[lang]
	dependencies := p.Dependencies[lang]
	switch lang {
	case langGo, langGolang:
		err = p.golang(source, dependencies...)
	case langJava:
		err = p.java(source, dependencies...)
	case langJs, langJavascript:
		err = p.js(source)
	case langDart:
		err = p.dart(source, dependencies...)
	}

	return
}
