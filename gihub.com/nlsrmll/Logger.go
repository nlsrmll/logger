package logger

import "os"

type Logger struct {
	path string
	file *os.File
}

func OpenFile(path string) (*Logger, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	logger := &Logger{
		path: path,
		file: file,
	}

	return logger, err
}
