package CryptoUtil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/mhthrh/Voiceflex/Utils/ConverUtil"
	"io"
	"os"
)

type CryptoInterface interface {
	Decrypt()
	Encrypt()
	Sha256()
	Md5Sum()
}
type Crypto struct {
	key      string
	FilePath string
	Text     string
	Result   string
}

func NewKey() *Crypto {
	c := new(Crypto)
	c.key = "AnKoloft@~delNazok!12345" // key parameter must be 16, 24 or 32,
	return c
}

func (k *Crypto) Sha256() {
	h := sha256.New()
	h.Write([]byte(k.Text))
	k.Result = fmt.Sprintf("%x", h.Sum(nil))
}

func (k *Crypto) Md5Sum() error {
	file, err := os.Open(k.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return err
	}
	k.Result = hex.EncodeToString(hash.Sum(nil))
	return nil
}

func (k *Crypto) Encrypt() error {
	key := []byte(k.key)
	plaintext := []byte(k.Text)
	c, err := aes.NewCipher(key)
	if err != nil {
		return err

	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err

	}

	k.Result = ConverUtil.Byte64(gcm.Seal(nonce, nonce, plaintext, nil))
	return nil
}

func (k *Crypto) Decrypt() error {
	key := []byte(k.key)
	bb, _ := base64.StdEncoding.DecodeString(k.Text)
	ciphertext := []byte(bb)
	c, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	t, e := gcm.Open(nil, nonce, ciphertext, nil)
	if e != nil {
		return err
	}
	k.Result = ConverUtil.BytesToString(t)
	return nil

}
