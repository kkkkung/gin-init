package authController

import (
	"gin-init/models/Auth"
	"gin-init/pkg/e"
	"gin-init/pkg/logging"
	"gin-init/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	msg := ""
	mobile := c.PostForm("mobile")
	password := c.PostForm("password")

	valid := validation.Validation{}
	a := auth{Username: mobile, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.BAD_REQUEST
	if ok {
		id := Auth.CheckAndReturnId(mobile, password)
		if id > 0 {
			token, err := util.GenerateToken(id)
			if err == nil {
				code = e.CREATED
				data["token"] = token
			}
		} else {
			code = e.UNAUTHORIZED
			msg = "用户名或密码错误"
		}

	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : msg,
		"data" : data,
	})
}
