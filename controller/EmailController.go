package controller

import (
	"certdeck/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmailController struct {
	EmailService service.EmailServiceInterface
}

func (ctrl *EmailController) EmailSend(c *gin.Context) {
	type email struct {
		Email string `json:"email"`
	}
	req := &email{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	err = ctrl.EmailService.SendEmail(req.Email)
	if err != nil {
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}
