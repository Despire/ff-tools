package main

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Despire/ff-tools/crypto"
	"github.com/Despire/ff-tools/formats"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if len(os.Args) > 1 && os.Args[1] == "pngToPdf" {
		if len(os.Args[2:]) != 3 {
			return errors.New("need 3 arguments for collision detection (Key||png||pdf)")
		}

		return crypto.PngToPdf(os.Args[2:])
	}

	if len(os.Args) > 1 && os.Args[1] == "encrypt" {
		if len(os.Args[2:]) < 2 {
			return errors.New("need 2 arguments for encrpytion (Key||data||?IV)")
		}

		if _, err := hex.DecodeString(os.Args[2]); err != nil {
			return err
		}

		iv := make([]byte, aes.BlockSize)
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			return err
		}

		if len(os.Args[2:]) > 2 {
			prefix, err := hex.DecodeString(os.Args[4])
			if err != nil {
				return err
			}

			copy(iv, prefix)
		}

		return crypto.Encrypt(os.Args[2:], iv)
	}

	if len(os.Args) > 1 && os.Args[1] == "decrypt" {
		if len(os.Args[2:]) < 2 {
			return errors.New("need 2 arguments for decryption (Key||data)")
		}

		if _, err := hex.DecodeString(os.Args[2]); err != nil {
			return err
		}

		iv := make([]byte, aes.BlockSize)
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			return err
		}

		if len(os.Args[2:]) > 2 {
			prefix, err := hex.DecodeString(os.Args[4])
			if err != nil {
				return err
			}

			copy(iv, prefix)
		}

		return crypto.Decrypt(os.Args[2:], iv)
	}

	if len(os.Args) != 3 {
		return errors.New("need exactly two arguments where each of them is a file")
	}

	var file1 formats.FormatChecker
	{
		b, err := os.ReadFile(os.Args[1])
		if err != nil {
			return err
		}

		file1, err = formats.Find(b)
		if err != nil {
			return err
		}
	}

	var file2 formats.FormatChecker
	{
		b, err := os.ReadFile(os.Args[2])
		if err != nil {
			return err
		}

		file2, err = formats.Find(b)
		if err != nil {
			return err
		}
	}

	out, err := formats.Combine(file1, file2)
	if err != nil {
		return err
	}

	for i, o := range out {
		b, _ := io.ReadAll(o)
		if err := os.WriteFile(
			fmt.Sprintf(
				"%d-%s-%s-combined.%s.%s",
				i,
				os.Args[1],
				os.Args[2],
				file1.Format().String(),
				file2.Format().String(),
			),
			b,
			os.ModePerm,
		); err != nil {
			return err
		}
	}

	return nil
}
