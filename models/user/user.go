package user

import (
	"fmt"

	"github.com/yaoapp/yao/lib/model"
)

// User the user model
type User struct {
	model.Model
}

// New create user
func New() *User {
	user := &User{}
	user.Model = model.New("user")
	return user
}

// Login the user login
func (user *User) Login() {
	option := user.GetOption()
	fmt.Printf("%s user login \n", option.Name)
}
