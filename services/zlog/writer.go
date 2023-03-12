package zlog

import (
	"fmt"
	"os"
	"time"
)

type writer struct {
	filename string
	file     *os.File
}

func (w *writer) Write(b []byte) (int, error) {
	filename := "logs/" + time.Now().Format("2006-01-02") + ".txt"
	if filename != w.filename {
		if err := w.openFile(filename); err != nil {
			return 0, err
		}
	}
	n, err := w.file.Write(b)
	if err != nil {
		fmt.Println(moduleName, "write,", "err:", err)
		return 0, err
	}
	return n, nil
}

func (w *writer) openFile(filename string) error {
	if w.file != nil {
		if err := w.file.Close(); err != nil {
			fmt.Println(moduleName, "open file,", "err:", err)
			return err
		}
	}
	if err := createDirectory(); err != nil {
		return err
	}
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(moduleName, "open file,", "err:", err)
		return err
	}
	newFilename := "zlog.txt"
	createSymlink(filename, newFilename)
	w.filename = filename
	w.file = file
	return nil
}

func createDirectory() error {
	if _, err := os.Stat("logs"); err == nil {
		return nil
	} else if !os.IsNotExist(err) {
		fmt.Println(moduleName, "create directory,", "err:", err)
		return err
	}
	if err := os.Mkdir("logs", 0755); err != nil {
		fmt.Println(moduleName, "create directory,", "err:", err)
		return err
	}
	return nil
}

func createSymlink(filename string, newFilename string) {
	if _, err := os.Lstat(newFilename); err != nil {
		if !os.IsNotExist(err) {
			fmt.Println(moduleName, "create symlink,", "err:", err)
			return
		}
	} else {
		if err := os.Remove(newFilename); err != nil {
			fmt.Println(moduleName, "create symlink,", "err:", err)
			return
		}
	}
	if err := os.Symlink(filename, newFilename); err != nil {
		fmt.Println(moduleName, "create symlink,", "err:", err)
		return
	}
}
