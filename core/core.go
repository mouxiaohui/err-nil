package core

import (
	"fmt"

	"github.com/mouxiaohui/err-nil/cmd"

	"github.com/gookit/color"
)

func Run() error {
	// 获取需要查找的go文件列表
	fileList, err := getGoFiles(cmd.PROJECT_PATH)
	if err != nil {
		return err
	}

	// 分析所有文件
	errNilFiles, total, err := findErrNil(&fileList)
	if err != nil {
		return err
	}

	if total < 10 {
		print("🥳 ")
	} else if total < 30 {
		print("😨 ")
	} else {
		print("🤢 ")
	}
	color.New(color.FgYellow).Print("一共有 ")
	color.New(color.FgYellow, color.OpBold, color.OpUnderscore).Print(total)
	color.New(color.FgYellow).Print(" 个 `err != nil`, `err == nil`\n")

	for _, errNil := range errNilFiles {
		if errNil.counter > 0 {
			color.Notice.Print(fmt.Sprintf("%d 个", errNil.counter))
			color.Note.Print(" => ")
			color.Info.Print(fmt.Sprintf("%s\n", errNil.path))
		}
	}

	return nil
}
