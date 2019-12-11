package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (zipEntry *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(zipEntry.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer reader.Close()

	for _, f := range reader.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			// Possible resource leak, 'defer' is called in a for loop.
			// 这句代码会造成循环结束后才开始回收资源，而不是执行了一次循环就回收一次资源
			// defer rc.Close()

			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			rc.Close()
			return data, zipEntry, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (zipEntry *ZipEntry) String() string {
	return zipEntry.absPath
}
