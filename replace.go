package main

import (
	`strings`
)

type replace struct {
	from module
	to   module
}

func parseReplaces(originals ...string) (replaces []replace) {
	// 防止有nil的依赖，因为originals始终有一个值（上层在用strings.Split时，空字符串也会分隔成一个有一个空字符串的数组）
	replaces = make([]replace, 0)
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
