package core

import (
	"fmt"

	"github.com/mouxiaohui/err-nil/cmd"
)

func Run() error {
	// 获取需要查找的go文件列表
	fileList, err := getGoFiles(cmd.PROJECT_PATH)
	if err != nil {
		return err
	}

	// 分析所有文件
	errNilFiles, err := findErrNil(&fileList)
	if err != nil {
		return err
	}

	for _, v := range errNilFiles {
		fmt.Println(v)
	}

	return nil
}
