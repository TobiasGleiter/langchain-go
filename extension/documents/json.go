package extension

import (
	"io/ioutil"
)

type JsonFileReader struct {
	FileName string
}

func NewJsonFileHandler(fileName string) *JsonFileReader {
	return &JsonFileReader{
		FileName: fileName,
	}
}

func (handler *JsonFileReader) ReadAllToString() (string, error) {
	data, err := ioutil.ReadFile(handler.FileName)
    if err != nil {
        return "", err
    }
    return string(data), nil
}