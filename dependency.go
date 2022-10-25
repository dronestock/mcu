package main

type dependency struct {
	// 类型
	// nolint:lll
	Type string `default:"go" validate:"required,oneof=go golang java js javascript dart"`
	// 源文件目录
	Source string `default:"."`
	// 模块列表
	Modules []string `validate:"required"`
}
