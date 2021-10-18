package main

import (
	`fmt`
	`strings`
)

type replace struct {
	from module
	to   module
}

func (r *replace) String() string {
	return fmt.Sprintf("%s => %s", r.from.String(), r.to.String())
}

func parseReplaces(key string) (replaces []replace) {
	_config := env(key)
	if "" == _config {
		return
	}

	originals := strings.Split(_config, `,`)
	replaces = make([]replace, len(originals))
	for _, original := range originals {
		replaces = append(replaces, parseReplace(original))
	}

	return
}

func parseReplace(original string) (replace replace) {
	var _configs []string
	defer func() {
		replace = newReplace(_configs)
	}()

	if _configs = strings.Split(original, "=>"); 2 == len(_configs) {
		return
	}
	if _configs = strings.Split(original, "->"); 2 == len(_configs) {
		return
	}
	if _configs = strings.Split(original, " "); 2 == len(_configs) {
		return
	}

	return
}

func newReplace(configs []string) (_replace replace) {
	length := len(configs)
	if nil != configs && 2 == length {
		_replace.from = parseModule(configs[0])
		_replace.to = parseModule(configs[1])
	}

	return
}
