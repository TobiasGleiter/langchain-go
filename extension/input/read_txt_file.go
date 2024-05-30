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

	for scanner.Scan() {
		line := scanner.Bytes() // Get the bytes of the line
		content = append(content, line...)
		content = append(content, '\n') // Append newline character
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