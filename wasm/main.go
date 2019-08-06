package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"syscall/js"
)

func generateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, _ := rsa.GenerateKey(rand.Reader, 1024)
	return privkey, &privkey.PublicKey
}

func exportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkeybytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkeypem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkeybytes,
		},
	)

	return string(pubkeypem), nil
}

func exportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privkeybytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkeypem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkeybytes,
		},
	)
	return string(privkeypem)
}

func main() {
	// Create the keys
	priv, pub := generateRsaKeyPair()

	// Export the keys to pem string
	privKey := exportRsaPrivateKeyAsPemStr(priv)
	pubKey, _ := exportRsaPublicKeyAsPemStr(pub)

	js.Global().Call("updatePrivKey", privKey)
	js.Global().Call("updatePubKey", pubKey)
}
