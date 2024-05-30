package extension

import (
    "bufio"
	"io"
    "io/ioutil"
    "os"
)

const (
	chunkSize = 64 * 1024
)

type TextFileHandler struct {
	FileName string
}

func NewTextFileHandler(fileName string) *TextFileHandler {
	return &TextFileHandler{
		FileName: fileName,
	}
}

func (handler *TextFileHandler) ReadAll() (string, error) {
	data, err := ioutil.ReadFile(handler.FileName)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func (handler *TextFileHandler) ReadAllLines() ([]string, error) {
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

func (handler *TextFileHandler) ReadChunks(chunkSize int) ([]string, error) {
    byteChunks, err := handler.RawReadChunks(chunkSize) // Call helper function
    if err != nil {
        return nil, err
    }

    var textChunks []string
    for _, chunk := range byteChunks {
        textChunks = append(textChunks, string(chunk))
    }
    return textChunks, nil
}

func (handler *TextFileHandler) RawReadChunks(chunkSize int) ([][]byte, error) {
	file, err := os.Open(handler.FileName)
    if err != nil {
        return nil, err
    }
    defer file.Close()

	reader := bufio.NewReader(file)
    var chunks [][]byte

	buffer := make([]byte, chunkSize)
    for {
        n, err := reader.Read(buffer)
        if err == io.EOF {
            if n > 0 {
                chunks = append(chunks, buffer[:n])
            }
            break
        } else if err != nil {
            return nil, err
        }
        chunks = append(chunks, buffer[:n])
    }

    return chunks, nil
}

func (handler *TextFileHandler) WriteAll(content string) error {
    data := []byte(content)
    err := ioutil.WriteFile("save_" + handler.FileName, data, 0644)
    return err
}