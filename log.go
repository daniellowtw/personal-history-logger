package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type logFactory struct {
	current    *os.File
	dateString string
	format     string
	dirPath    string
}

func (f *logFactory) prepFolder() error {
	dateToday := time.Now().Format(f.format)
	fp := filepath.Join(f.dirPath, dateToday+".log")
	ff, err := os.OpenFile(fp, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0770)
	if os.IsNotExist(err) {
		return fmt.Errorf("error creating log: %v", err)
	}
	f.current = ff
	return nil
}

func (f *logFactory) Get(t time.Time) (io.Writer, error) {
	if f.dateString == t.Format(f.format) {
		return f.current, nil
	}
	f.current.Close()
	if err := f.prepFolder(); err != nil {
		return nil, err
	}
	return f.current, nil
}
