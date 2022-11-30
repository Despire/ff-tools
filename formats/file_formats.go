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
	JPG  FileFormat = 0x4
	LAST FileFormat = 0x5
)

func (f FileFormat) String() string {
	switch f {
	case PDF:
		return "pdf"
	case ZIP:
		return "zip"
	case PNG:
		return "png"
	case JPG:
		return "jpg"
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

	j, err := NewJpg(f)
	if err == nil {
		return j, nil
	}

	return nil, errors.New("no such format is registered")
}
