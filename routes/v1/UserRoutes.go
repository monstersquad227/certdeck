package v1

import "github.com/gin-gonic/gin"

func UserRegister(api *gin.RouterGroup) {
	api.POST("/login")
}
