package main

type config struct {
	lang         lang
	filepath     string
	version      string
	dependencies []module
	replaces     []replace
}

func (c *config) dependencyStrings() (strings []string) {
	strings = make([]string, 0, len(c.dependencies))
	for _, _dependency := range c.dependencies {
		strings = append(strings, _dependency.String())
	}

	return
}

func (c *config) replaceStrings() (strings []string) {
	strings = make([]string, 0, len(c.replaces))
	for _, _replace := range c.replaces {
		strings = append(strings, _replace.String())
	}

	return
}

func (c *config) isReplaced(_module module) (to module, replaced bool) {
	for _, _replace := range c.replaces {
		if "" != _module.name && _module.name == _replace.from.name {
			replaced = true
			to = _replace.to
		}
	}

	return
}
