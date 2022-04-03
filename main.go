package main

import (
	"log"

	_ "github.com/mouxiaohui/err-nil/cmd"
	"github.com/mouxiaohui/err-nil/core"
)

func main() {
	if err := core.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
