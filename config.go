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

func (c *config) isReplaced(_module module) (replaced bool, to module) {
	for _, _replace := range c.replaces {
		if _module.name == _replace.from.name {
			replaced = true
			to = _replace.to
		}
	}

	return
}

func (c *config) toReplacesMap() (replacesMap map[string]replace) {
	replacesMap = make(map[string]replace)
	for _, _replace := range c.replaces {
		replacesMap[_replace.from.name] = _replace
	}
	return
}
