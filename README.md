# jwks2pem

read jwks from stdin and write PEM to stdout

## What is it for?

OpenID Connect identity providers have a JWT keyset endpoint.

The most widely used JWT validation library in golang wants a PEM public key.

Converting them is a PITA, and I've wasted time figuring out the magic openssl incantations for the last time now.

## How to use it?

If you clone this repository outside your GOPATH, `go build -o jwks2pem main.go` should produce an simple executable
that takes care of the conversion for you.

It reads a jwks from standard input, and writes the PEM format of all keys in the set to stdout, so
you can do

```
curl -s <jwks-url> | ./jwks2pem
```

_Note: if your public keys in the key set aren't RSA keys, you will have to make a rather obvious change
to the code. Just replace the `&rsa.PublicKey{}` with the appropriate type._

## Acknowledgements

Many thanks to 
 - [the author of lestrrat-go/jwx](https://github.com/lestrrat-go/jwx)
 - [the author of square/go-jose](https://github.com/square/go-jose), even though I didn't end up using it, it gave me important pointers 
 - [everyone who participated in this stackoverflow discussion](https://stackoverflow.com/questions/55586188/how-to-generate-pem-from-jwk-in-go)
