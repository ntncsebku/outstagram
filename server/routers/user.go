package routers

import (
	"log"
	"outstagram/server/injection"

	"github.com/gin-gonic/gin"
)

func UserAPIRouter(router *gin.RouterGroup) {
	userController, err := injection.InitializeUserController()
	if err != nil {
		log.Fatal(err.Error())
	}

	router.GET("/:userID", userController.GetUsersInfo)
	router.GET("/:userID/storyboard", userController.GetUserStoryBoard)
}
