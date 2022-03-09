package main

func (p *plugin) updates() (undo bool, err error) {
	for lang := range p.Sources {
		if err = p.update(lang); nil != err {
			return
		}
	}

	return
}
