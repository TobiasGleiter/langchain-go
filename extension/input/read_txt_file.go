package extension

import (
	"os"
	"bufio"
)

const (
	chunkSize = 64 * 1024
)

func ReadTextFile(filename string) ([]byte, error) {
	file, err := os.Open(filename + ".txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var content []byte
	scanner := bufio.NewScanner(file)

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	totalSize := fileInfo.Size()

	content = make([]byte, 0, totalSize)

	for scanner.Scan() {
		line := scanner.Bytes()
		content = append(content, line...)
		content = append(content, '\n')
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return content, nil
}

func SaveToTextFile(filename string, content []byte) error {
	file, err := os.Create(filename + ".txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return nil
}