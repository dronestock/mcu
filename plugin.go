package main

import (
	"github.com/dronestock/drone"
	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base

	// 版本
	Version string `default:"${PLUGIN_VERSION=${VERSION=1.0.0}}"`
	// 模块列表
	Modules []*module `default:"${PLUGIN_MODULES=${MODULES}}" validate:"required,gte=1"`
	// 依赖列表
	Dependencies []*dependency `default:"${PLUGIN_DEPENDENCIES=${DEPENDENCIES}}"`

	modules map[string]*module
}

func newPlugin() drone.Plugin {
	return &plugin{
		modules: make(map[string]*module),
	}
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Setup() (unset bool, err error) {
	for _, _module := range p.Modules {
		p.modules[_module.Label] = _module
	}

	return
}

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		drone.NewStep(p.updates, drone.Name(`更新`)),
	}
}

func (p *plugin) Fields() gox.Fields {
	return []gox.Field{
		field.String(`version`, p.Version),
		field.Any(`modules`, p.Modules),
		field.Any(`dependencies`, p.Dependencies),
	}
}

func (p *plugin) each(labels []string, fun moduleFunc) (err error) {
	for _, label := range labels {
		if _module, ok := p.modules[label]; ok {
			fun(_module)
		} else {
			err = exc.NewField(`指定的模块没有定义`, field.String(`label`, label))
		}

		if nil != err {
			return
		}
	}

	return
}
