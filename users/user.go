package users

import (
	"fmt"
	"master-go/models"
	"time"
)

func UserRegister() {
	user := new(models.User)
	user.AddUser(10, "Milton", time.Now(), true)

	fmt.Println(user)
}
