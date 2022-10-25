package main

func (p *plugin) updates() (undo bool, err error) {
	for _, _dependency := range p.Dependencies {
		if err = p.update(_dependency); nil != err {
			return
		}
	}

	return
}
