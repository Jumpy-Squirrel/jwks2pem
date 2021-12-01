package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/lestrrat-go/jwx/jwk"
	"io/ioutil"
	"os"
)

func pemFromJwks(input *os.File, output *os.File) string {
	keySetBytes, err := ioutil.ReadAll(input)
	if err != nil {
		return fmt.Sprintf("failed to read jwks json from stdin: %s", err.Error())
	}

	keySet, err := jwk.Parse(keySetBytes)
	if err != nil {
		return fmt.Sprintf("failed to parse json keyset: %s", err.Error())
	}

	for i := 0; i < keySet.Len(); i++ {
		key, ok := keySet.Get(i)
		if !ok {
			return fmt.Sprintf("failed to get key #%d from keyset", i+1)
		}

		pubKey := &rsa.PublicKey{}
		err = key.Raw(pubKey)
		if err != nil {
			return fmt.Sprintf("failed to extract raw rsa public key for key #%d: %s", i+1, err.Error())
		}

		pubData, err := x509.MarshalPKIXPublicKey(pubKey)
		if err != nil {
			return fmt.Sprintf("failed to marshal key #%d to public key: %s", i+1, err.Error())
		}
		if err := pem.Encode(output, &pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: pubData,
		}); err != nil {
			return fmt.Sprintf("failed to pem encode key #%d: %s", i+1, err.Error())
		}
	}

	return ""
}

func main() {
	errMsg := pemFromJwks(os.Stdin, os.Stdout)
	if errMsg != "" {
		_, _ = os.Stderr.WriteString(errMsg + "\n")
		os.Exit(1)
	}
}
