package user

import (
	"testing"
)

func TestNew(t *testing.T) {
	user := New()
	user.Query()
	user.Login()
}
