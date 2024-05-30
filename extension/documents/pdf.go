package extension

import ()

type PdfFileReader struct {
	FileName string
}

func NewPdfFileReader(fileName string) *PdfFileReader {
	return &PdfFileReader{
		FileName: fileName,
	}
}