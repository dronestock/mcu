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
	conf.path = env("FILEPATH")
	conf.version = env("VERSION")
	conf.dependencies = parseDependencies(strings.Split(env("DEPENDENCIES"), ",")...)
	defaults.SetDefaults(conf)

	// 记录配置日志信息
	logger.Info(
		"功能加载配置完成",
		field.String("lang", string(conf.lang)),
		field.String("path", conf.path),
		field.String("version", conf.version),
		field.Strings("dependencies", conf.dependencyStrings()...),
	)

	switch conf.lang {
	case langGo:
		fallthrough
	case langGolang:
		err = golang(conf, logger)
	case langJavascript:
		fallthrough
	case langJs:
		err = js(conf, logger)
	case langDart:
		err = dart(conf, logger)
	default:
		err = notSupportLang
	}

	if nil != err {
		panic(err)
	}
}
