package main

import (
	"github.com/murlokswarm/app"
	_ "github.com/murlokswarm/mac"
)

func main() {
	app.OnLaunch = func() {
		appMenu := &AppMainMenu{}
		if menuBar, ok := app.MenuBar(); ok {
			menuBar.Mount(appMenu)
		}

		win := app.NewWindow(app.Window{
			Title: "RSA Tools GUI",
			Width: 800,
			Height: 640,
			TitlebarHidden: true,
		})

		codec := &Codec{}
		win.Mount(codec)
	}

	app.Run()
}