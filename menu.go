package main

import "github.com/murlokswarm/app"

type AppMainMenu struct{
}

func (a *AppMainMenu) Render() string {
	return `
<menu>
	<menu label="app">
		<menuitem label="About" selector="orderFrontStandardAboutPanel:"/>
		<menuitem label="Quit" selector="terminate:" shortcut="meta+q"/>
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
	<menuitem label="Undo" selector="undo:" shortcut="meta+z" />
	<menuitem label="Redo" selector="redo:" shortcut="meta+shift+z" />
    <menuitem label="Copy" selector="copy:" shortcut="meta+c" />
	<menuitem label="Paste" selector="paste:" shortcut="meta+v" />
	<menuitem label="Delete" selector="delete:" />
	<menuitem label="Select All" selector="selectAll:" shortcut="meta+a" />
</menu>
`
}

func init() {
	app.RegisterComponent(&AppMainMenu{})
	app.RegisterComponent(&EditMenu{})
}