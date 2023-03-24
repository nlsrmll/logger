package logger

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Logger struct {
	path string
	file *os.File
}

func OpenFile(path string) (*Logger, error) {
	logger := &Logger{
		path: path,
	}

	return logger, nil
}

func (l *Logger) Log(s string) error {
	return l.writeNewLine(s)
}

func (l *Logger) Error(s string) error {
	return l.writeNewLine(fmt.Sprintf("ERROR: %s", s))
}

func (l *Logger) open() error {
	file, err := os.OpenFile(l.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	l.file = file

	return nil

}

func (l *Logger) close() error {
	return l.file.Close()
}

func (l *Logger) writeNewLine(s string) error {

	l.open()

	defer l.close()

	writer := bufio.NewWriter(l.file)
	_, err := writer.WriteString(fmt.Sprintf("%s %s\n", timeStamp(), s))

	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

func timeStamp() string {
	now := time.Now()

	return fmt.Sprintf("[%s %s]", now.Format("02.01.2006"), now.Format("15:04:05"))

}
