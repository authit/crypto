package crypto

import (
	"bytes"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"fmt"
	"io/ioutil"
	"log"
)

func printMessage() {
	decbuf := bytes.NewBuffer([]byte(encryptedMessage))
	result, err := armor.Decode(decbuf)
	if err != nil {
		log.Fatal(err)
	}

	md, err := openpgp.ReadMessage(result.Body, nil, func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		return []byte("golang"), nil
	}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("dec version:", result.Header["Version"])
	fmt.Println("dec type:", result.Type)
	bytes, err := ioutil.ReadAll(md.UnverifiedBody)
	fmt.Println("md:", string(bytes))
}

func main() {
	ent, err := openpgp.NewEntity("John Andersen", "", "johnandersenpdx@gmail.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ent)
}

const encryptedMessage = `-----BEGIN PGP MESSAGE-----
Version: GnuPG v1.4.15 (Darwin)

jA0EAwMCSk50dj2NcPtgySLEBzaZ+zgxODr+/7BeQPHyW4JsOrYXptQKFgQtewBg
HBi7
=dz85
-----END PGP MESSAGE-----`
