package formats

import (
	"bytes"
	"io"
)

type Png struct {
	contents []byte
}

func NewPng(contents []byte) (*Png, error) {

}

func (p *Png) Format() FileFormat { return PNG }

func (p *Png) IsParasite() bool { return false }

func (p *Png) Reader() io.Reader { return bytes.NewReader(p.contents) }

func (p *Png) Infect(file Parasite) ([]byte, error) {

}

func (p *Png) Attach(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return append(append(p.contents, "\n"...), b...), nil
}
