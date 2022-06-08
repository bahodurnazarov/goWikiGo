package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	fn := p.Title + ".txt"
	return ioutil.WriteFile(fn, p.Body, 0600)
}

func loadPage(title string) *Page {
	fn := title + ".txt"
	body, _ := ioutil.ReadFile(fn)
	return &Page{Title: title, Body: body}
}