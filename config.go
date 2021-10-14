package main

type config struct {
	lang         lang
	filepath     string
	version      string
	dependencies []*dependency
}

func (c *config) dependencyStrings() (strings []string) {
	strings = make([]string, 0, len(c.dependencies))
	for _, _dependency := range c.dependencies {
		strings = append(strings, _dependency.String())
	}

	return
}
