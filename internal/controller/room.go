package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/umarkotak/lul-er_BE/internal/models"
	"github.com/umarkotak/lul-er_BE/internal/service"
	"github.com/umarkotak/lul-er_BE/internal/utils"
)

func CreateRoom(c *gin.Context) {
	
	var create_room models.Room
	c.BindJSON(&create_room)

	result, err := service.CreateRoom(create_room)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	utils.RenderSuccess(c, result)
}