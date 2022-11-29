package formats

import (
	"errors"
	"io"
)

type FileFormat int

const (
	PDF  FileFormat = 0x1
	ZIP  FileFormat = 0x2
	PNG  FileFormat = 0x3
	LAST FileFormat = 0x4
)

func (f FileFormat) String() string {
	switch f {
	case PDF:
		return "pdf"
	case ZIP:
		return "zip"
	case PNG:
		return "png"
	default:
		panic("unknown fileformat")
	}
}

type FormatChecker interface {
	Format() FileFormat
}

type Attacher interface {
	Attach(reader io.Reader) ([]byte, error)
}

type Parasite interface {
	IsParasite() bool
	Reader() io.Reader
}

func Find(f []byte) (FormatChecker, error) {
	pdf, err := NewPdf(f)
	if err == nil {
		return pdf, nil
	}

	z, err := NewZip(f)
	if err == nil {
		return z, nil
	}

	p, err := NewPng(f)
	if err == nil {
		return p, nil
	}

	return nil, errors.New("no such format is registered")
}
