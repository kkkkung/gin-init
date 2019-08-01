package Admin

import (
	"gin-init/models"
	"gin-init/pkg/logging"
)

type Admin struct {
	models.Model

	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Password string `json:"-"`
}

func (Admin) TableName() string {
	return "admins"
}

func GetInfo(id int) (user Admin) {
	if err := models.Db.Debug().Where(map[string]interface{}{"id": id}).Find(&user).Error; err != nil {
		logging.Info(err)
	}
	return
}

func Login(user Admin) bool {
	var find Admin
	if models.Db.Where(user).Select("id").First(&find); find.ID > 0 {
		return true
	}
	return false
}
