package core

import (
	"io/ioutil"
	"path/filepath"
)

// 递归查找文件夹下所有的 .go 文件
func getGoFiles(path string) ([]string, error) {
	var filesList []string

	filesInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range filesInfo {
		if fileInfo.IsDir() {
			if list, err := getGoFiles(filepath.Join(path, fileInfo.Name())); err != nil {
				return nil, err
			} else {
				filesList = append(filesList, list...)
			}
		} else {
			if filepath.Ext(fileInfo.Name()) == ".go" {
				filesList = append(filesList, filepath.Join(path, fileInfo.Name()))
			}
		}
	}

	return filesList, nil
}
