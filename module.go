package main

type module struct {
	// 标签
	Label string `json:"label" validate:"required"`
	// 模块
	Name string `json:"name"`
	// 组
	Group string `json:"group"`
	// 坐标
	Artifact string `json:"artifact"`
	// 版本
	Version string `json:"version" validate:"required"`
	// 作用域
	Scope string `json:"scope"`
}
