package crypto

import (
	"bytes"
	"crypto/aes"
	"errors"
	"fmt"
	"io"
	"os"
)

func PngToPdf(args []string) error {
	key := args[0]
	png := args[1]
	pdf := args[2]

	pngfd, err := os.Open(png)
	if errors.Is(err, os.ErrNotExist) {
		return err
	}

	pngdata, err := io.ReadAll(pngfd)
	if err != nil {
		return err
	}

	pdffd, err := os.Open(pdf)
	if errors.Is(err, os.ErrNotExist) {
		return err
	}

	pdfdata, err := io.ReadAll(pdffd)
	if err != nil {
		return err
	}

	const marker = "%PDF-1obj\nstream"

	iv, err := DecryptPlainTextToIV(pngdata, key, marker)
	if err != nil {
		return err
	}

	fmt.Printf("iv: %x", iv)

	ciphertext, err := EncryptRaw(pngdata, key, iv)
	if err != nil {
		return err
	}

	ciphertext = append(ciphertext, "\nendstream\nendobj\n"...)
	ciphertext = append(ciphertext, pdfdata...)

	//ciphertext, err = pkcs7Pad(ciphertext, aes.BlockSize)
	//if err != nil {
	//	return err
	//}

	fd, err := os.Create(fmt.Sprintf("pngToPdf") + ".stg")
	if err != nil {
		return err
	}

	if _, err := io.Copy(fd, bytes.NewReader(ciphertext)); err != nil {
		return err
	}

	return fd.Close()
}

func xor(a, b []byte) []byte {
	out := make([]byte, aes.BlockSize)
	for i, val := range b {
		if i >= aes.BlockSize {
			break
		}

		out[i] = a[i] ^ val
	}

	return out
}

func DecryptPlainTextToIV(plaintext []byte, key string, target string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	out := make([]byte, aes.BlockSize)
	block.Decrypt(out, []byte(target))
	return xor(out, plaintext[:aes.BlockSize]), nil
}
