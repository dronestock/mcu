package main

type config struct {
	lang         lang
	path         string
	version      string
	dependencies []*dependency
}
