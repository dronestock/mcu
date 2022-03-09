package main

func (p *plugin) update(lang string) (undo bool, err error) {
	source := p.Sources[lang]
	dependencies := p.Dependencies[lang]
	switch lang {
	case langGo, langGolang:
		undo, err = p.golang(source, dependencies...)
	case langJava:
		undo, err = p.java(source, dependencies...)
	case langJs, langJavascript:
		undo, err = p.js(source)
	case langDart:
		undo, err = p.dart(source, dependencies...)
	}

	return
}
