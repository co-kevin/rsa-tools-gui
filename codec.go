package main

import (
	"encoding/base64"
	"github.com/murlokswarm/app"
)

type Codec struct {
	RsaKey string // user input rsa private/public key
	Data   string // user input encrypt data or origin data
	Output string // output data
}

func (c *Codec) Render() string {
	return `
<div class="WindowLayout">
    <div class="HelloBox">
		<h1>Encode or Decode use RSA and Base64</h1>
        <textarea row="15" placeholder="Enter your rsa key" onchange="OnChangeKey">
		</textarea>
		<textarea row="15" placeholder="Enter your data" onchange="OnChangeData">
		</textarea>
		<button onclick="OnClickEncodeButton">Encode</button><button onclick="OnClickDecodeButton">Decode</button>
		<p class="output">{{html .Output}}</p>
    </div>
</div>
`
}

// on change key, set data to context
func (c *Codec) OnChangeKey(arg app.ChangeArg) {
	c.RsaKey = arg.Value
	app.Render(c)
}

// on change data, set data to context
func (c *Codec) OnChangeData(arg app.ChangeArg) {
	c.Data = arg.Value
	app.Render(c)
}

// on click encode button, encode use public key
func (c *Codec) OnClickEncodeButton(arg app.EventArg) {
	defer app.Render(c)
	if c.RsaKey == "" {
		c.Output = "Please enter your rsa public key"
		return
	}

	if c.Data == "" {
		c.Output = "Please enter your data"
		return
	}

	rsa := RSASecurity{}
	if err := rsa.SetPublicKey(c.RsaKey); err != nil {
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

// on click decode button, decode use private key
func (c *Codec) OnClickDecodeButton(arg app.EventArg) {
	defer app.Render(c)
	if c.RsaKey == "" {
		c.Output = "Please enter your rsa private key"
		return
	}

	if c.Data == "" {
		c.Output = "Please enter your data"
		return
	}

	rsa := RSASecurity{}
	if err := rsa.SetPrivateKey(c.RsaKey); err != nil {
		c.Output = err.Error()
		return
	}

	if input, err := base64.StdEncoding.DecodeString(string(c.Data)); err != nil {
		c.Output = err.Error()
	} else {
		data, err := rsa.Decrypt(input)
		if err != nil {
			c.Output = err.Error()
			return
		}
		c.Output = string(data)
	}
}

// register component
func init() {
	app.RegisterComponent(&Codec{})
}
