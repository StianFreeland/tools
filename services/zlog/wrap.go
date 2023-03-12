package zlog

import "os"

type wrap struct {
	file *os.File
}

func (w wrap) Write(b []byte) (n int, err error) {
	return w.file.Write(b)
}

func (w wrap) Sync() error {
	return nil
}
