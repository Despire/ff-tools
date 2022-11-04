package formats

import "errors"

func Combine(f1 FormatChecker, f2 FormatChecker) ([]byte, error) {
	switch f1.Format() {
	case PDF:
		return pdfWrap(f1.(*Pdf), f2)
	case ZIP:
		return zipWrap(f1.(*Zip), f2)
	default:
		return nil, errors.New("unknown fileformat for f1")
	}
}

func pdfWrap(pdf *Pdf, f2 FormatChecker) ([]byte, error) {
	switch f2.Format() {
	case PDF:
		return nil, errors.New("failed to mergeTemplate two file of the same type")
	case ZIP:
		return pdf.Infect(f2.(Parasite))
	default:
		return nil, errors.New("unknown fileformat for f2")
	}
}

func zipWrap(z *Zip, f2 FormatChecker) ([]byte, error) {
	switch f2.Format() {
	case PDF:
		return z.Infect(f2.(Parasite))
	case ZIP:
		return nil, errors.New("failed to mergeTemplate two files of the same type")
	default:
		return nil, errors.New("unknown fileformat for f2")
	}
}
