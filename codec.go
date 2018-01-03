package main

import (
	"github.com/murlokswarm/app"
	"encoding/base64"
)

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
		<p class="output">{{html .Output}}</p>
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
	defer app.Render(c)
	if c.RsaPrivateKey == "" {
		c.Output = "Please enter your rsa private key"
		return
	}

	if c.Data == "" {
		c.Output = "Please enter your data"
		return
	}

	rsa := RSASecurity{}
	if err := rsa.SetPublicKey(c.RsaPrivateKey); err != nil {
		c.Output = err.Error()
		return
	}

	data, err := rsa.Encrypt([]byte(c.Data))
	if err != nil {
		c.Output = err.Error()
		return
	}

	c.Output = base64.StdEncoding.EncodeToString(data)
}

// register component
func init() {
	app.RegisterComponent(&Codec{})
}
