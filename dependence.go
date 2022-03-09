package main

type dependency struct {
	// 模块
	Module string `json:"module"`
	// 组
	Group string `json:"group"`
	// 坐标
	Artifact string `json:"artifact"`
	// 版本
	Version string `json:"version"`
}
