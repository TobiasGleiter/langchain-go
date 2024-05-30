package extension

import (
	"io/ioutil"
	"bytes"
)

type PdfFileReader struct {
	FileName string
}

func NewPdfFileReader(fileName string) *PdfFileReader {
	return &PdfFileReader{
		FileName: fileName,
	}
}

func (reader *PdfFileReader) ReadRawContent() ([]byte, error) {
    data, err := ioutil.ReadFile(reader.FileName)
    if err != nil {
        return nil, err
    }
    return data, nil
}

  
