package main

import (
	"log"
	"time"

	_ "github.com/mouxiaohui/err-nil/cmd"

	"github.com/briandowns/spinner"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {
	// 加载动画
	spinner := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	spinner.Color("yellow", "bold")
	spinner.Suffix = " 查找中..., 骚等 🤗"
	spinner.Start()

	time.Sleep(4 * time.Second) // Run for some time to simulate work

	// 停止加载动画
	spinner.Stop()

	return nil
}
