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
	Lang string `default:"${PLUGIN_LANG=${LANG=go}}" validate:"required_without=Sources,oneof=go golang java js javascript dart"`
	// 源文件目录
	Source string `default:"${PLUGIN_SOURCE=${SOURCE=.}}"`
	// 源文件目录列表
	Sources map[string]string `default:"${PLUGIN_SOURCES=${SOURCES}}" validate:"required_without=Source"`
	// 版本
	Version string `default:"${PLUGIN_VERSION=${VERSION=1.0.0}}"`
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
	}
}

func (p *plugin) Fields() gox.Fields {
	return []gox.Field{
		field.Any(`sources`, p.Sources),
	}
}

func (p *plugin) isReplaced(from dependency, lang string) (to dependency, replaced bool) {
	for _, _replace := range p.Replaces[lang] {
		if `` != from.Module && from.Module == _replace.From.Module {
			replaced = true
			to = _replace.To
		}
	}

	return
}
