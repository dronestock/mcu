package main

import (
	"context"
	_ "embed"
)

type stepUpdate struct {
	*plugin
}

func newUpdateStep(plugin *plugin) *stepUpdate {
	return &stepUpdate{
		plugin: plugin,
	}
}

func (u *stepUpdate) Runnable() bool {
	return 0 != len(u.Dependencies)
}

func (u *stepUpdate) Run(_ context.Context) (err error) {
	for _, dep := range u.Dependencies {
		if err = u.update(dep); nil != err {
			return
		}
	}

	return
}
