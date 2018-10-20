package main

import (
	"html/template"
	"path"
	"sync"
)

type templateCache struct {
	m     sync.Mutex
	cache map[string]*template.Template
}

func newTemplateCache() *templateCache {
	return &templateCache{
		cache: make(map[string]*template.Template),
	}
}

func (tc *templateCache) get(viewPaths []string, helpers template.FuncMap) (*template.Template, error) {
	tc.m.Lock()
	defer tc.m.Unlock()
	key := viewPaths[len(viewPaths)-1]
	tmpl := tc.cache[key]
	if tmpl != nil {
		return tmpl, nil
	}

	var err error
	tmplName := path.Base(viewPaths[0])
	tmpl, err = template.New(tmplName).
		Funcs(helpers).
		ParseFiles(viewPaths...)
	if err != nil {
		return nil, err
	}
	tc.cache[key] = tmpl
	return tmpl, nil
}
