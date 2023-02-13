package components

import (
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"service_app/db"
	"service_app/model"
)

var User = db.User{}

func CreateUser(c *gin.Context) error {

	sql := db.ConnectDB()

	var body model.AccountRequest

	User.Phone = c.Param("phone")

	err := c.Bind(&body)

	if err != nil {
		return err
	}

	if sql.Migrator().HasTable(db.User{}) {
		err = sql.AutoMigrate(db.User{})
	}

	if err != nil {
		return err
	}

	sql.Model(db.User{}).Where("phone = ?", User.Phone).First(&User)

	if User.ID != 0 {
		return errors.New("该手机已注册")
	}

	u1 := uuid.Must(uuid.NewV4(), nil)

	User.Uuid = u1.String()

	User.Password = body.Password

	err = sql.Create(&User).Error

	return err
}

func LoginUser() {

}
