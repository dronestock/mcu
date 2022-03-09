package main

import (
	`github.com/dronestock/drone`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

type plugin struct {
	drone.PluginBase

	// 语言
	// nolint:lll
	Lang string `default:"${PLUGIN_LANG=${LANG=go}}" validate:"required_without=Inputs,oneof=go golang java js javascript dart"`
	// 源文件目录
	Source string `default:"${PLUGIN_SOURCE=${SOURCE=.}}"`
	// 源文件目录列表
	Sources map[string]string `default:"${PLUGIN_SOURCES=${SOURCES}}" validate:"required_without=Source"`
	// 依赖列表
	Dependencies map[string][]dependency `default:"${PLUGIN_DEPENDENCIES=${DEPENDENCIES}}"`
	// 替换列表
	Replaces map[string][]replace `default:"${PLUGIN_REPLACES=${REPLACES}}"`
}

func newPlugin() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Setup() (unset bool, err error) {
	p.Sources[p.Lang] = p.Source

	return
}

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		drone.NewStep(p.updates, drone.Name(`更新`)),
		drone.NewStep(p.replace, drone.Name(`替换`)),
	}
}

func (p *plugin) Fields() gox.Fields {
	return []gox.Field{
		field.Any(`sources`, p.Sources),
	}
}
