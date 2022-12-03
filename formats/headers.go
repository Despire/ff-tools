package formats

var (
	PdfHeaderStart  = []byte("%PDF")
	ZipHeaderStart  = []byte("PK\x03\x04")
	PngHeaderStart  = []byte("\x89PNG\r\n\x1a\n")
	JpgHeaderStart  = []byte("\xff\xd8\xff")
	WASMHEaderStart = []byte("\x00\x61\x73\x6D\x01\x00\x00\x00")
)
