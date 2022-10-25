package main

import (
	"os"

	"github.com/tidwall/sjson"
)

func (p *plugin) json(filename string, elements ...jsonElement) (err error) {
	if `` == filename {
		return
	}

	var content []byte
	if content, err = os.ReadFile(filename); nil != err {
		return
	}

	for _, element := range elements {
		if content, err = sjson.SetBytes(content, element.path, element.value); nil != err {
			return
		}
	}
	err = os.WriteFile(filename, content, os.ModePerm)

	return
}
