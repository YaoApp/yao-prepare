package main

import (
	"github.com/yaoapp/yao/models/user"
)

func main() {
	user := user.New()
	user.Login()
}
