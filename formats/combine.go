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
	case JPG:
		return jpgWrap(f1.(*Jpg), f2)
	default:
		return nil, errors.New("unknown fileformat for f1")
	}
}

func jpgWrap(jpg *Jpg, f2 FormatChecker) ([]io.Reader, error) {
	switch f2.Format() {
	case PDF:
		var result []io.Reader

		first, err := jpg.Infect(f2.(Parasite))
		if err != nil {
			log.Printf("couldn't infect JPG with PDF skipping")
		}

		result = append(result, bytes.NewReader(first))

		second, err := jpg.Attach(f2.(Parasite).Reader())
		if err != nil {
			log.Printf("couldn't attach PDF to JPG")
		}

		return append(result, bytes.NewReader(second)), nil
	case ZIP:
		var result []io.Reader

		first, err := jpg.Infect(f2.(Parasite))
		if err != nil {
			log.Printf("couldn't infect JPG with ZIP skipping")
		}

		result = append(result, bytes.NewReader(first))

		second, err := jpg.Attach(f2.(Parasite).Reader())
		if err != nil {
			log.Printf("couldn't attach ZIP to JPG")
		}

		return append(result, bytes.NewReader(second)), nil
	case PNG:
		return nil, errors.New("PNG required offset at 0 can't attach or inject into JPG")
	case JPG:
		return nil, errors.New("failed to merge two file of the same type")
	default:
		return nil, errors.New("unknown fileformat for f2")
	}
}

func pngWrap(png *Png, f2 FormatChecker) ([]io.Reader, error) {
	switch f2.Format() {
	case PDF:
		var result []io.Reader

		first, err := png.Infect(f2.(Parasite))
		if err != nil {
			log.Printf("couldn't infect PNG with PDF skipping")
		}

		result = append(result, bytes.NewReader(first))

		second, err := png.Attach(f2.(Parasite).Reader())
		if err != nil {
			log.Printf("couldn't attach PDF to PNG")
		}

		return append(result, bytes.NewReader(second)), nil
	case ZIP:
		var result []io.Reader

		first, err := png.Infect(f2.(Parasite))
		if err != nil {
			log.Printf("couldn't infect PNG with ZIP skipping")
		}

		result = append(result, bytes.NewReader(first))

		second, err := png.Attach(f2.(Parasite).Reader())
		if err != nil {
			log.Printf("couldn't attach ZIP to PNG")
		}

		return append(result, bytes.NewReader(second)), nil
	case PNG:
		return nil, errors.New("failed to merge two file of the same type")
	case JPG:
		return nil, errors.New("JPG required offset at 0 can't inject or attach to PNG")
	default:
		return nil, errors.New("unknown fileformat for f2")
	}
}

func pdfWrap(pdf *Pdf, f2 FormatChecker) ([]io.Reader, error) {
	switch f2.Format() {
	case PDF:
		return nil, errors.New("failed to merge two file of the same type")
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
		return nil, errors.New("PNG requires offset at 0 can't attach or inject into PDF")
	case JPG:
		return nil, errors.New("JPG required offset at 0 can't attach or inject into PDF")
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
		return nil, errors.New("failed to merge two files of the same type")
	case PNG:
		return nil, errors.New("PNG requires offset at 0 can't attach or inject into ZIP")
	case JPG:
		return nil, errors.New("JPG required offset at 0 can't attach or inject into ZIP")
	default:
		return nil, errors.New("unknown fileformat for f2")
	}
}
