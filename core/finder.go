package core

import (
	"bufio"
	"os"
	"regexp"
)

var (
	regex *regexp.Regexp
)

func generateRegexp() error {
	r, err := regexp.Compile(`err [!|=]= nil`)
	if err != nil {
		return err
	}
	regex = r

	return nil
}

type ErrNilFile struct {
	path    string
	counter uint
}

// 查找 `err != nil` 与 `err == nil`
func (e *ErrNilFile) parseFile() error {

	file, err := os.Open(e.path)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := bufio.NewScanner(file)
	for {
		if !buf.Scan() {
			break
		}
		line := buf.Text()
		if regex.MatchString(line) {
			e.counter++
		}
	}

	return nil
}

// 分析所有文件，查找 `err != nil` 与 `err == nil`
func findErrNil(fileList *[]string) ([]ErrNilFile, uint, error) {
	var total uint
	if err := generateRegexp(); err != nil {
		return nil, total, err
	}

	var errNilFiles []ErrNilFile
	for _, file := range *fileList {
		errNil := ErrNilFile{path: file}
		if err := errNil.parseFile(); err != nil {
			return nil, total, err
		}
		total += errNil.counter
		errNilFiles = append(errNilFiles, errNil)
	}

	return errNilFiles, total, nil
}
