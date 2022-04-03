package core

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mouxiaohui/err-nil/cmd"
)

func Run() error {
	// 加载动画
	spinner := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	spinner.Color("yellow", "bold")
	spinner.Suffix = " 查找中..., 骚等 🤗"
	spinner.Start()

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

	// 停止加载动画
	spinner.Stop()

	for _, v := range errNilFiles {
		fmt.Println(v)
	}

	return nil
}
