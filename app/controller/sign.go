package controller

import (
	"gin-user-center/app/common"

	"gin-user-center/app/schema"
	"gin-user-center/app/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var logger = common.Logger

func Login(c *gin.Context) {
	session := sessions.Default(c)
	ctx := util.Context{Ctx: c}

	var user schema.AddUser
	// user := &schema.User{}
	if err := ctx.ValidateJSON(&user); err != nil {
		return
	}

	// form-data
	// name := c.PostForm("name")
	// password := c.PostForm("password")

	// json
	name := "admin"
	password := "admin"
	logger.Info("name:", name, "password:", password)

	// todo从db验证
	if name == "admin" && password == "admin" {
		session.Set("user_name", name)
		session.Save()

		content := map[string]string{
			"data": "success",
		}
		ctx.Response(0, nil, content)
	} else {
		ctx.Response(401, common.LOGIN_FAIL, nil)
	}
	return
}
