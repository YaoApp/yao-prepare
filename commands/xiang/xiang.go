package main

import (
	"github.com/gobuffalo/packr"
	"github.com/yaoapp/yao/models/user"
)

func main() {
	// pack model
	packr.NewBox("../../models")
	user := user.New()
	user.Login()
}
