package main

import (
	"github.com/murlokswarm/app"
	_ "github.com/murlokswarm/mac"
)

func main() {
	app.OnLaunch = func() {
		win := app.NewWindow(app.Window{
			Title: "Hello World",
			Width: 1280,
			Height: 720,
			TitlebarHidden: true,
		})

		hello := &Hello{}
		win.Mount(hello)
	}

	app.Run()
}