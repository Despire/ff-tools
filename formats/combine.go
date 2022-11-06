package formats

import (
	"bytes"
	"errors"
	"io"
	"log"
)

func Combine(f1 FormatChecker, f2 FormatChecker) ([]io.Reader, error) {
	switch f1.Format() {
	case PDF:
		return pdfWrap(f1.(*Pdf), f2)
	case ZIP:
		return zipWrap(f1.(*Zip), f2)
	case PNG:
		return pngWrap(f1.(*Png), f2)
	default:
		return nil, errors.New("unknown fileformat for f1")
	}
}

func pngWrap(png *Png, f2 FormatChecker) ([]io.Reader, error) {
	switch f2.Format() {
	case PDF:
		panic("implemnt me!")
	case ZIP:
		panic("implemnt me!")
	case PNG:
		panic("implemnt me!")
	default:
		return nil, errors.New("unknown fileformat for f2")
	}
}

func pdfWrap(pdf *Pdf, f2 FormatChecker) ([]io.Reader, error) {
	switch f2.Format() {
	case PDF:
		return nil, errors.New("failed to mergeTemplate two file of the same type")
	case ZIP:
		var result []io.Reader

		first, err := pdf.Infect(f2.(Parasite))
		if err != nil {
			log.Printf("couldn't infect PDF with ZIP skipping")
		}

		result = append(result, bytes.NewReader(first))

		second, err := pdf.Attach(f2.(Parasite).Reader())
		if err != nil {
			log.Printf("couldn't attach zip to pdf")
		}

		return append(result, bytes.NewReader(second)), nil
	case PNG:
		panic("implement me!")
	default:
		return nil, errors.New("unknown fileformat for f2")
	}
}

func zipWrap(z *Zip, f2 FormatChecker) ([]io.Reader, error) {
	switch f2.Format() {
	case PDF:
		var result []io.Reader

		first, err := z.Infect(f2.(Parasite))
		if err != nil {
			log.Printf("couldn't infect PDF with ZIP skipping")
		}

		result = append(result, bytes.NewReader(first))

		second, err := z.Attach(f2.(Parasite).Reader())
		if err != nil {
			log.Printf("couldn't attach zip to pdf")
		}

		return append(result, bytes.NewReader(second)), nil
	case ZIP:
		return nil, errors.New("failed to mergeTemplate two files of the same type")
	case PNG:
		panic("implement me!")
	default:
		return nil, errors.New("unknown fileformat for f2")
	}
}
