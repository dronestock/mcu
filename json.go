package main

import (
	`io/ioutil`
	`os`

	`github.com/tidwall/sjson`
)

func json(filename string, elements ...jsonElement) (err error) {
	if "" == filename {
		return
	}

	var content []byte
	if content, err = ioutil.ReadFile(filename); nil != err {
		return
	}

	for _, element := range elements {
		if content, err = sjson.SetBytes(content, element.path, element.value); nil != err {
			return
		}
	}
	err = ioutil.WriteFile(filename, content, os.ModePerm)

	return
}
