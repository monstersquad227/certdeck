package v1

import (
	"certdeck/controller"
	"certdeck/repository"
	"certdeck/service"
	"github.com/gin-gonic/gin"
)

func EmailRegister(api *gin.RouterGroup) {
	ctrl := &controller.EmailController{
		EmailService: &service.EmailService{
			EmailRepo: &repository.EmailRepository{},
		},
	}

	api.POST("/email/send", ctrl.EmailSend)
}
