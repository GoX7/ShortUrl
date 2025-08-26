package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

type (
	Engine struct {
		Block cipher.Block
		AEAD  cipher.AEAD
	}
)

func NewEngine(key string, engine *Engine) {
	hashKey := sha256.Sum256([]byte(key))
	block, err := aes.NewCipher(hashKey[:])
	if err != nil {
		fmt.Println("[-] crypto.newCipher: " + err.Error())
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("[-] crypto.newGCM: " + err.Error())
		return
	}

	fmt.Println("[+] crypto.new: ok")
	*engine = Engine{AEAD: gcm, Block: block}
}

func (eng *Engine) Seal(plaintext []byte) string {
	nonce := make([]byte, eng.AEAD.NonceSize())
	rand.Read(nonce)

	cipherResult := eng.AEAD.Seal(nonce, nonce, plaintext, nil)
	return base64.URLEncoding.EncodeToString(cipherResult)
}

func (eng *Engine) Open(cipherBase string) ([]byte, error) {
	cipherSlice, err := base64.URLEncoding.DecodeString(cipherBase)
	if err != nil {
		return nil, err
	}

	nonce := cipherSlice[:eng.AEAD.NonceSize()]
	ciphertext := cipherSlice[eng.AEAD.NonceSize():]

	return eng.AEAD.Open(nil, nonce, ciphertext, nil)
}
