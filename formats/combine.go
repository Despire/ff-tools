package formats

import (
	"bytes"
	"errors"
	"fmt"
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
	case WASM:
		return wasmWrap(f1.(*Wasm), f2)
	default:
		return nil, errors.New("unknown fileformat for f1")
	}
}

func wasmWrap(wasm *Wasm, f2 FormatChecker) ([]io.Reader, error) {
	switch f2.Format() {
	case PDF, ZIP:
		var result []io.Reader

		first, err := wasm.Infect(f2.(Parasite))
		if err != nil {
			log.Printf("couldn't infect %s with %s skipping", wasm.Format().String(), f2.Format().String())
		}

		result = append(result, bytes.NewReader(first))

		second, err := wasm.Attach(f2.(Parasite).Reader())
		if err != nil {
			log.Printf("couldn't attach %s to %s", f2.Format().String(), wasm.Format().String())
		}

		return append(result, bytes.NewReader(second)), nil
	case PNG, JPG:
		return nil, fmt.Errorf("%s requires offset at 0 can't attach or inject into %s", f2.Format().String(), wasm.Format().String())
	case WASM:
		return nil, errors.New("failed to merge file of the same type")
	default:
		return nil, errors.New("unknown fileformat for f2")
	}
}

func jpgWrap(jpg *Jpg, f2 FormatChecker) ([]io.Reader, error) {
	switch f2.Format() {
	case PDF, ZIP:
		var result []io.Reader

		first, err := jpg.Infect(f2.(Parasite))
		if err != nil {
			log.Printf("couldn't infect %s with %s skipping", jpg.Format().String(), f2.Format().String())
		}

		result = append(result, bytes.NewReader(first))

		second, err := jpg.Attach(f2.(Parasite).Reader())
		if err != nil {
			log.Printf("couldn't attach %s to %s", f2.Format().String(), jpg.Format().String())
		}

		return append(result, bytes.NewReader(second)), nil
	case PNG, WASM:
		return nil, fmt.Errorf("%s requires offset at 0 can't attach or inject into %s", f2.Format().String(), jpg.Format().String())
	case JPG:
		return nil, errors.New("failed to merge two file of the same type")
	default:
		return nil, errors.New("unknown fileformat for f2")
	}
}

func pngWrap(png *Png, f2 FormatChecker) ([]io.Reader, error) {
	switch f2.Format() {
	case PDF, ZIP:
		var result []io.Reader

		first, err := png.Infect(f2.(Parasite))
		if err != nil {
			log.Printf("couldn't infect %s with %s skipping", png.Format().String(), f2.Format().String())
		}

		result = append(result, bytes.NewReader(first))

		second, err := png.Attach(f2.(Parasite).Reader())
		if err != nil {
			log.Printf("couldn't attach %s to %s", f2.Format().String(), png.Format().String())
		}

		return append(result, bytes.NewReader(second)), nil
	case PNG:
		return nil, errors.New("failed to merge two file of the same type")
	case JPG, WASM:
		return nil, fmt.Errorf("%s requires offset at 0 can't inject or attach to %s", f2.Format().String(), png.Format().String())
	default:
		return nil, errors.New("unknown fileformat for f2")
	}
}

func pdfWrap(pdf *Pdf, f2 FormatChecker) ([]io.Reader, error) {
	switch f2.Format() {
	case ZIP, PDF:
		var result []io.Reader

		first, err := pdf.Infect(f2.(Parasite))
		if err != nil {
			log.Printf("couldn't infect %s with %s skipping", pdf.Format().String(), f2.Format().String())
		}

		result = append(result, bytes.NewReader(first))

		second, err := pdf.Attach(f2.(Parasite).Reader())
		if err != nil {
			log.Printf("couldn't attach %s to %s", f2.Format().String(), pdf.Format().String())
		}

		return append(result, bytes.NewReader(second)), nil
	case PNG, WASM, JPG:
		return nil, fmt.Errorf("%s requires offset at 0 can't attach or inject into %s", f2.Format().String(), pdf.Format().String())
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
	case WASM, PNG, JPG:
		return nil, fmt.Errorf("%s requires offset at 0 can't attach or inject into %s", f2.Format().String(), z.Format().String())
	default:
		return nil, errors.New("unknown fileformat for f2")
	}
}
