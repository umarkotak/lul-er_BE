package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/umarkotak/lul-er_BE/internal/models"
	"github.com/umarkotak/lul-er_BE/internal/service"
	"github.com/umarkotak/lul-er_BE/internal/utils"
)

func Register(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	result, err := service.Register(user)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	utils.RenderSuccess(c, result)
}

func Login(c *gin.Context) {

	var user models.User
	c.BindJSON(&user)

	result, err := service.Login(user)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	utils.RenderSuccess(c, result)

}
