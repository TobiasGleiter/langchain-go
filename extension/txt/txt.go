package extension

import (
    "bufio"
    "io/ioutil"
    "os"
)

type TextFileHandler struct {
	FileName string
}

func NewTextFileHandler(fileName string) *TextFileHandler {
	return &TextFileHandler{
		FileName: fileName,
	}
}

func (handler *TextFileHandler) ReadAllToString() (string, error) {
	data, err := ioutil.ReadFile(handler.FileName)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func (handler *TextFileHandler) ReadAllLinesToStringArray() ([]string, error) {
	file, err := os.Open(handler.FileName)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return lines, nil
}

func (handler *TextFileHandler) WriteContentToFile(content, fileName string) error {
    data := []byte(content)
    err := ioutil.WriteFile(fileName, data, 0644)
    return err
}