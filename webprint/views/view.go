package views

import "github.com/a-h/templ"

type View struct {
	View func() templ.Component
	Wasm func()
}
