package main

import (
	`errors`
	`strings`

	`github.com/mcuadros/go-defaults`
	`github.com/storezhang/glog`
	`github.com/storezhang/gox/field`
)

var notSupportLang = errors.New("不支持的语言")

func main() {
	var err error
	// 有错误，输出错误日志
	var logger glog.Logger
	if logger, err = glog.New(); nil != err {
		panic(err)
	}

	// 取各种参数
	conf := new(config)
	conf.lang = lang(env("LANG"))
	conf.path = env("PATH")
	conf.version = env("VERSION")
	conf.dependencies = parseDependencies(strings.Split(env("DEPENDENCIES"), ",")...)
	defaults.SetDefaults(conf)

	switch conf.lang {
	case langGo:
		fallthrough
	case langGolang:
		err = golang(conf)
	case langJavascript:
		fallthrough
	case langJs:
		err = js(conf)
	default:
		err = notSupportLang
	}

	if nil != err {
		logger.Error("处理失败", field.Error(err))
	}
}
