package extension

import (
	"bufio"
	"io"
	"os"
)

type TextSplitter struct {
	FileName  string
	ChunkSize int
}

func NewTextSplitter(fileName string, chunkSize int) *TextSplitter {
	return &TextSplitter{
		FileName:  fileName,
		ChunkSize: chunkSize,
	}
}

func (splitter *TextSplitter) TextSplitter() ([]string, error) {
	byteChunks, err := splitter.RawTextSplitter()
	if err != nil {
		return nil, err
	}

	var textChunks []string
	for _, chunk := range byteChunks {
		textChunks = append(textChunks, string(chunk))
	}
	return textChunks, nil
}

func (splitter *TextSplitter) RawTextSplitter() ([][]byte, error) {
	file, err := os.Open(splitter.FileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var chunks [][]byte

	buffer := make([]byte, splitter.ChunkSize)
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
