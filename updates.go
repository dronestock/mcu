package main

func (p *plugin) updates() (undo bool, err error) {
	for typ := range p.Sources {
		if err = p.update(typ); nil != err {
			return
		}
	}

	return
}
