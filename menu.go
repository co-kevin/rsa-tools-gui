package main

import "github.com/murlokswarm/app"

type AppMainMenu struct{
}

func (a *AppMainMenu) Render() string {
	return `
<menu>
	<menu label="app">
		<menuitem label="About Us" />
	</menu>
	<EditMenu />
</menu>
`
}

type EditMenu struct {
}

func (e *EditMenu) Render() string {
	return `
<menu label="Edit">
    <menuitem label="Copy" shortcut="meta+c" />
	<menuitem label="Paste" shortcut="meta+v" />
</menu>
`
}

func init() {
	app.RegisterComponent(&AppMainMenu{})
	app.RegisterComponent(&EditMenu{})
}