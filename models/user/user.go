package user

import (
	"github.com/yaoapp/yao/lib/json"
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
	db := user.GetDB()
	rows := []map[string]interface{}{}
	db.Raw("SELECT name,mobile from `user` WHERE mobile IS NOT NULL limit 2").Find(&rows)
	json.PrettyPrint(rows)
}
