package common

import (
	"go.go.org/go/isokit"
	"honnef.co/go/js/dom"
)

type Env struct {
	TemplateSet    *isokit.TemplateSet
	Window         dom.Window
	Document       dom.Document
	PrimaryContent dom.Element
}
