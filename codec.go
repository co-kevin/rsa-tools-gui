package main

import "github.com/murlokswarm/app"

type Codec struct {
	RsaPrivateKey string // user input rsa private key
	Data          string // user input encrypt data or origin data
	Output        string // output data
}

func (c *Codec) Render() string {
	return `
<div class="WindowLayout">
    <div class="HelloBox">
		<h1>Encode data with rsa and base64</h1>
        <textarea row="8" placeholder="Enter your rsa private key" onchange="OnChangePrivateKey">
		</textarea>
		<textarea row="8" placeholder="Enter your data" onchange="OnChangeData">
		</textarea>
		<button onclick="OnClickButton">Encode</button>
		<h2>{{html .Output}}</h2>
    </div>
</div>
`
}

// on swatch encoder or decoder
func (c *Codec) OnSwitchMethod(arg app.EventArg) {

}

// on change private key, set data to context
func (c *Codec) OnChangePrivateKey(arg app.ChangeArg) {
	c.RsaPrivateKey = arg.Value
	app.Render(c)
}

// on change data, set data to context
func (c *Codec) OnChangeData(arg app.ChangeArg) {
	c.Data = arg.Value
	app.Render(c)
}

// on click submit button
func (c *Codec) OnClickButton(arg app.EventArg) {
	c.Output = "on click button"
	app.Render(c)
	println(c.RsaPrivateKey)
	println(c.Data)
}

// register component
func init() {
	app.RegisterComponent(&Codec{})
}
