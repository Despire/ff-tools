package formats

var (
	PdfHeaderStart = []byte("%PDF")
	ZipHeaderStart = []byte("PK\x03\x04")
	PngHeaderStart = []byte("\x89PNG\r\n\x1a\n")
	JpgHeaderStart = []byte("\xff\xd8\xff")
)
