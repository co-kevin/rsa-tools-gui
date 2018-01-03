package main

import (
	"crypto/rsa"
	"bytes"
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
	"crypto/rand"
	"io"
	"errors"
)

var RSA = &RSASecurity{}

type RSASecurity struct {
	pubStr string          //公钥字符串
	priStr string          //私钥字符串
	pubkey *rsa.PublicKey  //公钥
	prikey *rsa.PrivateKey //私钥
}

// 公钥加密
func (rsas *RSASecurity) Encrypt(input []byte) ([]byte, error) {
	if rsas.pubkey == nil {
		return []byte(""), errors.New(`请先设置公钥`)
	}
	output := bytes.NewBuffer(nil)
	err := pubKeyIO(rsas.pubkey, bytes.NewReader(input), output)
	if err != nil {
		return []byte(""), err
	}
	return ioutil.ReadAll(output)
}

// 私钥解密
func (rsas *RSASecurity) Decrypt(input []byte) ([]byte, error) {
	if rsas.prikey == nil {
		return []byte(""), errors.New(`请先设置私钥`)
	}
	output := bytes.NewBuffer(nil)
	err := priKeyIO(rsas.prikey, bytes.NewReader(input), output)
	if err != nil {
		return []byte(""), err
	}

	return ioutil.ReadAll(output)
}

// 设置公钥
func (rsas *RSASecurity) SetPublicKey(pubStr string) (err error) {
	rsas.pubStr = pubStr
	rsas.pubkey, err = rsas.getPublickey()
	return err
}

// 设置私钥
func (rsas *RSASecurity) SetPrivateKey(priStr string) (err error) {
	rsas.priStr = priStr
	rsas.prikey, err = rsas.getPrivatekey()
	return err
}

// *rsa.PrivateKey
func (rsas *RSASecurity) getPublickey() (*rsa.PublicKey, error) {
	return getPubKey([]byte(rsas.pubStr))
}

// *rsa.PublicKey
func (rsas *RSASecurity) getPrivatekey() (*rsa.PrivateKey, error) {
	return getPriKey([]byte(rsas.priStr))
}

// 设置私钥
func getPriKey(privatekey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privatekey)
	if block == nil {
		return nil, errors.New("获取私钥错误")
	}
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return pri, nil
	}
	pri2, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pri2.(*rsa.PrivateKey), nil
}

// 设置公钥
func getPubKey(publickey []byte) (*rsa.PublicKey, error) {
	// 堆
	block, _ := pem.Decode(publickey)
	if block == nil {
		return nil, errors.New("获取公钥错误")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), err
}

// 公钥加密Reader
func pubKeyIO(pub *rsa.PublicKey, in io.Reader, out io.Writer) (err error) {
	k := (pub.N.BitLen() + 7) / 8

	k = k - 11

	buf := make([]byte, k)
	var b []byte
	size := 0
	for {
		size, err = in.Read(buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if size < k {
			b = buf[:size]
		} else {
			b = buf
		}
		b, err = rsa.EncryptPKCS1v15(rand.Reader, pub, b)
		if err != nil {
			return err
		}
		if _, err = out.Write(b); err != nil {
			return err
		}
	}
	return nil
}

// 私钥解密Reader
func priKeyIO(pri *rsa.PrivateKey, r io.Reader, w io.Writer) (err error) {
	k := (pri.N.BitLen() + 7) / 8

	buf := make([]byte, k)
	var b []byte
	size := 0
	for {
		size, err = r.Read(buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if size < k {
			b = buf[:size]
		} else {
			b = buf
		}
		b, err = rsa.DecryptPKCS1v15(rand.Reader, pri, b)

		if err != nil {
			return err
		}
		if _, err = w.Write(b); err != nil {
			return err
		}
	}
	return nil
}
