package main

import (
	"errors"
	"fmt"
	"github.com/Despire/ff-tools/formats"
	"io"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
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
